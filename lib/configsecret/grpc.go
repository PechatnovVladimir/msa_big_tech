package configsecrets

type Grpc struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
