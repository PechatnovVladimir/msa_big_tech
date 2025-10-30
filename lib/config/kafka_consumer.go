package config

type KafkaConsumer struct {
	Brokers string `mapstructure:"brokers"`
}
