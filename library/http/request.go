package http

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"Pigeon/library/utils"
)

var DefaultClient = http.DefaultClient

// NewClient create new http client
func NewClient() *http.Client {
	return &http.Client{}
}

// Get config：*Config or request url
// client：*http.Client or nil (using default client)
func Get(config interface{}, client *http.Client) (string, error) {
	rsp, err := Request(config, client)
	if err != nil {
		return ``, err
	}
	return readResponseBody(rsp.Body)
}

func Head(config interface{}, client *http.Client) http.Header {
	h, _ := head(parseToRequestConfig(config), client)
	return h.Header
}

func head(config *Config, client *http.Client) (*http.Response, error) {
	config.Method = "HEAD"
	return Request(config, client)
}

func Post(config interface{}, client *http.Client) (string, error) {
	rsp, err := post(parseToRequestConfig(config), client)
	if err != nil {
		return ``, err
	}
	return readResponseBody(rsp.Body)
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
func Request(config interface{}, client *http.Client) (*http.Response, error) {
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

func parseRequestData(reqData interface{}) io.Reader {
	data := url.Values{}
	switch requestData := reqData.(type) {
	case map[string]string:
		for k, v := range requestData {
			data.Add(k, v)
		}
	case string:
		data, _ = url.ParseQuery(requestData)
	case *Files:
		return requestData.Encode()
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

func ReadResponseBody(rsp *http.Response) string {
	body, err := readResponseBody(rsp.Body)
	if err != nil {
		return ""
	}
	return body
}

func readResponseBody(body io.ReadCloser) (string, error) {
	// read response
	rsp, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	return utils.BytesToString(rsp), nil
}
