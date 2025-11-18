package config

import "time"

type GrpcClientConfig struct {
	Timeout        time.Duration `mapstructure:"timeout"`
	Retry          RetryConfig   `mapstructure:"retry"`
	CircuitBreaker CBConfig      `mapstructure:"circuitBreaker"`
	Metrics        bool          `mapstructure:"metrics"`
}

type RetryConfig struct {
	MaxAttempts    uint          `mapstructure:"maxAttempts"`
	Backoff        BackoffConfig `mapstructure:"backoff"`
	RetryableCodes []string      `mapstructure:"retryableCodes"`
}

type BackoffConfig struct {
	Base           time.Duration `mapstructure:"base"`
	Max            time.Duration `mapstructure:"max"`
	Jitter         bool          `mapstructure:"jitter"`
	JitterFraction float64       `mapstructure:"jitterFraction"`
}

type CBConfig struct {
	FailuresForOpen  uint32        `mapstructure:"failuresForOpen"`
	Window           time.Duration `mapstructure:"window"`
	HalfOpenMaxCalls uint32        `mapstructure:"halfOpenMaxCalls"`
	OpenStateFor     time.Duration `mapstructure:"openStateFor"`
}
