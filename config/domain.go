package config

import (
	"Pigeon/library/env"
)

type Domain struct {
	Name string
	Subs DomainMap
}

type DomainMap map[string]*Domain

var (
	// RouterMode : 域名路由模式
	//	domain: 限制域名访问
	//	freedom (默认): 只要绑定了服务端口就都能访问
	RouterMode = env.GetOrDefault("ROUTER_MODE", RouterModeFreedom)
	// RootDomain : 可以在 .env 里设置一个域名
	RootDomain = env.GetOrDefault("ROUTER_DOMAIN", "test.com")
	// DomainConfig : 域名绑定，支持多个域名，支持无限嵌套
	// 一级必须为主域名
	// 子域名对应路由的关系：
	// 	 test.com -> /
	//	 a.test.com -> /a/ ==  test.com/a/
	//	 b.test.com -> /b/ ==  test.com/b/
	// 	 c.test.com -> /c/ ==  test.com/c/
	// 	 d.c.test.com -> /c/d/ == test.com/c/d/ == c.test.com/d/
	// 	 e.d.c.test.com -> /c/d/e/ == test.com/c/d/e/ == c.test.com/d/e/ == c.test.com/d/
	DomainConfig = DomainMap{
		// demo: test.com
		RootDomain: &Domain{
			Name: RootDomain,
			Subs: DomainMap{
				// a.test.com
				"a": nil,
				// b.test.com
				"b": nil,
				// c.test.com
				"c": &Domain{
					// c1.c.test.com
					Subs: DomainMap{
						"d": &Domain{
							// e.d.c.test.com
							Subs: DomainMap{
								"e": nil,
							},
						},
					},
				},
			},
		},
		// demo: dev.com
		"dev.com": &Domain{
			Name: "dev.com",
			Subs: DomainMap{
				// proxy.dev.com
				"proxy": &Domain{
					// pac.proxy.dev.com
					Subs: DomainMap{
						"pac": nil,
					},
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
	println("route domain:", RootDomain)
}
