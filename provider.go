package provider

import (
	"github.com/pkg/errors"
	"fmt"
)

type Provider interface {
	Get(input interface{}) ([]byte, string, error)
}

type ProviderFactory func(opts ...interface{}) (Provider, error)

var factories = map[string]ProviderFactory{
	"file": FileProviderFactory,
	"http": HttpProviderFactory,
}

type factory struct {
	Kind    string
	Factory ProviderFactory
	Opts    []interface{}
}

func Create(kind string) *factory {
	creator, err := get(kind)
	if err != nil {
		return &factory{
			Kind: kind,
		}
	}
	return &factory{
		Factory: creator,
	}
}

func (factory *factory) Setup(opts ...interface{}) *factory {
	factory.Opts = opts
	return factory
}

func (factory *factory) Build() (Provider, error) {
	if factory.Factory == nil {
		return nil, errors.New(factory.Kind + " provider providers not found !")
	}
	return factory.Factory(factory.Opts)
}

func Add(kind string, creator ProviderFactory) error {
	if kind == "" {
		return errors.New("kind is empty or null !")
	}
	factories[kind] = creator
	return nil
}

func Has(kind string) bool {
	_, err := get(kind)
	return err == nil
}

func get(kind string) (ProviderFactory, error) {
	for k, v := range factories {
		if kind == k {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("%s provider not found in providers", kind))
}
