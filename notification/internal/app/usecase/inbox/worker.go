package inbox

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"time"
)

type WorkerOption func(*InboxWorker)

func WithBatchSize(n int) WorkerOption {
	return func(w *InboxWorker) { w.batchSize = n }
}

type InboxWorker struct {
	batchSize    int
	maxAttempts  int
	pullInterval time.Duration
}

func NewInboxWorker(opts ...WorkerOption) InboxWorker {
	w := InboxWorker{
		batchSize:    100,
		maxAttempts:  10,
		pullInterval: 5 * time.Second,
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
	logger.Info(ctx, "InboxFriendRequestWorker started")

	t := time.NewTicker(w.pullInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			logger.Info(ctx, "InboxMessageWorker tick")

			// Один "тик" — одна транзакция
			if err := w.tm.RunRepeatableRead(ctx, w.Fetch); err != nil {
				logger.Errorf(ctx, "Inbox Message: error: %v\n", err)
			}

		}
	}
}

// Fetch обработка событий
func (w *InboxMessageWorker) Fetch(ctx context.Context) error {
	logger.Info(ctx, "InboxMessageWorker.Fetch start")
	defer logger.Info(ctx, "InboxFetchMessageWorker.Fetch end")

	messages := w.repo.SearchMessage(
		ctx,
		WithStatus([]string{model.StatusReceived, model.StatusFailed}),
		WithMaxAttempts(w.maxAttempts),
		WithLimit(w.batchSize),
		WithLock(), // FOR UPDATE
	)
	if len(messages) == 0 {
		logger.Info(ctx, "Inbox no messages")
		return nil
	}

	allIds := make([]string, len(messages))
	for i, message := range messages {
		allIds[i] = message.ID
	}

	if len(allIds) > 0 {
		_ = w.repo.UpdateMessages(
			ctx,
			WithUpdateIDs(allIds...),
			WithUpdateStatus(model.StatusProcessing),
			WithUpdateAttempts(1),
		)
	}

	succeeded, failed, err := w.handler.HandleBatch(ctx, messages)
	if err != nil {
		logger.Errorf(ctx, "Inbox batch handle error: %v", err)
		return err
	}

	if len(succeeded) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			WithUpdateIDs(succeeded...),
			WithUpdateProcessedAt(time.Now()),
			WithUpdateStatus(model.StatusProcessed),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	if len(failed) > 0 {
		e := w.repo.UpdateMessages(
			ctx,
			WithUpdateIDs(failed...),
			WithUpdateStatus(model.StatusFailed),
			WithUpdateLastError("error"),
		)
		if e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
