package controller

import (
	"fmt"
	"main/dto"
	"main/response"
	"main/service"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// 自定义请求
func CustomRequest(ctx *gin.Context) {

	// 先判断是否有需要返回的请求
	needResponseBytes := service.JudgeWhetherContainsNeedResponseData(ctx)

	// 尝试Json序列化
	jsonresult, err := service.TryJsonUnmarshal(needResponseBytes)
	if err == nil {
		Meta := dto.SuccessResponseMeta{Message: fmt.Sprintf("自定义%s请求", ctx.Request.Method), StatusCode: 200}
		response.Success(ctx, jsonresult, utils.Struct2MapViaJson(Meta))
		return
	}

	// 尝试xml序列化
	xmlresult, err1 := service.TryXMLUnmarshal(needResponseBytes)
	if err1 == nil && xmlresult != "" {
		ctx.XML(200, xmlresult)
		return
	}
	// 都不行,直接返回

	Meta := dto.SuccessResponseMeta{Message: fmt.Sprintf("自定义%s请求", ctx.Request.Method), StatusCode: 200}
	response.Success(ctx, gin.H{"Content": utils.BytesToString(needResponseBytes)}, utils.Struct2MapViaJson(Meta))
}
