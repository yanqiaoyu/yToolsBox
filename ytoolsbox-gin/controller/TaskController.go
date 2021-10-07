package controller

import (
	"main/common"
	"main/dao"
	"main/dto"
	"main/response"
	"main/service"
	"main/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// 新增一个任务
func PostNewTask(ctx *gin.Context) {
	db := common.GetDB()
	PostNewTaskParam := dto.PostNewTaskDTOReq{}

	if util.ResolveParam(ctx, &PostNewTaskParam) != nil {
		return
	}
	// 这里获取到的PostNewTaskParam是一个字符串形式的数组，所以还需要处理
	PostNewTaskParam.ConfigList = strings.TrimPrefix(PostNewTaskParam.ConfigList, "[")
	PostNewTaskParam.ConfigList = strings.TrimSuffix(PostNewTaskParam.ConfigList, "]")
	// log.Println(PostNewTaskParam.ConfigList)

	configIDList := strings.Split(PostNewTaskParam.ConfigList, ",")

	// log.Println(configIDList)
	// 获取每一个配置ID，新增一个任务，新增一个任务进度条目
	for i := 0; i < len(configIDList); i++ {
		// 从库里面把这个配置ID的全部信息查出来传给service
		configList := dao.SelectConfigByToolID(db, string(configIDList[i]))
		// 新增一个条目
		dao.InsertTaskItem(db, configList)
		go service.CreateNewTaskService(configList)
	}

	// 返回
	Meta := dto.SuccessResponseMeta{Message: "新建任务成功", StatusCode: 200}
	response.Success(ctx, nil, util.Struct2MapViaJson(Meta))
}

// 获取级联选择器中的信息
func GetCascader(ctx *gin.Context) {
	db := common.GetDB()
	CascaderTree := dao.SelectCascaderInfo(db)

	// 构造返回的结构体
	ToolData := dto.CascaderInfo{Total: int64(len(CascaderTree)), CascaderList: CascaderTree}
	Meta := dto.SuccessResponseMeta{Message: "获取配置信息成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(ToolData), util.Struct2MapViaJson(Meta))
}

// 获取所有TasksItem(任务进度)
func GetTaskItem(ctx *gin.Context) {
	db := common.GetDB()

	GetAllTaskItemParam := dto.GetAllTaskItemDTOReq{}
	if util.ResolveParam(ctx, &GetAllTaskItemParam) != nil {
		return
	}

	// log.Println(GetAllTaskItemParam)

	TaskItemList, DefaultLength := dao.SelectAllTaskItem(db, GetAllTaskItemParam)

	// log.Println(TaskItemList)

	// 构造返回的结构体
	TaskItemData := dto.GetAllTaskItemDTOResp{Total: int64(DefaultLength), TaskItemList: TaskItemList}
	Meta := dto.SuccessResponseMeta{Message: "获取任务列表成功", StatusCode: 200}

	response.Success(ctx, util.Struct2MapViaJson(TaskItemData), util.Struct2MapViaJson(Meta))
}
