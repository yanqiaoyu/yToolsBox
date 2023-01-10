/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 14:50:51
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-06-22 18:28:05
 * @FilePath: \golang_web\router.go
 */

package main

import (
	"main/controller"
	"main/middleware"
	"main/service"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	URL_Prefix := "/api/auth"
	/* 用路由组重新归纳了一下路由 */
	v1 := r.Group(URL_Prefix)
	{
		// r.POST(URL_Prefix + "/signup", controller.SignUp)
		// 登录
		v1.POST("/login", controller.Login)
		// r.GET(URL_Prefix + "/info", middleware.AuthMiddleWare(), controller.Info)
		// 获取菜单信息
		v1.GET("/menus", service.GetMenus)
		// 获取所有用户的信息
		v1.GET("/users", controller.GetAllUser)
		// 获取特定用户的信息
		v1.GET("/users/:userID", controller.GetSpecifiedUser)
		// 更新特定用户的状态
		v1.PUT("/users/state", controller.PutUserState)
		// 更新特定用户的信息
		v1.PUT("/users/:userID", controller.PutUserInfo)
		// 新增用户
		v1.POST("/users", controller.PostNewUser)
		// 删除用户
		v1.DELETE("/users/:userID", controller.DeleteSpecifiedUser)

		// 获取权限
		v1.GET("/rights", controller.GetRights)

		/***
			以下是工具相关的路由表
		***/
		// 更新某个工具的使用说明
		v1.PUT("/tools/tutorial", controller.PutSpecifiedToolTutorialByToolID)

		// 添加新工具
		v1.POST("/tools", controller.PostNewTool)
		// 查询所有工具
		v1.GET("/tools", controller.GetAllTools)
		// 删除所有工具
		v1.DELETE("/tools", controller.DeleteAllTools)
		// 查询某个工具的所有配置
		v1.GET("/tools/config/:toolID", controller.GetSpecifiedToolConfig)
		// 查询某个工具的某个配置
		v1.GET("/tools/config/:toolID/:configID", controller.GetSpecifiedToolConfigByConfigID)
		// 更新某个工具的某个配置
		v1.PUT("/tools/config/:toolID/:configID", controller.PutSpecifiedToolConfigByConfigID)
		// 为某个工具新增配置
		v1.POST("/tools/config/:toolID", controller.PostNewConfig)
		// 删除某个工具下的某个配置
		v1.DELETE("/tools/config/:toolID/:configID", controller.DeleteSpecifiedConfig)
		// 上传脚本文件
		v1.POST("/upload", controller.PostScriptFile)

		/***
			以下是任务相关的路由表
		***/

		// 新建一个任务
		v1.POST("/tasks", controller.PostNewTask)
		// 查询Cascader里面的信息
		v1.GET("/tasks/cascader", controller.GetCascader)
		// 查询所有的TaskItem(任务进度)
		v1.GET("/tasks", controller.GetTaskItem)
		// 清空所有任务
		v1.DELETE("/tasks", controller.DeleteAllTask)
		// 删除特定任务
		v1.DELETE("/tasks/:taskID", controller.DeleteSpecifiedTask)
		// 重新开始执行一个任务
		v1.POST("/tasks/restart", controller.PostRestartTask)

		/***
			以下是定时任务相关的路由表
		***/

		// 新建一个定时任务
		v1.POST("/crontasks", controller.PostNewCronTask)
		// 清除所有定时任务
		v1.DELETE("/crontasks", controller.DeleteAllCronTask)
		// 清除特定定时任务
		v1.DELETE("/crontasks/:cronTaskOriginID/:cronTaskScheduleID", controller.DeleteSpecifiedCrontask)

		// 查询所有特定定时任务
		v1.GET("/crontasks", controller.GetAllCronTask)
		// 根据scheduleID查询特定定时任务
		v1.GET("/crontasks/:cronTaskScheduleID", controller.GetSpecifiedCrontaskByScheduleID)

		// 查询所有定时任务执行结果
		v1.GET("/crontasksresult", controller.GetAllCronTaskResult)
		// 删除所有定时任务执行结果
		v1.DELETE("/crontasksresult", controller.DeleteAllCronTaskResult)
		// 删除特定定时任务执行结果
		v1.DELETE("/crontasksresult/:cronTaskResultID", controller.DeleteSpecifiedCrontaskResult)

		/***
			定制:以下是测试工具相关接口
		***/
		v1.GET("/custom/request/:path", middleware.AllowCookieMiddleWare(), controller.CustomRequest)
		v1.POST("/custom/request/:path", middleware.AllowCookieMiddleWare(), controller.CustomRequest)

		/***
			定制:以下是DSP清除器相关接口
		***/
		v1.DELETE("/custom/cleaner/dsp/deleteRiskAndVunl", controller.DeleteRiskAndVunl)

		/***
			定制:以下是POC配置相关接口
		***/

		// 获取配置
		v1.GET("/custom/pocconfig", controller.GetPOCConfig)
		// 保存配置
		v1.POST("/custom/pocconfig", controller.PostPOCConfig)
		// 获取安装agent的配置
		v1.GET("/custom/installconfig", controller.GetAgentInstallConfig)
		// 保存安装agent的配置
		v1.POST("/custom/installconfig", controller.PostAgentInstallConfig)
		// 执行安装agent的操作
		v1.POST("/custom/installagent", controller.PostInstallAgent)
		// 更新工具盒中的agent
		v1.GET("/custom/updatedscagent", controller.UpdateDSCAgentInToolBox)
		// 更新大脑中的账号提取配置
		v1.GET("/custom/updatedscaccountextract", controller.UpdateDSCAccountExtract)
		// 测试SSH链接是否成功
		v1.POST("/custom/testssh", controller.TestSSHConnection)
		// 调整大脑的阈值
		v1.POST("/custom/modifydscthreshold", controller.ModifyDSCThreshold)

		/***
			定制:以下是泄密溯源相关接口
		***/
		v1.GET("/custom/dataleakage", controller.DataLeakage)

		/***
			定制:以下是分类分级相关的业务接口
		***/
		v1.POST("/custom/adddataclassify", controller.DataClassify)

		/***
			定制:以下是脆弱性与风险相关的业务接口
		***/

		// 查询所有脆弱性与风险
		v1.GET("/custom/riskandvulnerability", controller.GetAllRiskAndVulnerability)
		// 查询所有脆弱性与风险的触发记录
		v1.GET("/custom/riskandvulnerabilitylog", controller.GetAllRiskAndVulnerabilityLog)
		// 查看工具盒中Agent的配置
		v1.GET("/custom/dscagentconfig", controller.GetDSCAgentConfig)
		// 查看大脑中账号提取的配置
		v1.GET("/custom/getaccountextractconfig", controller.GetAccountExtractConfig)
		// 强制审计日志
		v1.GET("/custom/forceaudit", controller.ForceAudit)

		/***
			定制:以下是修改时间的接口
		***/
		v1.GET("/custom/modifydate/:date", controller.ModifyDate)

		/***
			定制:以下是模拟脆弱性与风险的接口
		***/
		Vulnerability := v1.Group("/custom/mock/vulnerability")
		Vulnerability.Use(middleware.SaveAllTriggerLog())
		// 1. url中存在密码信息
		Vulnerability.GET("/PswdInURL", controller.MockURLContainsPasswd)
		// 2. 响应数据存在密码信息
		Vulnerability.GET("/PswdInResp", controller.MockResponseContainsPasswd)
		// 3. cookie中存在密码信息
		Vulnerability.GET("/PswdInCookie", middleware.AllowCookieMiddleWare(), controller.MockCookieContainsPasswd)
		// 4. 请求数据存在明文密码信息
		Vulnerability.POST("/PswdInPlainText", controller.MockPlainTextContainsPasswd)
		// 5. 登录弱密码
		Vulnerability.POST("/WeakPasswd", controller.MockWeakPasswd)
		// 6. 响应数据存在明文密码信息
		Vulnerability.GET("/PswdInPlainTextResp", controller.MockPlainTextResponseContainsPasswd)
		// 7. GET方式执行危险操作
		Vulnerability.GET("/DeleteInGet", controller.MockDeleteInGet)
		// 8. 鉴权信息在url中
		Vulnerability.GET("/AuthInURL", controller.MockAuthInURL)
		// 9. 敏感信息在url中
		Vulnerability.GET("/SensitiveDataInURL", controller.MockSensitiveDataInURL)
		// 10. 敏感接口未鉴权
		Vulnerability.POST("/SensitiveAPINotSec", controller.MockSensitiveAPINotAuth)
		// 11. 非敏感接口未鉴权
		Vulnerability.POST("/NoneSensitiveAPINotSec", controller.MockNoneSensitiveAPINotAuth)
		// 12. 敏感接口返回数据量可修改
		Vulnerability.GET("/DataAmountCanBeModified", controller.MockDataAmountCanBeModified)
		// 13. 单次访问数据量过大
		Vulnerability.GET("/TooMuchDataInSingleRequest", controller.MockTooMuchDataInSingleRequest)
		// 14. 单次访问敏感类型过多
		Vulnerability.GET("/TooMuchTypeInSingleRequest", controller.MockTooMuchTypeInSingleRequest)
		// 15. 脱敏策略不一致
		Vulnerability.GET("/DifferentDesensePolicy", controller.MockDifferentDesensePolicy)
		// 16. Hadoop未授权访问
		Vulnerability.POST("/ws/v1/cluster/apps/new-application", controller.MockHadoopUnAuthorizedAccess)
		// 17. SonarQube未授权访问
		r.GET("/api/settings/values", controller.MockSonarQubeUnAuthorizedAccess)
		// 18. ssh secret key信息泄露
		Vulnerability.POST("/fileDownload", controller.MockSSHSecretKeyLeakage)
		// 19. JDBC连接字符串信息泄露
		Vulnerability.POST("/xxl-conf-admin/conf/find", controller.MockJDBCStringLeakage)
		// 20. Alibaba Nacos未授权访问
		Vulnerability.POST("/nacos/v1/auth/users", controller.MockHAlibabaNacosUnAuthorizedAccess)

		Risk := v1.Group("/custom/mock/risk")
		Risk.Use(middleware.SaveAllTriggerLog())

		// 1. 账号多地访问
		Risk.GET("/OneAccountWithMultiPlace", controller.OneAccountWithMultiPlace)
		// 2. 账号多IP访问
		Risk.GET("/OneAccountWithMultiIP", controller.OneAccountWithMultiIP)
		// 3. 境内IP有多个账号身份
		Risk.GET("/LocalIPWithMultiAccount", controller.LocalIPWithMultiAccount)
		// 4. 境外IP有多个账号身份
		Risk.GET("/ForeignIPWithMultiAccount", controller.ForeignIPWithMultiAccount)
		// 5. 单个账号一段时间内返回大量敏感数据
		Risk.GET("/SingleAccountReturnTooMuchSensitiveDataPeriod", controller.MockSingleAccountReturnTooMuchSensitiveDataPeriod)
		// 6. 单个IP一段时间内返回大量敏感数据
		Risk.GET("/SingleIPReturnTooMuchSensitiveDataPeriod", controller.MockSingleIPReturnTooMuchSensitiveDataPeriod)
		// 7. 单个账号单次返回大量敏感数据
		Risk.GET("/SingleAccountReturnTooMuchSensitiveDataOnce", controller.MockSingleAccountReturnTooMuchSensitiveDataOnce)
		// 8. 单个IP单次返回大量敏感数据
		Risk.GET("/SingleIPReturnTooMuchSensitiveDataOnce", controller.MockSingleIPReturnTooMuchSensitiveDataOnce)
		// 9. 单个账号单次返回敏感数据类型超过15种
		Risk.GET("/SingleAccountReturnTooManyKindsOfSensitiveDataOnce", controller.MockSingleAccountReturnTooManyKindsOfSensitiveDataOnce)
		// 10. 单个IP单次返回敏感数据类型超过15种
		Risk.GET("/SingleIPReturnTooManyKindsOfSensitiveDataOnce", controller.MockSingleIPReturnTooManyKindsOfSensitiveDataOnce)

		// 11. 单个账号单次返回新类型的敏感数据
		Risk.GET("/SingleAccountReturnNewTypeSensiDataOnce", controller.SingleAccountReturnNewTypeSensiDataOnce)
		// 12. 单个IP单次返回新类型的敏感数据
		Risk.GET("/SingleIPReturnNewTypeSensiDataOnce", controller.SingleIPReturnNewTypeSensiDataOnce)
		// 13. 单个账号在一段时间内进行请求参数值遍历
		Risk.GET("/SingleAccountRequestTraversePeriod", controller.SingleAccountRequestTraversePeriod)
		// 14. 单个IP在一段时间内进行请求参数值遍历
		Risk.GET("/SingleIPRequestTraversePeriod", controller.SingleIPRequestTraversePeriod)
		// 15. 请求参数值出现新类型
		Risk.GET("/NewTypeInRequest", controller.NewTypeInRequest)
		// 16. 请求方法异常
		Risk.GET("/AbnormalRequestMethod", controller.AbnormalRequestMethod)
		// 17. 单个账号在一段时间内返回大量4XX
		Risk.GET("/SingleAccountReturnTooMuch4XXPeriod", controller.SingleAccountReturnTooMuch4XXPeriod)
		// 18. 单个IP在一段时间内返回大量4XX
		Risk.GET("/SingleIPReturnTooMuch4XXPeriod", controller.SingleIPReturnTooMuch4XXPeriod)
		// 19. 单个账号在一段时间内频繁访问同一API
		Risk.GET("/SingleAccountVisitSameAPIPeriod", controller.SingleAccountVisitSameAPIPeriod)
		// 20. 单个账号在异常时间段频繁访问同一API
		Risk.GET("/SingleAccountVisitSameAPIAbnormalPeriod", controller.SingleAccountVisitSameAPIAbnormalPeriod)
		// 21. 单个IP在一段时间内频繁访问同一API
		Risk.GET("/SingleIPVisitSameAPIPeriod", controller.SingleIPVisitSameAPIPeriod)
		// 22. 单个IP在异常时间段频繁访问同一API
		Risk.GET("/SingleIPVisitSameAPIAbnormalPeriod", controller.SingleIPVisitSameAPIAbnormalPeriod)
		// 23. 请求参数名缺失
		Risk.GET("/LackOfVarName", controller.LackOfVarName)
		// 24. 请求参数出现非预期的参数名
		Risk.GET("/UnexpectedVarName", controller.UnexpectedVarName)
		// 25. 单个账号在一段时间内进行路径遍历
		Risk.GET("/SingleAccountPathTraversePeriod/:name", controller.SingleAccountPathTraversePeriod)
		// 26. 单个IP在一段时间内进行路径遍历
		Risk.GET("/SingleIPPathTraversePeriod/:name", controller.SingleIPPathTraversePeriod)

		/***
			定制:以下是安全事件相关的业务接口
		***/

		// 查询所有脆弱性与风险
		v1.GET("/custom/securityevents", controller.GetAllSecurityEvents)

		SecurityEvents := v1.Group("/custom/mock/securityevents")

		// 参数遍历获取大量敏感数据
		SecurityEvents.GET("/RequestTraverseAndReturnTooMuchSensitiveData", controller.RequestTraverseAndReturnTooMuchSensitiveData)
	}
	return r
}
