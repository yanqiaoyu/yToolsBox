package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, data map[string]interface{}, meta map[string]interface{}) {
	ctx.JSON(httpStatus, map[string]interface{}{
		"data": data,
		"meta": meta,
	})
}

func Success(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusOK, data, meta)
}

func Fail(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusOK, data, meta)
}
