package Router

import (
	"log"
	"strings"

	"Pigeon/app/Interfaces"
	"Pigeon/config"
	"github.com/gin-gonic/gin"
)

var (
	domainsMap = config.DomainConfig
	routerMode = config.RouterMode
)

func DomainRouter(c *gin.Context) {
	// 自由模式, 或者根目录
	if routerMode == config.RouterModeFreedom {
		c.Next()
		return
	}
	var (
		host, path   = c.Request.Host, c.Request.URL.Path
		hostSplit    = strings.Split(host, ".")
		hostSplitLen = len(hostSplit)
		routePath    = ""
	)
	if hostSplitLen > 1 {
		// 拼接一级域名 b.com
		host = hostSplit[hostSplitLen-2] + "." + hostSplit[hostSplitLen-1]
		subdomains, ok := domainsMap[host]
		// not found
		if !ok {
			c.Status(404)
			c.Abort()
			return
		}
		// 寻找子域名的处理函数
		if hostSplitLen > 2 {
			var (
				endIndex   = hostSplitLen - 2
				subs       = hostSplit[:endIndex]
				startMatch bool
			)
			// deep subdomain to request path
			for i := len(subs); i > 0; i-- {
				if startMatch {
					c.Request.URL.Path = "/" + subs[i-1] + c.Request.URL.Path
					continue
				}
				// not found subdomain for this layer
				subdomains, ok = subdomains.Subs[subs[i-1]]
				if !ok {
					c.Status(404)
					c.Abort()
					return
				}
				// 动态域名，拼接到 route 上
				if subdomains != nil && subdomains.Match {
					startMatch = true
				}
				routePath += "/" + subs[i-1]
			}
		}
	}
	// subdomain: a.bc.com/test -> bc.com/a/test
	// a.b.c.d.com/test -> d.com/c/b/a/test
	log.Println("routing domain", c.Request.Host+path, "to path", routePath+c.Request.URL.Path)
	c.Set("domain_route_path", routePath)
	// find handler in route path map
	if r, p := Interfaces.ResolveRouteHandler(c); r != nil {
		c.Request.URL.Path = routePath + c.Request.URL.Path
		// log.Println(c.Request.URL.Path, r)
		c.Params = p
		r.HandlerFunc(c)
	} else {
		c.Status(404)
	}
	c.Abort()
}
