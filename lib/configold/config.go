package configold

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App           App           `mapstructure:"app"`
	Grpc          Grpc          `mapstructure:"grpc"`
	Postgres      Postgres      `mapstructure:"postgres"`
	KafkaProducer KafkaProducer `mapstructure:"kafka_producer"`
	KafkaConsumer KafkaConsumer `mapstructure:"kafka_consumer"`
}

func LoadConfig(ctx context.Context, fileCfg string) (*Config, error) {

	//если есть переменные окружения загрузим их для последующего приоритетного переопределения данных yaml
	_ = godotenv.Load(".env")

	v := viper.New()
	v.SetConfigType("yaml")
	if fileCfg != "" {
		v.SetConfigFile(fileCfg)
	} else {
		v.SetConfigName("config")
		v.AddConfigPath("./config/")
		v.AddConfigPath(".")
	}

	//устанавливаем дефолтные значения, если нет в конфиге
	v.SetDefault("app.name", "my-app")
	v.SetDefault("app.version", "v.0.0.0")
	v.SetDefault("app.mode", "debug")

	v.SetDefault("postgres.host", "localhost")
	v.SetDefault("postgres.port", 5432)
	v.SetDefault("postgres.user", "user")
	v.SetDefault("postgres.password", "psw")
	v.SetDefault("postgres.database", "dbname")
	v.SetDefault("postgres.sslmode", "disable")

	v.SetDefault("grpc.port", 50051)
	v.SetDefault("grpc.host", "localhost")

	v.SetDefault("kafka_producer.brokers", "localhost:9092")
	v.SetDefault("kafka_consumer.brokers", "localhost:9092")

	//переписываем переменными окружения, если они есть
	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.BindEnv("app.mode", "APP_MODE")
	v.BindEnv("app.name", "APP_NAME")
	v.BindEnv("app.version", "APP_VERSION")
	v.BindEnv("postgres.host", "POSTGRES_HOST")
	v.BindEnv("postgres.port", "POSTGRES_PORT")
	v.BindEnv("postgres.user", "POSTGRES_USER")
	v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	v.BindEnv("postgres.database", "POSTGRES_DATABASE")
	v.BindEnv("postgres.sslmode", "POSTGRES_SSLMODE")

	v.BindEnv("grpc.port", "GRPC_PORT")
	v.BindEnv("grpc.host", "GRPC_HOST")

	v.BindEnv("kafka_producer.brokers", "PRODUCER_BROKERS")
	v.BindEnv("kafka_consumer.brokers", "CONSUMER_BROKERS")

	if err := v.ReadInConfig(); err != nil {
		// файл не обязателен — можно работать только с ENV/дефолтами
		logger.Warnf(context.TODO(), "cannot read config file: %v (using env/defaults)", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	if cfg.App.Mode == "" {
		cfg.App.Mode = "debug"
	}

	return &cfg, nil
}
