package configsecrets

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Mode    string `mapstructure:"mode"`
}
