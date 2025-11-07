package config

type Grpc struct {
	Server GrpcServerConfig `mapstructure:"server"`
	Client GrpcClientConfig `mapstructure:"client"`
}
