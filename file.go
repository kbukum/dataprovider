package provider

import (
	"io/ioutil"
	"path/filepath"
	"net/http"
)

type FileProvider struct {
	opts interface{}
}

func (provider *FileProvider) Get(input interface{}) ([]byte, string, error) {
	path, err := filepath.Abs(input.(string))
	if err != nil {
		return nil, "", err
	}
	result, err := ioutil.ReadFile(path)
	var contentType = filepath.Ext(path)
	if contentType == "" {
		contentType = http.DetectContentType(result)
	}
	return result, contentType, err
}

func FileProviderFactory(opts ...interface{}) (Provider, error) {
	return &FileProvider{
		opts: nil,
	}, nil
}
