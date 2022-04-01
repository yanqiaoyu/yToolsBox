package dto

// 前端要求的cascader信息格式如下
// [
//         {
//           value: '东南',
//           label: '东南',
//           children: [
//             { value: '普陀', label: '普陀' },
//             { value: '黄埔', label: '黄埔' },
//             { value: '徐汇', label: '徐汇' },
//           ],
//         },
// ]

type CascaderInfo struct {
	Total        int64
	CascaderList []CascaderFatherNode
}

// 父节点是工具
type CascaderFatherNode struct {
	// 工具的ID
	// 工具的名称
	Value    int               `json:"value" gorm:"column:id"`
	Label    string            `json:"label" gorm:"column:toolName"`
	Children []CascaderSonNode `json:"children"`
}

// 子节点是工具的配置
type CascaderSonNode struct {
	// 配置的ID
	// 配置的名称
	Value int    `json:"value" gorm:"column:id"`
	Label string `json:"label" gorm:"column:toolConfigName"`
}

type PostNewTaskDTOReq struct {
	ConfigList string `form:"ConfigList" json:"ConfigList" binding:"required"`
}

type PostRestartTaskDTOReq struct {
	ToolName       string `form:"toolName" json:"toolName" binding:"required"`
	ToolConfigName string `form:"toolConfigName" json:"toolConfigName" binding:"required"`
}

type GetAllTaskItemDTOReq struct {
	Query    string `json:"query" form:"query" `
	Pagenum  int    `json:"pagenum" form:"pagenum" binding:"required"`
	Pagesize int    `json:"pagesize" form:"pagesize" binding:"required"`
}

type GetAllTaskItemDTOResp struct {
	Total        int64
	TaskItemList []map[string]interface{}
}
