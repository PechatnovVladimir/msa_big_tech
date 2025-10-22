package socialeventshandler

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/modules/outbox"
	"log"
	"slices"
	"time"
)

type TopicResolver func(e *outbox.Event) (topic string, key string)

type KafkaBatchHandler struct {
	producer     sarama.SyncProducer
	resolve      TopicResolver
	maxBatchSize int
	closeTimeout time.Duration
}

type KafkaHandlerOption func(*KafkaBatchHandler)

// WithTopic фиксирует один топик, ключ = AggregateID.
func WithTopic() KafkaHandlerOption {
	return func(h *KafkaBatchHandler) {
		h.resolve = func(e *outbox.Event) (string, string) { return string(e.EventType), e.AggregateID }
	}
}

// WithTopicResolver позволяет выбрать топик/ключ динамически.
func WithTopicResolver(r TopicResolver) KafkaHandlerOption {
	return func(h *KafkaBatchHandler) { h.resolve = r }
}

// WithMaxBatchSize настраивает размер чанка для SendMessages (по умолчанию 500).
func WithMaxBatchSize(n int) KafkaHandlerOption {
	return func(h *KafkaBatchHandler) {
		if n > 0 {
			h.maxBatchSize = n
		}
	}
}

// WithCloseTimeout задаёт таймаут при закрытии продюсера.
func WithCloseTimeout(d time.Duration) KafkaHandlerOption {
	return func(h *KafkaBatchHandler) { h.closeTimeout = d }
}

func NewKafkaBatchHandler(producer sarama.SyncProducer, opts ...KafkaHandlerOption) *KafkaBatchHandler {
	handler := &KafkaBatchHandler{
		producer:     producer,
		maxBatchSize: 100,
		closeTimeout: 10 * time.Second,
		resolve:      func(e *outbox.Event) (string, string) { return string(e.AggregateType), e.AggregateID },
	}
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// Close аккуратно закрывает продюсер.
func (h *KafkaBatchHandler) Close() error {
	done := make(chan struct{})
	var cerr error
	go func() {
		cerr = h.producer.Close()
		close(done)
	}()
	select {
	case <-done:
		return cerr
	case <-time.After(h.closeTimeout):
		return errors.New("kafka: close timeout")
	}
}

const headerEventID = "event_id"

func (h *KafkaBatchHandler) HandleBatch(ctx context.Context, events []*outbox.Event) (succeeded []string, failed []string, err error) {
	if len(events) == 0 {
		log.Println("KafkaBatchHandler", "nothing to send")
		return nil, nil, nil
	}

	defer func() {
		if err != nil {
			log.Println("HandleBatch", err)
		} else {
			log.Println("HandleBatch", "succeeded", succeeded, "failed", failed)
		}
	}()

	chunks := chunkEvents(events, h.maxBatchSize)
	succeeded = make([]string, 0, len(events))
	failed = make([]string, 0, len(events))

	for _, evs := range chunks {
		select {
		case <-ctx.Done():
			return succeeded, append(failed, ids(evs)...), ctx.Err()
		default:
		}

		// Собираем батч сообщений
		msgs := make([]*sarama.ProducerMessage, 0, len(evs))
		for _, e := range evs {
			topic, key := h.resolve(e)
			msg := &sarama.ProducerMessage{
				Topic:     topic,
				Key:       sarama.StringEncoder(key), // партиционирование по ключу
				Value:     sarama.ByteEncoder(e.Payload),
				Timestamp: e.CreatedAt, // полезно для таймлайнов
				Metadata:  e.ID,        // чтобы распознать ошибку по id
				Headers: []sarama.RecordHeader{
					{
						Key:   []byte(headerEventID),
						Value: []byte(e.ID),
					},
				},
			}
			msgs = append(msgs, msg)
		}

		// Отправляем батчом
		if sendErr := h.producer.SendMessages(msgs); sendErr != nil {
			// Частичные ошибки приходят как sarama.ProducerErrors
			if perrs, ok := sendErr.(sarama.ProducerErrors); ok {
				failedSet := make(map[string]struct{}, len(perrs))
				for _, pe := range perrs {
					log.Println("Write to kafka failed:", pe)

					if pe == nil || pe.Msg == nil {
						continue
					}
					if id, ok2 := pe.Msg.Metadata.(string); ok2 {
						failedSet[id] = struct{}{}
					}
				}
				// Разносим успех/провал по спискам
				for _, m := range msgs {
					id := m.Metadata.(string)
					if _, bad := failedSet[id]; bad {
						failed = append(failed, id)
					} else {
						succeeded = append(succeeded, id)
					}
				}
				// продолжаем другие чанки; общий err вернём последним не-nil
				err = sendErr
				continue
			}

			// Фатальная ошибка всего чанка — считаем всё failed
			for _, m := range msgs {
				failed = append(failed, m.Metadata.(string))
			}
			// Возвращаем ошибку, но часть до этого могла быть отправлена успешно
			return succeeded, failed, sendErr
		}

		// Весь чанк успешен
		for _, m := range msgs {
			succeeded = append(succeeded, m.Metadata.(string))
		}
	}

	return succeeded, failed, err
}

func chunkEvents(src []*outbox.Event, size int) [][]*outbox.Event {
	return slices.Collect(slices.Chunk(src, size))
}

func ids(src []*outbox.Event) []string {
	ids := make([]string, len(src))
	for i, e := range src {
		ids[i] = e.ID
	}
	return ids
}
