package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/model"
	"main/response"
	"main/service"
	"main/util"

	"github.com/gin-gonic/gin"
)

// 新增一个工具
func PostNewTool(ctx *gin.Context) {
	// 提取工具的基本信息
	PostNewToolBasicInfoDTOReq := dto.PostNewToolBasicInfoDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolBasicInfoDTOReq) != nil {
		return
	}

	// 提取工具的配置信息
	PostNewToolConfigInfoDTOReq := dto.PostNewToolConfigInfoDTOReq{}
	// 提取参数
	if util.ResolveParam(ctx, &PostNewToolConfigInfoDTOReq) != nil {
		return
	}

	// log.Print(util.Struct2MapViaJson(PostNewToolBasicInfoDTOReq), util.Struct2MapViaJson(PostNewToolConfigInfoDTOReq))

	db := common.GetDB()

	// 写入基本信息
	result := dao.InsertNewToolBasicInfo(db, &PostNewToolBasicInfoDTOReq)
	if result.Error != nil {
		// 再调用一次Error(),转换为字符串
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"tools_toolName_key\" (SQLSTATE 23505)" {
			msg := dto.FailResponseMeta{StatusCode: 400, Message: "不允许重复的工具名称"}
			response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
			return
		}
	}

	// 写入配置信息
	PostNewToolConfigInfoDTOReq.ToolID = int(PostNewToolBasicInfoDTOReq.ID)
	result = dao.InsertNewToolConfigInfo(db, &PostNewToolConfigInfoDTOReq)
	if result.Error != nil {
		// 再调用一次Error(),转换为字符串
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "工具配置信息写入错误"}
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
		return
	}

	Meta := model.Meta{Msg: "新增工具成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 查询所有工具
func GetAllTools(ctx *gin.Context) {
	db := common.GetDB()

	GetAllToolsDTOReq := dto.GetAllToolsDTOReq{}

	if util.ResolveParam(ctx, &GetAllToolsDTOReq) != nil {
		return
	}

	tools_List := dao.SelectAllTools(db, GetAllToolsDTOReq)

	// 构造返回的结构体
	ToolData := dto.GetAllToolsDTOResp{Total: len(tools_List), Tools: tools_List}
	Meta := dto.SuccessResponseMeta{Message: "获取工具成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

// 查询某个工具的所有配置
func GetSpecifiedToolConfig(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedToolConfigDTOReqURI := dto.GetSpecifiedToolConfigDTOReqURI{}
	GetSpecifiedToolConfigDTOReqQuery := dto.GetSpecifiedToolConfigDTOReqQuery{}

	if util.ResolveURI(ctx, &GetSpecifiedToolConfigDTOReqURI) != nil {
		return
	}
	if util.ResolveQuery(ctx, &GetSpecifiedToolConfigDTOReqQuery) != nil {
		return
	}

	configList, DefaultLength := dao.SelectSpecifiedToolConfig(db, int(GetSpecifiedToolConfigDTOReqURI.ToolID), GetSpecifiedToolConfigDTOReqQuery)
	// log.Println(configList)

	// 构造返回的结构体
	ToolData := dto.GetSpecifiedToolConfigDTOResp{Total: DefaultLength, ToolConfig: configList, ID: GetSpecifiedToolConfigDTOReqURI.ToolID}
	Meta := model.Meta{Msg: "查询配置成功", Status_code: 200}
	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

// 上传脚本文件
func PostScriptFile(ctx *gin.Context) {
	db := common.GetDB()
	toolName, FileDST, err := service.SaveScriptFile(ctx)
	if err != nil {
		msg := dto.FailResponseMeta{StatusCode: 400, Message: "上传文件失败"}
		response.Fail(ctx, nil, util.Struct2MapViaJson(msg))
	}

	dao.UpdateToolConfigScriptLocalPath(db, toolName, FileDST)

	Meta := model.Meta{Msg: "上传文件成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 新增配置
func PostNewConfig(ctx *gin.Context) {
	db := common.GetDB()
	PostNewConfig := dto.PostNewToolConfigInfoDTOReq{}
	if util.ResolveParam(ctx, &PostNewConfig) != nil {
		return
	}
	// log.Println(util.Struct2MapViaJson(PostNewConfig))

	dao.InsertExistToolConfigInfo(db, &PostNewConfig)

	Meta := model.Meta{Msg: "新增配置成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 删除特定的配置文件
func DeleteSpecifiedConfig(ctx *gin.Context) {
	db := common.GetDB()
	DeleteSpecifiedConfigDTOReq := dto.DeleteSpecifiedConfigDTOReq{}

	if util.ResolveURI(ctx, &DeleteSpecifiedConfigDTOReq) != nil {
		return
	}

	dao.DeleteSpecifiedConfig(db, DeleteSpecifiedConfigDTOReq.ID)

	Meta := model.Meta{Msg: "删除配置成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 查询某个工具的某个配置
func GetSpecifiedToolConfigByConfigID(ctx *gin.Context) {
	db := common.GetDB()
	GetSpecifiedToolConfigByConfigID := dto.GetSpecifiedToolConfigByConfigIDDTOReq{}

	if util.ResolveURI(ctx, &GetSpecifiedToolConfigByConfigID) != nil {
		return
	}

	toolConfig := dao.SelectSpecifiedToolConfigByConfigID(db, GetSpecifiedToolConfigByConfigID.ID)

	// 构造返回的结构体
	ToolData := dto.GetSpecifiedToolConfigByConfigIDDTOResp{ToolConfig: toolConfig}
	Meta := model.Meta{Msg: "查询配置成功", Status_code: 200}
	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

// 更新某个工具的某个配置
func PutSpecifiedToolConfigByConfigID(ctx *gin.Context) {
	db := common.GetDB()
	PutSpecifiedToolConfigByConfigIDDTOReqURI := dto.PutSpecifiedToolConfigByConfigIDDTOReqURI{}
	PutSpecifiedToolConfigByConfigIDDTOReqQuery := dto.PutSpecifiedToolConfigByConfigIDDTOReqQuery{}

	if util.ResolveURI(ctx, &PutSpecifiedToolConfigByConfigIDDTOReqURI) != nil {
		return
	}
	if util.ResolveParam(ctx, &PutSpecifiedToolConfigByConfigIDDTOReqQuery) != nil {
		return
	}

	dao.UpdateSpecifiedToolConfigByConfigID(db, PutSpecifiedToolConfigByConfigIDDTOReqURI.ID, PutSpecifiedToolConfigByConfigIDDTOReqQuery)

	Meta := model.Meta{Msg: "更新配置成功", Status_code: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}
