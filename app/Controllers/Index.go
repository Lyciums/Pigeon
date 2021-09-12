package Controllers

import (
	"github.com/gin-gonic/gin"
)

type Index struct{}

func (i *Index) RegisterRouter(r *gin.RouterGroup) {
	r.GET("", i.Index)
	r.GET("a/", getRoutePath)
	r.GET("b/", getRoutePath)
	r.GET("c/", getRoutePath)
	r.GET("c/d/", getRoutePath)
	r.GET("c/d/e/", getRoutePath)
}

func (i *Index) Index(c *gin.Context) {
	c.String(200, "index.")
}

func getRoutePath(c *gin.Context) {
	domainRoutePath, ok := c.Get("domain_route_path")
	msg := "域名路由 - 自由模式\n访问路径：" + c.FullPath()
	if ok {
		msg = "域名路由 - 域名限定模式\n请求的域名: " + c.Request.Host + "\n解析分配到的路由：" + domainRoutePath.(string)
	}
	c.String(200, msg)
}
