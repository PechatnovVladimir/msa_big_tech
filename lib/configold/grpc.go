package configold

type Grpc struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
