package middleware

import (
	"main/common"
	"main/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin 的中间件就是这种格式
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// 如果这个token为空，或者不以 Bearer开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "权限不足",
				"code": 401,
			})
			// 那就将这次请求抛弃，并返回
			ctx.Abort()
			return
		}

		// run到这里，说明token有效 ，提取有效部分
		tokenString = tokenString[7:]

		// 定义一个函数，解析token
		token, claims, err := common.ParseToken(tokenString)

		// 如果解析出错,或者返回的token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "权限不足",
				"code": 401,
			})
			// 那就将这次请求抛弃，并返回
			ctx.Abort()
			return
		}

		// 否则，token通过验证，拿取userID
		userID := claims.UserID
		DB := common.GetDB()
		var user model.User

		// 去数据库里找user
		DB.First(&user, userID)

		// 没找到
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "权限不足",
				"code": 401,
			})
			// 那就将这次请求抛弃，并返回
			ctx.Abort()
			return
		}

		// 找到了，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}

}
