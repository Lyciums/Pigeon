package Interfaces

import (
	"strings"

	"Pigeon/app/Caches/App"
	"github.com/gin-gonic/gin"
)

type (
	RouterController interface {
		RegisterRouter(*gin.RouterGroup)
	}
	Domain struct {
		Name  string
		Subs  DomainRouterMap
		Match bool
	}
	DomainRouterMap      map[string]*Domain
	RouterControllerList []RouterController
)

func SetupRouter(routerGroup *gin.RouterGroup, controllers RouterControllerList) *gin.RouterGroup {
	for _, controller := range controllers {
		controller.RegisterRouter(routerGroup)
	}
	return routerGroup
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
