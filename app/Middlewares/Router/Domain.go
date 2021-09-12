package Router

import (
	"log"
	"strings"

	"Marvel/app/Caches/App"
	"Marvel/app/Caches/Router"
	"Marvel/app/Interfaces"
	"Marvel/config"
	"github.com/gin-gonic/gin"
)

var (
	domainsMap = config.DomainConfig
	routerMode = config.RouterMode
)

func DomainRouter(c *gin.Context) {
	// 自由模式
	if routerMode == config.RouterModeFreedom {
		c.Next()
		return
	}
	var (
		host         = c.Request.Host
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
				endIndex = hostSplitLen - 2
				subs     = hostSplit[:endIndex]
			)
			// deep subdomain to request path
			for i := len(subs); i > 0; i-- {
				// not found subdomain for this layer
				if subdomains, ok = subdomains.Subs[subs[i-1]]; !ok {
					c.Status(404)
					c.Abort()
					return
				}
				routePath += "/" + subs[i-1]
			}
		}
	}
	// map root router
	if Router.RootRouterMap == nil {
		Router.RootRouterMap = Interfaces.MakeRouterMap(App.Engine.Routes())
	}
	// subdomain: a.bc.com/test -> bc.com/a/test
	// a.b.c.d.com/test -> d.com/c/b/a/test
	routePath += c.Request.URL.Path
	log.Println("routing domain", c.Request.Host, "to path", routePath)
	// find handler in route path map
	if h, ok := Router.RootRouterMap[routePath]; ok {
		h.HandlerFunc(c)
	} else {
		c.Status(404)
	}
	c.Abort()
}
