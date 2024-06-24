package middleware

import (
	"log"
	"main/common"
	"main/dao"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveAllTriggerLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.String(), "/custom/mock/vulnerability") {
			db := common.GetDB()
			for _, item := range common.RiskAndVulnerabilityList {
				if item.Name == "重要接口未授权访问" || item.Name == "HTTP Basic Authorization弱口令" {
					if item.TriggerMethod == c.Request.URL.String() {
						log.Println("记录触发脆弱性操作: ", c.Request.URL, item.Name)
						dao.InsertTriggerLog(db, "脆弱性", item.Name, c.Request.URL.String())
						break
					}
				} else if strings.Contains(item.TriggerMethod, c.Request.URL.String()) {
					log.Println("记录触发脆弱性操作: ", c.Request.URL, item.Name)
					dao.InsertTriggerLog(db, "脆弱性", item.Name, c.Request.URL.String())
					break
				}
			}
		} else if strings.Contains(c.Request.URL.String(), "/custom/mock/risk") {
			db := common.GetDB()
			for _, item := range common.RiskAndVulnerabilityList {
				url := strings.Split(c.Request.URL.String(), "?")
				if strings.Contains(item.TriggerMethod, url[0]) {
					log.Println("记录触发风险操作: ", c.Request.URL, item.Name)
					dao.InsertTriggerLog(db, "风险", item.Name, c.Request.URL.String())
					break
				}
			}
		} else {
			db := common.GetDB()
			for _, item := range common.RiskAndVulnerabilityList {
				if item.TriggerMethod == c.Request.URL.String() {
					log.Println("记录触发脆弱性操作: ", c.Request.URL, item.Name)
					dao.InsertTriggerLog(db, "脆弱性", item.Name, c.Request.URL.String())
				}
			}
		}
		// 处理请求
		c.Next()
	}
}
