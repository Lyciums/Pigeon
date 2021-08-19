package utils

import (
	"bytes"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func LoadFromString(s string) (*html.Node, error) {
	return htmlquery.Parse(strings.NewReader(s))
}

func LoadFromBytes(b []byte) (*html.Node, error) {
	return htmlquery.Parse(bytes.NewReader(b))
}
