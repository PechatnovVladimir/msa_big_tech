package config

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/PechatnovVladimir/msa_big_tech/lib/secrets"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	App           App           `mapstructure:"app"`
	Grpc          Grpc          `mapstructure:"grpc"`
	Postgres      Postgres      `mapstructure:"postgres"`
	KafkaProducer KafkaProducer `mapstructure:"kafka_producer"`
	KafkaConsumer KafkaConsumer `mapstructure:"kafka_consumer"`
	Logger        Logger        `mapstructure:"logger"`
}

type ConfigLoader struct {
	viper           *viper.Viper
	secretsProvider secrets.SecretsProvider //Провайдер секретов
	secretKeys      []string                //Список ключей-секретов
	envBindings     map[string]string       //Привязка ключей к переменным окружения
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{
		viper:       viper.New(),
		secretKeys:  []string{},
		envBindings: make(map[string]string),
	}
}

// WithSecretsProvider устанавливает провайдер секретов
func (cl *ConfigLoader) WithSecretsProvider(provider secrets.SecretsProvider) *ConfigLoader {
	cl.secretsProvider = provider
	return cl
}

// SetDefault устанавливает значение по умолчанию для ключа
func (cl *ConfigLoader) SetDefault(key string, value interface{}) *ConfigLoader {
	cl.viper.SetDefault(key, value)
	return cl
}

// SetDefaults устанавливает несколько значений по умолчанию
func (cl *ConfigLoader) SetDefaults(defaults map[string]interface{}) *ConfigLoader {
	for key, value := range defaults {
		cl.viper.SetDefault(key, value)
	}
	return cl
}

// BindEnv привязывает ключ конфигурации к переменной окружения
func (cl *ConfigLoader) BindEnv(key string, envVar string) *ConfigLoader {
	err := cl.viper.BindEnv(key, envVar)
	if err != nil {
		return nil
	}
	return cl
}

// BindEnvPrefix привязывает все ключи к переменным окружения с префиксом
func (cl *ConfigLoader) BindEnvPrefix(prefix string, keys ...string) *ConfigLoader {
	for _, key := range keys {
		envVar := fmt.Sprintf("%s_%s", strings.ToUpper(prefix), strings.ToUpper(key))
		cl.BindEnv(key, envVar)
	}
	return cl
}

// AddSecretKey добавляет ключ, который должен быть загружен как секрет
func (cl *ConfigLoader) AddSecretKey(key string) *ConfigLoader {
	cl.secretKeys = append(cl.secretKeys, key)
	return cl
}

// AddSecretKeys добавляет несколько ключей для загрузки как секреты
func (cl *ConfigLoader) AddSecretKeys(keys ...string) *ConfigLoader {
	cl.secretKeys = append(cl.secretKeys, keys...)
	return cl
}

func (cl *ConfigLoader) LoadConfig(fileCfg string) (*Config, error) {

	//если есть переменные окружения загрузим их для последующего приоритетного переопределения данных yaml
	_ = godotenv.Load(".env")

	cl.viper.SetConfigType("yaml")
	if fileCfg != "" {
		cl.viper.SetConfigFile(fileCfg)
	} else {
		cl.viper.SetConfigName("config")
		cl.viper.AddConfigPath("./config/")
		cl.viper.AddConfigPath(".")
	}

	//переписываем переменными окружения, если они есть
	cl.viper.AutomaticEnv()

	if err := cl.viper.ReadInConfig(); err != nil {
		// файл не обязателен — можно работать только с ENV/дефолтами
		logger.Warnf(context.TODO(), "cannot read config file: %v (using env/defaults)", err)
	}

	if cl.secretsProvider != nil {
		if cl.secretsProvider.IsSet() {
			cl.loadSecrets()
		}
	}

	var cfg Config
	if err := cl.viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

// loadSecrets загружает секреты из провайдера
func (cl *ConfigLoader) loadSecrets() {
	for _, key := range cl.secretKeys {
		//Секреты в приоритете, поэтому перепишем, даже если уже установленное есть
		// Проверяем, нужно ли загружать секрет (значение не установлено или пустое)
		if !cl.viper.IsSet(key) || cl.viper.GetString(key) == "" {
			secretValue, err := cl.secretsProvider.Get(key)
			if err != nil {
				logger.Warnf(context.TODO(), "secret not found for key %s: %v\n", key, err)
				continue
			}

			// Устанавливаем значение в Viper
			if secretValue != "" {
				cl.viper.Set(key, secretValue)
			}
		}
	}
	return
}

// Get возвращает значение конфигурации
func (cl *ConfigLoader) Get(key string) interface{} {
	return cl.viper.Get(key)
}

// GetString возвращает строковое значение конфигурации
func (cl *ConfigLoader) GetString(key string) string {
	return cl.viper.GetString(key)
}

// GetInt возвращает целочисленное значение конфигурации
func (cl *ConfigLoader) GetInt(key string) int {
	return cl.viper.GetInt(key)
}

// GetBool возвращает булево значение конфигурации
func (cl *ConfigLoader) GetBool(key string) bool {
	return cl.viper.GetBool(key)
}

// GetStringMap возвращает map[string]interface{}
func (cl *ConfigLoader) GetStringMap(key string) map[string]interface{} {
	return cl.viper.GetStringMap(key)
}

// GetStringSlice возвращает []string
func (cl *ConfigLoader) GetStringSlice(key string) []string {
	return cl.viper.GetStringSlice(key)
}

// IsSet проверяет, установлено ли значение для ключа
func (cl *ConfigLoader) IsSet(key string) bool {
	return cl.viper.IsSet(key)
}

// AllSettings возвращает все настройки
func (cl *ConfigLoader) AllSettings() map[string]interface{} {
	return cl.viper.AllSettings()
}

// GetEnvBindings возвращает привязки к переменным окружения
func (cl *ConfigLoader) GetEnvBindings() map[string]string {
	return cl.envBindings
}

// Configure полная настройка загрузчика одним вызовом
func (cl *ConfigLoader) Configure(options ...ConfigOption) *ConfigLoader {
	for _, option := range options {
		option(cl)
	}
	return cl
}

// ConfigOption тип функции для конфигурации
type ConfigOption func(*ConfigLoader)

// WithDefaults опция для установки значений по умолчанию
func WithDefaults(defaults map[string]interface{}) ConfigOption {
	return func(cl *ConfigLoader) {
		cl.SetDefaults(defaults)
	}
}

// WithEnvBindings опция для привязки переменных окружения
func WithEnvBindings(bindings map[string]string) ConfigOption {
	return func(cl *ConfigLoader) {
		for key, envVar := range bindings {
			cl.BindEnv(key, envVar)
		}
	}
}
