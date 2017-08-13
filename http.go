package provider

import (
	"io/ioutil"
	"net/http"
	"mime"
	"strings"
	"github.com/pkg/errors"
	"path/filepath"
)

type HttpProvider struct {
	opts interface{}
}

type HttpInput struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}

func getOrDefault(value string, def string) string {
	if value == "" {
		return def
	}
	return value
}

func (provider *HttpProvider) Get(input interface{}) ([]byte, string, error) {
	client := &http.Client{}
	var req *http.Request
	var err error
	var contentType string
	switch t := input.(type) {
	case string:
		// t is string
		req, err = http.NewRequest("GET", t, nil)
		contentType = filepath.Ext(t)
	case *HttpInput:
		// t is HttpInput
		req, err = http.NewRequest(getOrDefault(t.Method, "GET"), t.Url, strings.NewReader(t.Body))
		filepath.Ext(t.Url)
		if t.Headers != nil && len(t.Headers) > 0 {
			for key, element := range t.Headers {
				req.Header.Set(key, element)
			}
		}
		contentType = filepath.Ext(t.Url)
	default:
		err = errors.New("unknown type")
	}
	if err != nil {
		return nil, "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	if contentType == "" {
		header := resp.Header.Get("content-type")
		contentType = strings.Split(header, ";")[0]
		if contentType == "" {
			contentType = http.DetectContentType(result)
		}
		extensions, err := mime.ExtensionsByType(contentType)
		if err != nil && len(extensions) > 0 {
			contentType = extensions[0]
		}
	}
	return result, contentType, nil
}

func HttpProviderFactory(opts ...interface{}) (Provider, error) {
	return &HttpProvider{
		opts: nil,
	}, nil
}
