package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/service"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// 清除脆弱性与风险的日志
func DeleteRiskAndVunl(ctx *gin.Context) {
	// 1.拿到大脑的SSH相关信息
	db := common.GetDB()
	//查询配置
	POCConfig, err := dao.SelectPOCConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "查询配置失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
	}

	_, err = service.DeleteRiskAndVunlService(POCConfig)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "删除脆弱性与风险日志失败: " + err.Error(), StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	// 2.执行脚本

	Meta := dto.SuccessResponseMeta{Message: "删除脆弱性与风险日志成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}
