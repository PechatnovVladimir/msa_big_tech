package boot

import "github.com/PechatnovVladimir/msa_big_tech/lib/secrets"

type Config struct {
	ConfigFile     string
	DefaultValue   map[string]interface{}
	EnvBinding     map[string]string
	SecretKeys     []string
	SecretProvider secrets.SecretsProvider
}
