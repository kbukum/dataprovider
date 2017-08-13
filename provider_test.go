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

func TestAdd(t *testing.T) {
	provider.Add("myProvider", )
}
