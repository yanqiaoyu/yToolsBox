package model

type Menus struct {
	// 存放数据的
	Data []MenusData `json:"data"`
	// 存放返回值信息的
	Meta MenusMeta `json:"meta"`
}

type MenusData struct {
	// 模块的ID
	Id int `json:"id"`
	// 模块的名称
	AuthName string `json:"authName"`
	// 模块的路径
	Path string `json:"path"`
	// 模块的子模块
	ChildMenus []ChildMenus `json:"child"`
}

type MenusMeta struct {
	Msg         string `json:"msg"`
	Status_code int    `json:"status_code"`
}

type ChildMenus struct {
	// 子模块的ID
	Id int `json:"id"`
	// 子模块的名称
	AuthName string `json:"authName"`
	// 子模块的路径
	Path string `json:"path"`
}
