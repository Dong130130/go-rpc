package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "rpc_demo/docs" // swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 定义了服务的配置结构体，包含服务监听地址、端口号和用户认证信息
type Config struct {
	Host string            `json:"host"`
	Port int               `json:"port"`
	Auth map[string]string `json:"auth"`
}

// cfg 保存全局配置，在 main 中读取 config.json 初始化
var cfg Config

func main() {
	// 读取配置文件 config.json
	data, _ := os.ReadFile("config.json")
	json.Unmarshal(data, &cfg)
	// 初始化 Gin 引擎
	r := gin.Default()

	r.POST("/add", addHandler)
	r.POST("/division", divisionHandler)
	r.POST("/system", systemHandler)

	// 注册 Swagger UI，提供接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务，监听配置文件中的 host:port
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	r.Run(addr)
}
