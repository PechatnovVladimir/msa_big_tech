package secrets

import (
	"context"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type FileProvider struct {
	data map[string]interface{}
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

func NewFileProvider(file string) *FileProvider {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return nil
	}
	return &FileProvider{
		data: m,
	}
}
