package provider_test

import (
	"testing"
	"github.com/kbukum/dataprovider"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	p := provider.Create("file")
	assert.NotNil(t, p, "Couldn't create Provider Factory !")
	assert.NotNil(t, p.Factory, "Couldn't create Provider Factory Function !")
	assert.Empty(t, p.Kind, "Kind is not empty.")
}

func TestFactory_Build(t *testing.T) {
	var p, err = provider.Create("file").Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
	p, err = provider.Create("http").Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
}

func TestFactory_Setup(t *testing.T) {
	var p, err = provider.Create("http").Setup().Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
	p, err = provider.Create("http").Setup(map[string]string{}).Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
}

func Test_Get(t *testing.T) {
	p, err := provider.Create("file").Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
	data, ext, err := p.Get("/Users/kamilbukum/go/src/github.com/kbukum/dataprovider/file.go")
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.Equal(t, ".go", ext, "File extension must be '.go' !")
	assert.NotNil(t, data, "Got data must be not nil !", err)
}

type MyProvider struct {
	opts interface{}
}

func (provider *MyProvider) Get(input interface{}) ([]byte, string, error) {
	return []byte(" Content: " + input.(string)), ".txt", nil
}

func MyProviderFactory(opts ...interface{}) (provider.Provider, error) {
	return &MyProvider{
		opts: nil,
	}, nil
}

func TestAdd(t *testing.T) {
	provider.Add("txt", MyProviderFactory)
	p, err := provider.Create("txt").Build()
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.NotNil(t, p, "Couldn't create file provider factory !", err)
	data, ext, err := p.Get("This is string content")
	assert.Nil(t, err, "Couldn't create file provider factory Error : %v", err)
	assert.Equal(t, ".txt", ext, "File extension must be '.go' !")
	assert.NotNil(t, data, "Got data must be not nil !", err)
}

func TestHas(t *testing.T) {
	result := provider.Has("file")
	assert.True(t, result, "Provider is not exist");
}
