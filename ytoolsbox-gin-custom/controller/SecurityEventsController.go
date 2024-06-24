package controller

import (
	// "strconv"
	// "log"
	// "strings"
	// "os/exec"

	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/service"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// 查询安全事件的条目
func GetAllSecurityEvents(ctx *gin.Context) {
	db := common.GetDB()

	GetAllSecurityEventsParam := dto.GetAllSecurityEventsDTOReq{}
	if utils.ResolveParam(ctx, &GetAllSecurityEventsParam) != nil {
		return
	}

	SecurityEventsList, DefaultLength := dao.SelectAllSecurityEventsDAO(db, GetAllSecurityEventsParam)

	// 构造返回的结构体
	SecurityEventsListData := dto.GetAllSecurityEventsDTOResp{Total: int64(DefaultLength), SecurityEventsList: SecurityEventsList}
	Meta := dto.SuccessResponseMeta{Message: "获取安全事件成功", StatusCode: 200}

	response.Success(ctx, utils.Struct2MapViaJson(SecurityEventsListData), utils.Struct2MapViaJson(Meta))
}

// 参数遍历获取大量敏感数据 或 频繁访问获取大量敏感数据 或 异常时间段频繁访问获取大量敏感数据
func RequestTraverseAndReturnTooMuchSensitiveData(ctx *gin.Context) {
	var data map[string]interface{} = make(map[string]interface{})
	data["1"] = "张一"
	data["2"] = "张二"
	data["3"] = "张三"
	data["4"] = "张四"
	data["5"] = "张五"
	Meta := dto.SuccessResponseMeta{Message: "模拟参数遍历获取大量敏感数据", StatusCode: 200}
	response.Success(ctx, data, utils.Struct2MapViaJson(Meta))
}

// 账号失陷并下载了大量敏感数据
func OneAccountGetSensitiveData(ctx *gin.Context) {
	err := service.ReplayMultiAccount("InnerSourceIP", "tcpreplayconfig")
	if err != nil {
		Meta := dto.SuccessResponseMeta{Message: "回放包失败", StatusCode: 401}
		response.Fail(ctx, nil, utils.Struct2MapViaJson(Meta))
		return
	}

	Meta := dto.SuccessResponseMeta{Message: "模拟疑似账号失陷并获取大量敏感数据", StatusCode: 200}
	response.Success(ctx, nil, utils.Struct2MapViaJson(Meta))
}
