package theme

import (
	"context"
)

type ComponentLoader interface {
	Load(ctx context.Context, data interface{}) ([]byte, error)
}

type componentLoader struct {
	loader    Loader
	cache     []byte
	hash      interface{}
	category  string
	component string
}

func NewComponentLoader(loader Loader, component string) ComponentLoader {
	return &componentLoader{
		loader:    loader,
		cache:     []byte{},
		hash:      0,
		component: component,
	}
}

func (s *componentLoader) Load(ctx context.Context, data interface{}) ([]byte, error) {
	if s.hash != data {
		result, err := s.loader.LoadComponent(ctx, s.component, data)
		if err != nil {
			return nil, err
		}
		s.cache = result
		s.hash = data
	}
	return s.cache, nil
}
