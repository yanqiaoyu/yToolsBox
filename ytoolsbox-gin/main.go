package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Cors()) //开启中间件 允许使用跨域请求
	r.POST("/login", func(c *gin.Context) {

		name := c.PostForm("name")
		password := c.PostForm("password")
		log.Print(name, password)

		// 账户名密码错误
		if password != "admin" || name != "admin" {
			data := map[string]interface{}{
				"data": map[string]interface{}{
					"status_code": 401,
					"message":     "账号密码错误",
				},
			}
			c.JSON(200, data)
			return
		}

		data := map[string]interface{}{
			"data": map[string]interface{}{
				"status_code": 200,
				"message":     "登录成功",
				"token":       "123456",
			},
		}

		c.JSON(200, data)
	})
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
