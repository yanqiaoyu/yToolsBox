/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu?tab=repositories
 * @Date: 2021-09-12 12:18:07
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 12:18:13
 * @FilePath: /ytoolsbox-gin/middleware/CORS.go
 */
package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		if origin != "" {

			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Cache-Control")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 一个局部中间件,用来允许cookie交互
func AllowCookieMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("使用的是允许携带Cookie的中间件")
		cookie := c.Request.Header.Get("Cookie")
		log.Println("携带的cookie如下: ", cookie)

		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		if origin != "" {

			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Credentials")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Cookie, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Cache-Control")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
