package consumer

import (
	"github.com/IBM/sarama"
	"time"
)

type Config struct {
	Brokers          []string
	Topics           []string
	GroupID          string
	SaramaConfig     *sarama.Config
	BatchSize        int
	BatchTimeout     time.Duration
	FlushOnRebalance bool
	ConsumerID       string //Уникальный ID инстанса
}
