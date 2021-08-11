package http

import (
	"crypto/tls"
	"net/http"
	"time"
	"unsafe"
)

type Config struct {
	URL        string
	Method     string
	Data       interface{}
	Headers    HeaderMap
	Timeout    int
	Payload    bool
	SkipVerify bool
}

func (cfg *Config) ApplyToHttpClient(client *http.Client) *http.Client {
	client.Timeout = time.Duration(cfg.Timeout)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.SkipVerify},
	}
	return client
}

func (cfg *Config) Copy() Config {
	return *(*Config)(unsafe.Pointer(cfg))
}

func (cfg *Config) ClearHeader() *Config {
	cfg.Headers = NewHeaderMap(10)
	return cfg
}

func (cfg *Config) SetHeaders(keyValues ...string) *Config {
	lens := len(keyValues)
	if cfg.Headers == nil {
		cfg.Headers = NewHeaderMap(10)
	}
	// 补参数值
	if lens&1 > 0 {
		keyValues = append(keyValues, "")
	}
	for lens > 0 {
		cfg.Headers[keyValues[lens-2]] = keyValues[lens-1]
		lens -= 2
	}
	return cfg
}

func (cfg *Config) GetHeader() HeaderMap {
	return cfg.Headers
}