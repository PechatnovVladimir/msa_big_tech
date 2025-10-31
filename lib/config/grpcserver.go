package config

type GrpcServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
