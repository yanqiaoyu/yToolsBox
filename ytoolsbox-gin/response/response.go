package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data map[string]interface{}, meta map[string]interface{}) {
	ctx.JSON(httpStatus, map[string]interface{}{
		// "code": code,
		// "data": data,
		// "msg":  msg,
		"data": data,
		"meta": meta,
	})
}

func Success(ctx *gin.Context, msg map[string]interface{}, data map[string]interface{}) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// func Fail(ctx *gin.Context, msg string, data gin.H) {
// 	Response(ctx, http.StatusOK, 400, data, msg)
// }
