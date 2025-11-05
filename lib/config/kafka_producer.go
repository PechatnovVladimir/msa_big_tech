package config

type KafkaProducer struct {
	Brokers string `mapstructure:"brokers"`
}
