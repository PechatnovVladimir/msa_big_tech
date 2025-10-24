package secrets

type VaultProvider struct {
	isSet bool
}

func (v *VaultProvider) GetBytes(key string) ([]byte, error) {
	return nil, nil
}

func (v *VaultProvider) Get(key string) (string, error) {
	return "", nil
}

func (v *VaultProvider) IsSet() bool {
	return v.isSet
}

func NewVaultProvider() *VaultProvider {
	return &VaultProvider{
		isSet: true,
	}
}
