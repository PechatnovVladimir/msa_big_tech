package secrets

import (
	vault "github.com/hashicorp/vault/api"
)

type Secrets struct {
	DBUser string
	DBPass string
}

func LoadSecrets() (*Secrets, error) {
	cfg := vault.DefaultConfig()
	_ = cfg
	return nil, nil
}
