package resources

import (
	"Pigeon/config"
	"github.com/gin-gonic/gin"
)

type resource struct {
	route, path string
	listable    bool
}

var resources = []resource{
	// 静态资源文件
	{"/statics/", config.StaticResourcePath, false},
	{"/files/", config.FileResourcePath, true},
}

func Mount(e *gin.Engine) *gin.Engine {
	// 定义模板分隔符
	e.Delims(config.LeftDelimit, config.RightDelimit)
	// html 模板文件
	e.LoadHTMLGlob(config.HTMLViewPath + "/*")
	// 其他资源文件
	for _, resource := range resources {
		e.StaticFS(resource.route, gin.Dir(resource.path, resource.listable))
	}
	return e
}
