package configsecrets

type KafkaConsumer struct {
	Brokers string `mapstructure:"brokers"`
}
