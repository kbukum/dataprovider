package provider_test

import (
	"testing"
	"github.com/kbukum/dataprovider"
	"github.com/stretchr/testify/assert"
)

func TestHttpProvider_Get(t *testing.T) {
	p, err := provider.HttpProviderFactory()
	// create provider
	assert.Nil(t, err, "Couldn't create provider Error: %v", err)
	// get file by http provider
	result, contentType, err := p.Get("https://www.google.com.tr/logos/2017/hiphop/cta_bg.jpg")
	assert.Nil(t, err, "Error is not null")
	assert.Equal(t, ".jpg", contentType, "ContentType must be .jpg")
	assert.NotNil(t, result, "Contet is nil ")

	// get file by http provider
	result, contentType, err = p.Get(&provider.HttpInput{
		Method:  "GET",
		Url:     "https://www.google.com.tr/logos/2017/hiphop/cta_bg.jpg",
		Headers: map[string]string{},
	})
	assert.Nil(t, err, "Error is not null")
	assert.Equal(t, ".jpg", contentType, "ContentType must be .jpg")
	assert.NotNil(t, result, "Contet is nil ")

}
