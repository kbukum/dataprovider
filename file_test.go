package provider_test

import (
	"testing"
	"github.com/kbukum/dataprovider"
	"github.com/stretchr/testify/assert"
)

func TestFileProvider_Get(t *testing.T) {
	var m = map[string]interface{}{}
	p, err := provider.FileProviderFactory(m)
	// create provider
	assert.Nil(t, err, "Couldn't create provider Error: %v", err)
	result, contentType, err := p.Get("./file.go")
	assert.Nil(t, err, "Error is not null")
	assert.Equal(t, ".go", contentType, "ContentType must be .jpg")
	assert.NotNil(t, result, "Content is nil ")
}
