package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Args struct {
	A        float64 `json:"a"`
	B        float64 `json:"b"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

type SystemCmd struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkAuth(user, pass string) bool {
	if p, ok := cfg.Auth[user]; ok && p == pass {
		return true
	}
	return false
}

// @Summary Add two numbers
// @Description 计算两个数的和
// @Accept json
// @Produce json
// @Param data body Args true "请求体"
// @Success 200 {object} map[string]float64
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Router /add [post]
func addHandler(c *gin.Context) {
	var a Args
	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数有问题"})
		return
	}
	if !checkAuth(a.Username, a.Password) {
		c.JSON(401, gin.H{"error": "认证失败"})
		return
	}
	c.JSON(200, gin.H{"result": a.A + a.B})
}

// @Summary Divide two numbers
// @Description 计算两个数的商
// @Accept json
// @Produce json
// @Param data body Args true "请求体"
// @Success 200 {object} map[string]float64
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Router /division [post]
func divisionHandler(c *gin.Context) {
	var a Args
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(400, gin.H{"error": "参数解析失败"})
		return
	}
	if !checkAuth(a.Username, a.Password) {
		c.JSON(401, gin.H{"error": "认证失败"})
		return
	}
	if a.B == 0 {
		c.JSON(400, gin.H{"error": "除数不能是0"})
		return
	}
	c.JSON(200, gin.H{"result": a.A / a.B})
}

// @Summary Run system command
// @Description 执行系统命令并返回结果
// @Accept json
// @Produce json
// @Param data body SystemCmd true "请求体"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Router /system [post]
func systemHandler(c *gin.Context) {
	var cmd SystemCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(400, gin.H{"error": "参数解析失败"})
		return
	}
	if !checkAuth(cmd.Username, cmd.Password) {
		c.JSON(401, gin.H{"error": "认证失败"})
		return
	}

	var ec *exec.Cmd
	if runtime.GOOS == "windows" {
		ec = exec.Command("cmd", "/C", cmd.Cmd)
	} else {
		ec = exec.Command("bash", "-c", cmd.Cmd)
	}

	out, err := ec.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{
			"error":  fmt.Sprintf("命令失败: %v", err),
			"output": string(out),
		})
		return
	}

	c.JSON(200, gin.H{"output": string(out)})
}
