package configsecrets

import "fmt"

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

func (p *Postgres) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.Database, p.SSLMode,
	)
}
