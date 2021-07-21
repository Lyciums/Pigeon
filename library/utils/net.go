package utils

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

const (
	// DefaultCheckTimeout 默认 4 秒
	DefaultCheckTimeout = time.Second << 2
)

// GetHostIPAddr 根据域名获取 ip 地址
func GetHostIPAddr(host string) string {
	ip, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return ""
	}
	return ip.String()
}

// CheckPortIsOpen 检查端口是否开放
func CheckPortIsOpen(addr string, port uint16, timeout time.Duration) bool {
	if addr == "" {
		return false
	}
	addrSplit := strings.Split(addr, "//")
	return CheckAddrPortIsOpen(addrSplit[len(addrSplit)-1], port, timeout)
}

func CheckAddrPortIsOpen(addr string, port uint16, timeout time.Duration) bool {
	if addr == "" {
		return false
	}
	if 1 > timeout {
		timeout = DefaultCheckTimeout
	}
	addr = fmt.Sprintf("%s:%d", addr, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func CheckHostAvailable(u string) bool {
	p, _ := url.Parse(u)
	if p.Host == "" {
		return false
	}
	var port uint16 = 80
	if p.Scheme == "https" {
		port = 443
	}
	return CheckPortIsOpen(p.Host, port, 0)
}
