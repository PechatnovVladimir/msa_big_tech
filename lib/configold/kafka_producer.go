package configold

type KafkaProducer struct {
	Brokers string `mapstructure:"brokers"`
}
