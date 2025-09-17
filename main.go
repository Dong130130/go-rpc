package main

import (
	"fmt"
	"os"
	"encoding/json"

	_ "rpc_demo/docs" // swagger docs
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


type Config struct {
	Host string            `json:"host"`
	Port int               `json:"port"`
	Auth map[string]string `json:"auth"`
}

var cfg Config

func main() {
	// 读取配置
	data, _ := os.ReadFile("config.json")
	json.Unmarshal(data, &cfg)

	r := gin.Default()

	r.POST("/add", addHandler)
	r.POST("/division", divisionHandler)
	r.POST("/system", systemHandler)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	r.Run(addr)
}
