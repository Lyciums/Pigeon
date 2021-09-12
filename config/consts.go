package config

const (
	// RouterModeFreedom 允许任意域名访问到框架路由
	RouterModeFreedom = "freedom"
	// RouterModeDomain 限制只有设定的域名能够访问到框架路由
	RouterModeDomain  = "domain"
	// LeftDelimit 模板变量左边分隔符
	LeftDelimit = "{*"
	// RightDelimit 模板变量右边边分隔符
	RightDelimit = "*}"
	// ResourceRootPath 资源路径根目录
	ResourceRootPath = "resources"
	// FileResourcePath 文件资源路径
	FileResourcePath = ResourceRootPath + "/files"
	// HTMLViewPath 网页视图路径
	HTMLViewPath = ResourceRootPath + "/views"
	// StaticResourcePath 静态资源路径
	StaticResourcePath = ResourceRootPath + "/statics"
)
