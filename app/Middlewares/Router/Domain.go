package Router

import (
	"log"
	"strings"

	"Pigeon/app/Caches/App"
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
	if r, p := ResolveRouteHandler(c); r != nil {
		c.Request.URL.Path = routePath + c.Request.URL.Path
		// log.Println(c.Request.URL.Path, r)
		c.Params = p
		r.HandlerFunc(c)
	} else {
		c.Status(404)
	}
	c.Abort()
}

func ResolveRouteHandler(c *gin.Context) (r *gin.RouteInfo, p gin.Params) {
	var (
		prefix        = c.MustGet("domain_route_path").(string)
		suffix        = prefix + c.Request.URL.Path
		rawSuffix     = strings.Trim(suffix, "/")
		suffixSplit   = strings.Split(rawSuffix, "/")
		suffixPathLen = len(suffixSplit)
		method        = c.Request.Method
	)
	// 匹配前缀一致的路由
	for _, route := range App.Engine.Routes() {
		p = p[:0]
		// 匹配子域名部分
		rawPath := strings.Trim(route.Path, "/")
		// 匹配方法一致，前缀一致
		if route.Method == method && (isParamPath(rawPath) || strings.HasPrefix(route.Path, prefix)) {
			// 路径为空
			if rawPath == "" || rawPath == "" && rawSuffix != "" {
				continue
			}
			// 删除左右两侧的 / 做对比
			routePathSplit := strings.Split(rawPath, "/")
			psl := len(routePathSplit)
			// 没有路径
			if psl == 0 && suffixPathLen == 0 {
				return &route, p
			}
			// 路由参数个数大于请求参数
			// 请求参数个数大于路由参数，且路由参数当中不包含路径匹配符 *
			// 没有路由参数，但有请求参数
			if psl > suffixPathLen || suffixPathLen > psl && !strings.Contains(rawPath, "*") || psl == 0 && suffixPathLen > 0 {
				continue
			}
			var matched bool
			// 路由匹配，参数捕获
			// route   /a/b/c => a | b | c
			// request /a/c/b => a | c | b -> not match
			// request /a/b/c => a | c | b -> matched
			// route   /a/:type/c => a | :type | c
			// request /a/  c  /b => a |  c    | b  -> not match
			// request /a/  c  /c => a |  c    | b  -> matched
			// route   /a/*path => a | *path
			// request /a/b/c/d/e => a | b/c/d/e  -> matched b/c/d/e
			for i, path := range routePathSplit {
				// 参数路径
				if isParamPath(path) {
					matched = true
					// * 路径匹配符，匹配后面所有路径
					if path[0] == '*' {
						p = append(p, gin.Param{
							Key:   path[1:],
							Value: strings.Join(suffixSplit[i:], "/"),
						})
						break
					}
					p = append(p, gin.Param{Key: path[1:], Value: suffixSplit[i]})
					continue
				}
				// 固定路径
				if path == suffixSplit[i] {
					matched = true
					continue
				}
				// 不匹配就跳过
				matched = false
				break
			}
			if matched {
				return &route, p
			}
		}
	}
	return nil, p
}

// 判断是否参数路径
func isParamPath(path string) bool {
	return path != "" && (path[0] == '*' || path[0] == ':')
}
