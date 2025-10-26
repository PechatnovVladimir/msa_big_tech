package secrets

import "errors"

type CompositeProvider struct {
	providers []SecretsProvider
	isSet     bool
}

func (c *CompositeProvider) Get(key string) (string, error) {
	for _, provider := range c.providers {
		if provider.IsSet() {
			value, err := provider.Get(key)
			if err == nil {
				return value, nil
			}
		}
	}
	return "", errors.New("key not found in any provider")
}

func (c *CompositeProvider) GetBytes(key string) ([]byte, error) {
	return nil, nil
}

func (c *CompositeProvider) IsSet() bool {
	return c.isSet
}

func NewCompositeProvider(vaultAddr string, vaultPath string, filePath string) *CompositeProvider {
	providers := []SecretsProvider{
		NewEnvProvider(),
		NewFileProvider(filePath),
		NewVaultProvider(vaultAddr, vaultPath),
	}

	isset := false
	for _, provider := range providers {
		if provider.IsSet() {
			isset = true
			break
		}
	}

	return &CompositeProvider{
		providers: providers,
		isSet:     isset,
	}
}
