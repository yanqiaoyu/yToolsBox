package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// 窃取数据接口
func DataLeakage(ctx *gin.Context) {
	dataLeakgeDB := common.GetDataLeakgeDB()

	Meta := dto.SuccessResponseMeta{Message: "窃取信息成功", StatusCode: 200}
	response.Success(ctx, gin.H{"Content": dao.SelectRandomLeakgeData(dataLeakgeDB)}, utils.Struct2MapViaJson(Meta))
}
