package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	db := common.GetDB()
	loginParam := dto.LoginDTO{}
	msg := dto.FailResponseMeta{}

	// 模型绑定获取参数
	utils.ResolveParam(ctx, &loginParam)

	if !dao.IsUserExist(db, loginParam) {
		msg.StatusCode = 400
		msg.Message = "账户不存在"
		response.Fail(ctx, nil, utils.Struct2MapViaJson(msg))
		return
	}

	if !dao.CheckPassword(db, loginParam) {
		msg.StatusCode = 400
		msg.Message = "账户或密码错误"
		response.Fail(ctx, nil, utils.Struct2MapViaJson(msg))
		return
	}

	msg.StatusCode = 200
	msg.Message = "登录成功"
	response.Success(ctx, gin.H{"token": "123456"}, utils.Struct2MapViaJson(msg))
}
