package secrets

import (
	"context"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type FileProvider struct {
	isSet bool
	data  map[string]interface{}
}

func (f *FileProvider) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (f *FileProvider) Get(ctx context.Context, key string) (string, error) {
	value, ok := f.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value.(string), nil
}

func (f *FileProvider) IsSet(ctx context.Context) bool {
	return f.isSet
}

func NewFileProvider(file string) *FileProvider {
	data, err := os.ReadFile(file)
	if err != nil {
		return &FileProvider{isSet: false}
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return &FileProvider{isSet: false}
	}
	return &FileProvider{
		data:  m,
		isSet: true,
	}
}
