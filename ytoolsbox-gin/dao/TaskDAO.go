package dao

import (
	"main/dto"
	"main/model"
	"main/service"
	"main/util"

	"gorm.io/gorm"
)

// 查询级联选择器的配置
func SelectCascaderInfo(db *gorm.DB) []dto.CascaderFatherNode {
	// 先找父节点
	FatherNode := []dto.CascaderFatherNode{}
	db.Model(&model.Tool{}).Find(&FatherNode)
	// log.Println(FatherNode)

	for i := 0; i < len(FatherNode); i++ {
		// 遍历每一个父节点，用父节点的ID查询所有子节点
		SonNode := []dto.CascaderSonNode{}
		db.Model(&model.ToolConfig{}).Where("tool_configs.\"toolID\" = ?", FatherNode[i].Value).Find(&SonNode)
		// log.Println(SonNode)
		FatherNode[i].Children = SonNode
	}

	return FatherNode
}

// 根据配置的ID查出一条配置
func SelectConfigByToolID(db *gorm.DB, toolID string) dto.BriefToolConfigDTO {
	configList := dto.BriefToolConfigDTO{}
	db.Model(&model.ToolConfig{}).Where("id = ?", toolID).Find(&configList)
	return configList
}

// 新增一个任务条目
func InsertTaskItem(db *gorm.DB, config dto.BriefToolConfigDTO) {
	tool := model.Tool{}
	db.Model(&model.Tool{}).Select("\"toolName\"").Where("id = ?", config.ToolID).First(&tool)

	TaskItem := model.Tasks{
		ToolConfigName: config.ToolConfigName,
		ToolName:       tool.ToolName,
		Progress:       0,
		IsDone:         false,
		ReturnContent:  "任务开始:\r\n",
	}
	db.Model(&model.Tasks{}).Create(&TaskItem)
}

// 查询所有任务的条目
func SelectAllTaskItem(db *gorm.DB, obj dto.GetAllTaskItemDTOReq) ([]map[string]interface{}, int) {
	taskItem := []model.Tasks{}
	query := obj.Query
	pagenum := obj.Pagenum
	pagesize := obj.Pagesize
	map_taskItemList := []map[string]interface{}{}

	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		// 按照时间升序
		db.Order("created_at desc").Model(&model.Tasks{}).Find(&taskItem)
	} else {
		// 按照时间升序
		db.Order("created_at desc").Where("\"toolName\" LIKE ?", "%"+query+"%").Model(&model.Tasks{}).Find(&taskItem)
	}

	DefaultLength := len(taskItem)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := util.Struct2MapViaJson(taskItem[i])
		map_taskItemList = append(map_taskItemList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := service.CalculateReturnMapLength(pagenum, pagesize, map_taskItemList)
	// 返回切片后的结果
	return map_taskItemList[ArrayStart:ArrayEnd], DefaultLength
}
