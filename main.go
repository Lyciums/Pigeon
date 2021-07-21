package main

import (
	"Pigeon/bootstrap"
	"Pigeon/library/env"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	// load .env file
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	engine := gin.Default()
	// 设置业务路由，中间件，资源目录等
	bootstrap.Setup(engine)
	// 启动服务
	_ = endless.ListenAndServe(":"+env.GetOrDefault("SERVICE_PORT", "8888"), engine)
}
