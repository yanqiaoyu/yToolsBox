package dao

import (
	"log"
	"main/model"
	"main/utils"

	"gorm.io/gorm"
)

// 查询配置
func SelectPOCConfig(db *gorm.DB) (model.POCConfig, error) {
	POCConfig := model.POCConfig{}
	result := db.Debug().Model(&model.POCConfig{}).First(&POCConfig)
	if result.Error != nil {
		return POCConfig, result.Error
	}
	return POCConfig, nil
}

// 保存配置
func InsertSavePOCConfig(db *gorm.DB, SavePOCConfigParam model.POCConfig) error {
	// 先查一下表中有没有数据,没有就create, 有就update
	flag := model.POCConfig{}

	result := db.Debug().Model(&model.POCConfig{}).First(&flag)

	if result.RowsAffected == 0 {
		log.Println("POC配置表中无数据,新建条目")
		db.Debug().Create(&SavePOCConfigParam)
	} else {
		log.Println("POC配置表中有数据,更新条目")
		db.Debug().Model(&model.POCConfig{}).Where("id = ?", flag.ID).Updates(SavePOCConfigParam)
	}

	return nil
}

// 查询Agent安装配置
func SelectAgentInstallConfig(db *gorm.DB) ([]map[string]interface{}, int, error) {
	AgentInstallConfig := []model.AgentInstallConfig{}
	map_AgentInstallConfigList := []map[string]interface{}{}

	result := db.Debug().Model(&model.AgentInstallConfig{}).Find(&AgentInstallConfig)

	DefaultLength := len(AgentInstallConfig)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < DefaultLength; i++ {
		map_item := utils.Struct2MapViaJson(AgentInstallConfig[i])
		map_AgentInstallConfigList = append(map_AgentInstallConfigList, map_item)
	}

	log.Print("查询的Agent安装的配置信息: ", AgentInstallConfig, result)

	return map_AgentInstallConfigList, DefaultLength, nil
}

// 保存Agent安装配置
func InsertAgentInstallConfig(db *gorm.DB, SaveAgentInstallConfigParam model.AgentInstallConfig) error {
	// 先查一下表中有没有数据,没有就create, 有就update
	flag := model.AgentInstallConfig{}

	result := db.Debug().Model(&model.AgentInstallConfig{}).Where("type = ?", SaveAgentInstallConfigParam.Type).First(&flag)

	if result.RowsAffected == 0 {
		log.Println("Agent配置表中无数据,新建条目")
		db.Debug().Create(&SaveAgentInstallConfigParam)
	} else {
		log.Println("Agent配置表中有数据,更新条目")
		db.Debug().Model(&model.AgentInstallConfig{}).Where("id = ?", flag.ID).Updates(SaveAgentInstallConfigParam)
	}

	return nil
}

// 从数据库中,拿到大脑的IP,前端账号,前端密码
func SelectModifyDSCThresholdConfig(db *gorm.DB) (model.POCConfig, error) {
	ModifyDSCThresholdConfig := model.POCConfig{}
	result := db.Debug().Model(&model.POCConfig{}).First(&ModifyDSCThresholdConfig)
	if result.Error != nil {
		return ModifyDSCThresholdConfig, result.Error
	}
	return ModifyDSCThresholdConfig, nil
}
