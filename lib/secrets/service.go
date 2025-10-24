package secrets

type SecretsProvider interface {
	Get(key string) (string, error)      // строки (пароли, токены)
	GetBytes(key string) ([]byte, error) // бинарь (TLS ключи/серты)
	IsSet() bool
}

type Secrets struct {
	SecretsProvider
}

func NewSecrets(secretsProvider SecretsProvider) *Secrets {
	return &Secrets{secretsProvider}
}
