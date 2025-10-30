package outbox

import (
	"context"
	"errors"
	"log"
	"time"
)

type WorkerOption func(*OutboxWorker)

func WithBatchSize(n int) WorkerOption {
	return func(w *OutboxWorker) { w.batchSize = n }
}

func WithPollInterval(d time.Duration) WorkerOption {
	return func(w *OutboxWorker) { w.pollInterval = d }
}

func WithRetryInterval(d time.Duration) WorkerOption {
	return func(w *OutboxWorker) { w.retryInterval = d }
}

func WithMaxRetry(n int) WorkerOption {
	return func(w *OutboxWorker) { w.maxRetry = n }
}

func WithWindow(d time.Duration) WorkerOption {
	return func(w *OutboxWorker) { w.window = d }
}

type OutboxWorker struct {
	batchSize     int
	maxRetry      int
	retryInterval time.Duration
	pollInterval  time.Duration
	window        time.Duration
}

func NewOutboxWorker(opts ...WorkerOption) OutboxWorker {
	w := OutboxWorker{
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
	FriendRequestEventsHandler interface {
		// Возвращает списки успешных и проваленных id; err — для фатальных ошибок батча.
		HandleBatch(ctx context.Context, events []*Event) (succeeded []string, failed []string, err error)
	}
)

// OutboxFriendRequestWorker — обработка outbox-событий именно по заявкам в друзья.
type OutboxFriendRequestWorker struct {
	OutboxWorker

	repo    Repository
	tm      TransactionManager
	handler FriendRequestEventsHandler
}

// NewOutboxFriendRequestWorker конструктор с дефолтами.
func NewOutboxFriendRequestWorker(
	repo Repository,
	tm TransactionManager,
	h FriendRequestEventsHandler,
	opts ...WorkerOption,
) *OutboxFriendRequestWorker {
	w := &OutboxFriendRequestWorker{
		OutboxWorker: NewOutboxWorker(opts...),
		repo:         repo,
		tm:           tm,
		handler:      h,
	}

	return w
}

// Run — запускает бесконечный цикл обработки до отмены ctx.
// Селектит batch с FOR UPDATE SKIP LOCKED, обрабатывает, коммитит.
func (w *OutboxFriendRequestWorker) Run(ctx context.Context) error {
	log.Println("OutboxFriendRequestWorker started")

	t := time.NewTicker(w.pollInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			log.Println("OutboxFriendRequestWorker tick")

			// Один "тик" — одна транзакция
			if err := w.tm.RunRepeatableRead(ctx, w.FetchFriendRequest); err != nil {
				log.Printf("outbox FriendRequest: error: %v\n", err)
			}

			log.Println("OutboxFriendUpdatedWorker tick")
			if err := w.tm.RunRepeatableRead(ctx, w.FetchFriendUpdated); err != nil {
				log.Printf("outbox: friendUpdated error: %v\n", err)
			}
		}
	}
}

// Fetch обработка событий
func (w *OutboxFriendRequestWorker) FetchFriendRequest(ctx context.Context) error {
	log.Println("OutboxFriendRequestWorker.Fetch start")
	defer log.Println("OutboxFetchFriendRequestWorker.Fetch end")

	err := w.fetch(ctx, AggregateTypeFriendRequest, EventTypeFriendRequest)

	return err
}

func (w *OutboxFriendRequestWorker) FetchFriendUpdated(ctx context.Context) error {
	log.Println("OutboxFriendRequestWorker.Fetch start")
	defer log.Println("OutboxFetchFriendRequestWorker.Fetch end")

	err := w.fetch(ctx, AggregateTypeFriendUpdated, EventTypeFriendUpdated)

	return err
}

func (w *OutboxFriendRequestWorker) fetch(ctx context.Context, aggregateType AggregateType, eventType EventType) error {
	var (
		now  = time.Now().UTC()
		from = now.Add(-w.window)
	)

	events := w.repo.SearchEvents(
		ctx,
		// 1-я ступень pruning
		//WithNotBefore(from),
		//WithNotAfter(now),
		// 2-я ступень pruning
		WithAggregateType(aggregateType),
		// 3-я ступень pruning
		WithEventType(eventType),
		// фильтрация
		WithOnlyUnpublished(),
		//WithDueAt(now),
		WithMaxRetryCount(w.maxRetry),
		WithLimit(w.batchSize),
		WithLock(), // FOR UPDATE
	)
	if len(events) == 0 {
		log.Println("outbox no events")
		return nil
	}

	succeeded, failed, err := w.handler.HandleBatch(ctx, events)
	if err != nil {
		log.Printf("outbox batch handle error: %v", err)
		return err
	}

	if len(succeeded) > 0 {
		e := w.repo.UpdateEvents(
			ctx,
			WithUpdateNotBefore(from),
			WithUpdateNotAfter(now),
			WithUpdateAggregateType(aggregateType),
			WithUpdateEventType(eventType),
			WithUpdateIDs(succeeded...),

			SetPublishedAt(now),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	if len(failed) > 0 {
		e := w.repo.UpdateEvents(
			ctx,
			WithUpdateNotBefore(from),
			WithUpdateNotAfter(now),
			WithUpdateAggregateType(aggregateType),
			WithUpdateEventType(eventType),
			WithUpdateIDs(failed...),

			IncRetry(1),
			SetNextAttemptAt(now.Add(w.retryInterval)),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
