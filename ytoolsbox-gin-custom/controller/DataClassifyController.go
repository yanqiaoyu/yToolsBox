package controller

import (
	"log"
	"main/common"
	"main/dao"
	"main/dto"
	grpcclient "main/grpc/grpc_client"
	"main/response"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// 窃取数据接口
func DataClassify(ctx *gin.Context) {
	db := common.GetDB()
	// 从数据库中,拿到大脑的IP,前端账号,前端密码
	result, err := dao.SelectModifyDSCThresholdConfig(db)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "大脑配置查询失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}
	log.Print("需要新增数据源的大脑IP: ", result.DSCAddress)
	log.Print("需要新增数据源的大脑前端账号: ", result.DSCWebUserName)
	log.Print("需要新增数据源的大脑前端密码: ", result.DSCWebPassword)
	log.Print("增数据源的IP: ", result.ToolBoxAddress)

	err = grpcclient.AddDataClassify(result)
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "新增数据源失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "新增数据源成功", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}
