package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"Pigeon/library/utils"
	json "github.com/json-iterator/go"
)

var DefaultClient = http.DefaultClient

// NewClient create new http client
func NewClient() *http.Client {
	return &http.Client{}
}

// Get config：*Config or request url
// client：*http.Client or nil (using default client)

func GetByURL(url string, client *http.Client) (*http.Response, error) {
	return GetByConfig(&Config{URL: url}, client)
}

func GetByConfig(config *Config, client *http.Client) (*http.Response, error) {
	if config == nil {
		return nil, errors.New("GetByConfig: config cannot be nil")
	}
	return get(config, client)
}

// Get returns string
func Get(config interface{}, client *http.Client) ([]byte, error) {
	rsp, err := get(parseToRequestConfig(config), client)
	if err != nil {
		return []byte{}, err
	}
	return ReadResponseBody(rsp.Body)
}

func get(config *Config, client *http.Client) (*http.Response, error) {
	config.Method = "GET"
	return Request(config, client)
}

func HeadByURL(url string, client *http.Client) (*http.Response, error) {
	return HeadByConfig(&Config{URL: url}, client)
}

func HeadByConfig(config *Config, client *http.Client) (*http.Response, error) {
	if config == nil {
		return nil, errors.New("HeadByConfig: config cannot be nil")
	}
	return head(config, client)
}

// Head returns map[string][]string
func Head(config interface{}, client *http.Client) http.Header {
	h, _ := head(parseToRequestConfig(config), client)
	return h.Header
}

func head(config *Config, client *http.Client) (*http.Response, error) {
	config.Method = "HEAD"
	return Request(config, client)
}

func PostByURL(url string, client *http.Client) (*http.Response, error) {
	return PostByConfig(&Config{URL: url}, client)
}

func PostByConfig(config *Config, client *http.Client) (*http.Response, error) {
	if config == nil {
		return nil, errors.New("PostByConfig: config cannot be nil")
	}
	return post(config, client)
}

// Post returns []byte
func Post(config interface{}, client *http.Client) ([]byte, error) {
	rsp, err := post(parseToRequestConfig(config), client)
	if err != nil {
		return []byte{}, err
	}
	return ReadResponseBody(rsp.Body)
}

func post(config *Config, client *http.Client) (*http.Response, error) {
	config.Method = "POST"
	if config.Headers == nil {
		config.Headers = make(HeaderMap, 10)
	}
	// check content-type
	// is not set content-type
	if _, ok := config.Headers[`Content-Type`]; !ok {
		config.Headers[`Content-Type`] = `application/x-www-form-urlencoded`
	}
	var data interface{}
	switch config.Data.(type) {
	// has file
	case *Files:
		if f := config.Data.(*Files); f.CountFile() > 0 {
			go f.PipeFile()()
			data = f.GetPipeReader()
			defer data.(*io.PipeReader).Close()
			config.Headers[`Content-Type`] = f.GetWriter().FormDataContentType()
		}
	default:
		data = parseRequestData(config.Data)
	}
	config.Data = data
	return Request(config, client)
}

// func Delete(config Config) *http.Response  {}
// func Put(config Config) *http.Response     {}
// func Options(config Config) *http.Response {}

// Request send base request
func Request(config *Config, client *http.Client) (*http.Response, error) {
	requestConfig := parseToRequestConfig(config)
	var body io.Reader
	if data, ok := requestConfig.Data.(io.Reader); ok {
		body = data
	}
	rsp, _ := http.NewRequest(requestConfig.Method, requestConfig.URL, body)
	return SendHttpRequest(rsp, client, requestConfig)
}

func SendHttpRequest(request *http.Request, client *http.Client, config *Config) (*http.Response, error) {
	if request == nil {
		return nil, errors.New(`request must be not nil`)
	}
	// create a new http client
	if client == nil {
		client = NewClient()
	}
	// apply configs to request
	if config != nil {
		config.ApplyToHttpClient(client)
		config.Headers.ApplyToRequest(request)
	}
	// send request by http.Client{} method.
	// http.DefaultClient is http.Client instance
	return client.Do(request)
}

func marshalRequestDataValue(v reflect.Value) string {
	vt := v.Kind().String()
	// 数字就直接转换，否则直接 json
	if strings.HasPrefix(vt, "int") || strings.HasPrefix(vt, "string") {
		return fmt.Sprintf("%v", v)
	}
	r, _ := json.Marshal(v.Interface())
	return string(r)
}

func parseRequestData(reqData interface{}) io.Reader {
	data := url.Values{}
	reqDataType := reflect.TypeOf(reqData)
	// 匹配到是 map 类型,把值转换成字符串
	if strings.HasPrefix(reqDataType.String(), "map[string]") {
		reqDataValue := reflect.ValueOf(reqData)
		keys := reqDataValue.MapKeys()
		for _, k := range keys {
			data.Add(k.String(), marshalRequestDataValue(reqDataValue.MapIndex(k)))
		}
	} else {
		switch requestData := reqData.(type) {
		case string:
			data, _ = url.ParseQuery(requestData)
		case *Files:
			return requestData.Encode()
		}
	}
	return strings.NewReader(data.Encode())
}

func parseToRequestConfig(config interface{}) *Config {
	switch v := config.(type) {
	case string:
		return &Config{URL: v}
	case Config:
		return &v
	case *Config:
		return v
	}
	return &Config{}
}

func ReadResponseBody(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

func ReadResponseBodyString(body io.ReadCloser) (string, error) {
	// read response
	rsp, err := ReadResponseBody(body)
	return utils.BytesToString(rsp), err
}
