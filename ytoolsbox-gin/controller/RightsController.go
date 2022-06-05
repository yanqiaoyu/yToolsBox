package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func GetRights(ctx *gin.Context) {
	db := common.GetDB()

	rightsList := dao.SelectAllRights(db)

	// 构造返回的结构体
	Meta := dto.SuccessResponseMeta{Message: "获取权限成功", StatusCode: 200}
	Data := dto.GetAllRightsResp{Total: int64(len(rightsList)), RightsList: rightsList}

	response.Success(ctx, utils.Struct2MapViaJson(Data), utils.Struct2MapViaJson(Meta))
}
