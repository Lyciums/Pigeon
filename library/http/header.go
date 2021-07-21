package http

import (
	"net/http"
	"regexp"
)

type HeaderMap map[string]string

func (hm HeaderMap) GetOrDefault(name, defaultValue string) string {
	if v, ok := hm[name]; ok {
		return v
	}
	return defaultValue
}

func (hm HeaderMap) Get(name string) string {
	return hm.GetOrDefault(name, "")
}

func (hm HeaderMap) GetMultiple(names ...string) HeaderMap {
	m := NewHeaderMap(len(names))
	for _, name := range names {
		m[name] = hm.Get(name)
	}
	return m
}

func (hm HeaderMap) Set(name, value string) HeaderMap {
	hm[name] = value
	return hm
}

func (hm HeaderMap) Delete(names ...string) HeaderMap {
	for _, name := range names {
		delete(hm, name)
	}
	return hm
}

func (hm HeaderMap) ApplyToRequest(req *http.Request) {
	if hm != nil {
		for key, value := range hm {
			req.Header.Set(key, value)
		}
	}
}

func NewHeaderMap(size int) HeaderMap {
	return make(HeaderMap, size)
}

func NewHeaderMapBySlice(values [][]string) HeaderMap {
	hm := NewHeaderMap(len(values))
	for _, v := range values {
		if len(v) > 1 {
			// 适配 MatchCURLHeadersRegexp 的值
			v = v[len(v)-2:]
			hm[v[0]] = v[1]
		}
	}
	return hm
}

var (
	MatchCURLHeadersRegexp = regexp.MustCompile(`-H '([^:]+):\s+([^']*?)'`)
)

func ParseHeaderByCURL(curl string) HeaderMap {
	return NewHeaderMapBySlice(MatchCURLHeadersRegexp.FindAllStringSubmatch(curl, -1))
}
