package inbox

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"log"
	"time"
)

type WorkerOption func(*InboxWorker)

func WithBatchSize(n int) WorkerOption {
	return func(w *InboxWorker) { w.batchSize = n }
}

func WithPollInterval(d time.Duration) WorkerOption {
	return func(w *InboxWorker) { w.pollInterval = d }
}

func WithRetryInterval(d time.Duration) WorkerOption {
	return func(w *InboxWorker) { w.retryInterval = d }
}

func WithMaxRetry(n int) WorkerOption {
	return func(w *InboxWorker) { w.maxRetry = n }
}

func WithWindow(d time.Duration) WorkerOption {
	return func(w *InboxWorker) { w.window = d }
}

type InboxWorker struct {
	batchSize     int
	maxRetry      int
	retryInterval time.Duration
	pollInterval  time.Duration
	window        time.Duration
}

func NewInboxWorker(opts ...WorkerOption) InboxWorker {
	w := InboxWorker{
		batchSize:     100,
		maxRetry:      10,
		retryInterval: 5 * time.Minute,
		pollInterval:  10 * time.Second,
		window:        24 * time.Hour,
	}
	for _, opt := range opts {
		opt(&w)
	}
	return w
}

type (
	//FriendREquestEventHandler - обработчик событий по заявкам в друзья
	MessagesHandler interface {
		// Возвращает списки успешных и проваленных id; err — для фатальных ошибок батча.
		HandleBatch(ctx context.Context, messages []*model.Inbox) (succeeded []string, failed []string, err error)
	}
)

// InboxFriendRequestWorker — обработка Inbox-событий именно по заявкам в друзья.
type InboxMessageWorker struct {
	InboxWorker

	repo    Repository
	tm      TransactionManager
	handler MessagesHandler
}

// NewInboxFriendRequestWorker конструктор с дефолтами.
func NewInboxMessageWorker(
	repo Repository,
	tm TransactionManager,
	h MessagesHandler,
	opts ...WorkerOption,
) *InboxMessageWorker {
	w := &InboxMessageWorker{
		InboxWorker: NewInboxWorker(opts...),
		repo:        repo,
		tm:          tm,
		handler:     h,
	}

	return w
}

// Run — запускает бесконечный цикл обработки до отмены ctx.
// Селектит batch с FOR UPDATE SKIP LOCKED, обрабатывает, коммитит.
func (w *InboxMessageWorker) Run(ctx context.Context) error {
	log.Println("InboxFriendRequestWorker started")

	t := time.NewTicker(w.pollInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			log.Println("InboxChatWorker tick")

			// Один "тик" — одна транзакция
			if err := w.tm.RunRepeatableRead(ctx, w.FetchMessageSent); err != nil {
				log.Printf("Inbox FriendRequest: error: %v\n", err)
			}

		}
	}
}

// Fetch обработка событий
func (w *InboxMessageWorker) FetchMessageSent(ctx context.Context) error {
	log.Println("InboxMessageWorker.Fetch start")
	defer log.Println("InboxFetchMessageWorker.Fetch end")

	err := w.fetch(ctx)

	return err
}

func (w *InboxMessageWorker) fetch(ctx context.Context) error {
	var (
		now  = time.Now().UTC()
		from = now.Add(-w.window)
		_    = from
	)

	messages := w.repo.SearchMessage(
		ctx,
		// 1-я ступень pruning
		//WithNotBefore(from),
		//WithNotAfter(now),
		// 2-я ступень pruning
		//WithAggregateType(aggregateType),
		// 3-я ступень pruning
		//WithEventType(eventType),
		// фильтрация
		//WithOnlyUnpublished(),
		//WithDueAt(now),
		//WithMaxRetryCount(w.maxRetry),
		WithLimit(w.batchSize),
		WithLock(), // FOR UPDATE
	)
	if len(messages) == 0 {
		log.Println("Inbox no messages")
		return nil
	}

	succeeded, failed, err := w.handler.HandleBatch(ctx, messages)
	if err != nil {
		log.Printf("Inbox batch handle error: %v", err)
		return err
	}

	if len(succeeded) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			//WithUpdateNotBefore(from),
			//WithUpdateNotAfter(now),
			//WithUpdateAggregateType(aggregateType),
			//WithUpdateEventType(eventType),
			WithUpdateIDs(succeeded...),

			//SetPublishedAt(now),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	if len(failed) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			//WithUpdateNotBefore(from),
			//WithUpdateNotAfter(now),
			//WithUpdateAggregateType(aggregateType),
			//WithUpdateEventType(eventType),
			WithUpdateIDs(failed...),

			//IncRetry(1),
			//SetNextAttemptAt(now.Add(w.retryInterval)),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
