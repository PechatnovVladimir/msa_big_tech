package secrets

import (
	"errors"
	vault "github.com/hashicorp/vault/api"
	"log"
	"os"
)

type VaultProvider struct {
	data  map[string]interface{}
	isSet bool
}

func (v *VaultProvider) GetBytes(key string) ([]byte, error) {
	return nil, nil
}

func (v *VaultProvider) Get(key string) (string, error) {
	value, ok := v.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value.(string), nil
}

func (v *VaultProvider) IsSet() bool {
	return v.isSet
}

func NewVaultProvider(addr string, path string) *VaultProvider {
	cfg := vault.DefaultConfig()
	cfg.Address = getenv("VAULT_ADDR", addr)
	client, err := vault.NewClient(cfg)
	if err != nil {
		log.Printf("WARN: Failed to create vault client: %v", err)
		return &VaultProvider{
			isSet: false,
		}
	}
	client.SetToken(getenv("VAULT_TOKEN", "root"))

	secretPath := getenv("VAULT_SECRET_PATH", path)

	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Printf("WARN: Failed to read secret from vault: %v", err)
		return &VaultProvider{
			isSet: false,
		}
	}

	if secret == nil || secret.Data == nil {
		return &VaultProvider{
			isSet: false,
		}
	}

	row, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return &VaultProvider{
			isSet: false,
		}
	}

	return &VaultProvider{
		data:  row,
		isSet: true,
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
