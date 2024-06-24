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

// func ResponseXml(ctx *gin.Context, httpStatus int, data map[string]interface{}) {
// 	ctx.XML(httpStatus, data)
// }

func Success(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusOK, data, meta)
}

func SuccessXml(ctx *gin.Context, data interface{}) {
	ctx.XML(http.StatusOK, data)
	// ResponseXml(ctx, http.StatusOK, data)
}

func SuccessString(ctx *gin.Context, data string) {
	ctx.String(http.StatusOK, data)
	// ResponseXml(ctx, http.StatusOK, data)
}

func SuccessJson(ctx *gin.Context, data map[string]interface{}) {
	ctx.JSON(http.StatusOK, data)
	// ResponseXml(ctx, http.StatusOK, data)
}

func SuccessString201(ctx *gin.Context, data string) {
	ctx.String(http.StatusCreated, data)
	// ResponseXml(ctx, http.StatusOK, data)
}

func Success201(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusCreated, data, meta)
}

func Success400(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusBadRequest, data, meta)
}

func Fail(ctx *gin.Context, data map[string]interface{}, meta map[string]interface{}) {
	Response(ctx, http.StatusBadRequest, data, meta)
}
