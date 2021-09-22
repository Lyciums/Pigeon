package config

import (
	"Pigeon/app/Interfaces"
	"Pigeon/library/env"
)

var (
	// RouterMode : 域名路由模式
	//	Domain: 限制域名访问
	//	freedom (默认): 只要绑定了服务端口就都能访问
	RouterMode = env.GetOrDefault("ROUTER_MODE", RouterModeFreedom)
	// Domain : 可以在 .env 里设置一个域名
	Domain = env.GetOrDefault("ROUTER_Interfaces.Domain", "test.com")
	// DomainConfig : 域名绑定，支持多个域名，支持无限嵌套
	// 一级必须为主域名
	// 子域名对应路由的关系：
	// 	 test.com -> /
	//	 a.test.com -> /a/ ==  test.com/a/
	//	 b.test.com -> /b/ ==  test.com/b/
	// 	 c.test.com -> /c/ ==  test.com/c/
	// 	 d.c.test.com -> /c/d/ == test.com/c/d/ == c.test.com/d/
	// 	 e.d.c.test.com -> /c/d/e/ == test.com/c/d/e/ == c.test.com/d/e/ == c.test.com/d/
	DomainConfig = Interfaces.DomainRouterMap{
		// demo: test.com
		Domain: &Interfaces.Domain{
			Name: Domain,
			Subs: Interfaces.DomainRouterMap{
				// a.test.com
				"a": nil,
				// b.test.com
				"b": nil,
				// c.test.com
				"c": &Interfaces.Domain{
					// c1.c.test.com
					Subs: Interfaces.DomainRouterMap{
						"d": &Interfaces.Domain{
							// e.d.c.test.com
							Subs: Interfaces.DomainRouterMap{
								"e": nil,
							},
						},
					},
				},
			},
		},
		// demo: dev.com
		"dev.com": &Interfaces.Domain{
			Name: "dev.com",
			Subs: Interfaces.DomainRouterMap{
				// proxy.dev.com
				"proxy": &Interfaces.Domain{
					// *.proxy.dev.com
					Match: true,
				},
				// download.dev.com
				"download": nil,
				// video.dev.com
				"video": nil,
			},
		},
	}
)

func init() {
	println("route mode:", RouterMode)
	println("route Interfaces.Domain:", Domain)
}
