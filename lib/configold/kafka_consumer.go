package configold

type KafkaConsumer struct {
	Brokers string `mapstructure:"brokers"`
}
