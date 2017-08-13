## Data Provider
    Get Data as Byte with Data Type from anywhere which providers are defined.

#### Motivation
Always we need to get data from anywhere. In a lot of application or libraries we want to find data by the given environment and options.

#### Provider Interfaces

* Provider Factory 
```go
type ProviderFactory func(opts ...interface{}) (Provider, error)
```

* Provider 
```go
type Provider interface{
	Get(input interface{}) ([]byte, string, error)
}
```

#### Create, Setup, Build Provider

* Create(provider type)
Create is take provider type like `file` or `http`
```go
var p, err = provider.Create("http")
```

* Setup
if provider needs options for setup then you can use this method.
```go
var p, err = provider.Create("http").Setup(opts...interface{})
```
* Build
Provides to create provider.
```go
var p, err = provider.Create("http").Build()
```

#### Exist Providers `file`, `http`

* `file` provider
Provides to get data from absolute or relative file path.

```go
p, err := provider.Create("file").Build()
if err != nil {
    fmt.Println("couldn't create file provider !")
} else {
    data, ext, err := p.Get("./file.go")
    
}
```

* `http` provider

Provides to get data from absolute or relative file path.
* Standard Http Provider
```go
p, err := provider.Create("http").Build()
if err != nil {
    fmt.Println("couldn't create http provider !")
} else {
    data, ext, err := p.Get("https://www.google.com.tr/logos/2017/hiphop/cta_bg.jpg")
}
```

* Custom Http Provider
if you need to apply some security rules then
```go
// get file by http provider
	result, contentType, err = p.Get(&provider.HttpInput{
		Method: "GET",
		Url:"https://www.google.com.tr/logos/2017/hiphop/cta_bg.jpg",
		Headers: map[string]string{},
})
```
#### New Provider

##### Create You Provider

* Create Provider
```go
type MyProvider struct {
	opts interface{}
}
```
* Create `Get` method for your provider
```go
func (provider *MyProvider) Get(input interface{}) ([]byte, string, error) {
	...
	result is must be data as []byte
	contentType is extension of the data.
	return result, contentType, err
}
```
* Create Provider Factory Function.
```go
func MyProviderFactory(opts ...interface{}) (Provider, error) {
	return &MyProvider{
		opts: nil,
	}, nil
}
```
#### Add Your Provider to Data Provider

```go
import (
	"github.com/kbukum/dataprovider"
)
func main() {
    provider.Add("myProvider", MyProviderFactory)
}
```


