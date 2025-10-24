package configsecrets

type KafkaProducer struct {
	Brokers string `mapstructure:"brokers"`
}
