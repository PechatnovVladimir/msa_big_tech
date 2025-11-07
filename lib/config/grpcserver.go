package config

type GrpcServerConfig struct {
	Host      string          `mapstructure:"host"`
	Port      int             `mapstructure:"port"`
	Timeout   TimeoutConfig   `mapstructure:"timeout"`
	RateLimit RateLimitConfig `mapstructure:"rateLimit"`
}

type TimeoutConfig struct {
	Enabled   bool          `mapstructure:"enabled"`
	Ignore    []string      `mapstructure:"ignore"`
	TimeoutMs int           `mapstructure:"timeoutMs"`
	Paths     []PathTimeout `mapstructure:"paths"`
}

type PathTimeout struct {
	Path      string `mapstructure:"path"`
	TimeoutMs int    `mapstructure:"timeoutMs"`
}

type RateLimitConfig struct {
	Enabled   bool            `mapstructure:"enabled"`
	Ignore    []string        `mapstructure:"ignore"`
	ReqPerSec int             `mapstructure:"reqPerSec"`
	Paths     []PathRateLimit `mapstructure:"paths"`
}

type PathRateLimit struct {
	Path      string `mapstructure:"path"`
	ReqPerSec int    `mapstructure:"reqPerSec"`
}
