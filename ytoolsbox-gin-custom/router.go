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

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// 定义一个结构体来表示请求和响应信息
type RequestResponse struct {
	Request struct {
		Method string                 `json:"method"`
		Body   map[string]interface{} `json:"body"`
	} `json:"request"`
	Response struct {
		Status int                    `json:"status"`
		Body   map[string]interface{} `json:"body"`
	} `json:"response"`
}

// 从JSON文件中读取路由配置并注册
func RegisterRoutesFromJSON(r *gin.RouterGroup, filePath string) {
	// 读取JSON文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}

	// 打印文件路径和内容
	// log.Printf("Reading routes from file: %s", filePath)
	// log.Printf("File content: %s", string(data))

	// 解析JSON数据
	var config map[string]RequestResponse
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("JSON解析错误: %v", err)
	}
 	// 打印解析后的数据
    // log.Printf("Parsed config: %+v", config)
	// 循环遍历请求信息并注册路由
	for url, reqRes := range config {
		switch reqRes.Request.Method {
		case "GET":
			r.GET(url, controller.HandleDynamicRequest)
		case "POST":
			r.POST(url, controller.HandleDynamicRequest)
		case "PUT":
			r.PUT(url, controller.HandleDynamicRequest)
		case "DELETE":
			r.DELETE(url, controller.HandleDynamicRequest)
		default:
			fmt.Printf("不支持的HTTP方法: %s\n", reqRes.Request.Method)
		}
		// 打印注册的路由
        log.Printf("Registering route: %s %s", reqRes.Request.Method, url)
	}
}

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.SaveAllTriggerLog())
	URL_Prefix := "/api/auth"
	/* 用路由组重新归纳了一下路由 */
	v1 := r.Group(URL_Prefix)
	{
		// 调用函数从JSON文件,动态注册路由
		RegisterRoutesFromJSON(v1, "config/multiple_http_json/http.json")

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
		// 5. 登录弱密码 或 存在空密码/空密钥问题 或 万能密码注入
		Vulnerability.POST("/WeakPasswd", controller.MockWeakPasswd)
		// 6. 响应数据存在明文密码信息
		Vulnerability.GET("/PswdInPlainTextResp", controller.MockPlainTextResponseContainsPasswd)
		// 7. GET方式执行危险操作r
		Vulnerability.GET("/DeleteInGet", controller.MockDeleteInGet)
		// 8. 鉴权信息在url中
		Vulnerability.GET("/AuthInURL", controller.MockAuthInURL)
		// 9. 敏感信息在url中
		Vulnerability.GET("/SensitiveDataInURL", controller.MockSensitiveDataInURL)
		// 10. 敏感接口未鉴权
		r.POST("/api/custom/mock/vulnerability/SensitiveAPINotSec", controller.MockSensitiveAPINotAuth)
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
		// 新增54个脆弱性检测
		// // 21. ElasticSearch根路径未授权访问 和 凭证长期生效
		// Vulnerability.GET("/", controller.MockElasticSearchRootPathUnauthorizedOrPermanentValidityOfCertificate)
		// 22. ElasticSearch查询数据接口未授权访问
		Vulnerability.GET("/baizhi/_search", controller.MockElasticSearchSearchDataAPIUnauthorized)
		// 23. ElasticSearch节点目录未授权访问
		Vulnerability.GET("/_nodes", controller.MockElasticSearchNodePathUnauthorized)
		// // 24. Docker部分路径未授权访问
		// Vulnerability.GET("/v1.40/images/json", controller.MockDockerUnauth)
		// 25. SpringBoot部分常见路径未授权访问
		Vulnerability.GET("/actuator/info", controller.MockSpringBootUnauth)
		// 26. Jolokia未授权访问
		Vulnerability.GET("/jolokia/list", controller.MockJolokiaUnauth)
		// 27. SonarQube未授权访问应用列表
		r.GET("/api/webservices/list", controller.MockSonarQubeUnauth)
		// 28. Swagger未授权访问
		Vulnerability.GET("/v2/api-docs", controller.MockSwaggerUnauth)
		// 29. Solr未授权访问
		Vulnerability.GET("/solr/admin/info/system", controller.MockSolrUnauth)
		// 30. 云主机Accesskey或Secret key泄漏
		Vulnerability.GET("/GetDefualtOssInfo", controller.MockAKSKReveal)
		// 31. Kubernetes API Server未授权访问
		Vulnerability.GET("/api/v1/namespaces/pods", controller.MockKubernetesApiServerUnAuthorizedAccess)
		// // 32. Kubelet 未授权访问
		// Vulnerability.GET("/pods/", controller.MockKubeletUnAuthorizedAccess)
		// 33. Kubernetes Dashboard面板未授权访问
		Vulnerability.GET("/api/v1/settings/pinner", controller.MockKubernetesDashboardUnAuthorizedAccess)
		// 34. GraphQL未授权访问
		Vulnerability.POST("/graphql", controller.MockGraphQLUnAuthorizedAccess)
		// 35. 短信验证码出现在响应包中
		Vulnerability.POST("/sendmessage", controller.MockVarifyCodeInResponse)
		// // 36. Kong未授权访问漏洞（CVE-2020-11710）
		// Vulnerability.POST("/services", controller.MockKongUnauthorizedAccess)
		// 37. Weblogic未授权访问漏洞
		Vulnerability.POST("/ws_utc/resources/setting/keystore", controller.MockWeblogicUnauthorizedAccess)
		// 38. Apache APISIX Dashboard 认证绕过（CVE-2021-45232）
		Vulnerability.GET("/apisix/admin/migrate/export", controller.MockApacheAPISIXDashboardUnauthorizedAccess)
		// 39. Apache ShenYu Admin 身份验证绕过漏洞（CVE-2021-37580）
		Vulnerability.GET("/dashboardUser", controller.MockApacheShenYuAdminAPIUnauthorizedAccess)
		// 40. Cookie没有添加HttpOnly属性
		Vulnerability.GET("/php/login.php", controller.MockCookieDoesNotHaveHttpOnlyPropertyConfigured)
		// // 41. 开发环境涉敏API在公网暴露
		// Vulnerability.GET("/api/getUser", controller.MockTestApiPublic)
		// 42. JWT认证错误地使用了空加密算法
		Vulnerability.GET("/api/islogin", controller.MockJWTUseWrongEncryptionAlgorithm)
		// 43.

		// 44. API接口存在水平越权

		// 45. 涉敏API存在资源未限速

		// // 46. referer包含敏感信息
		// Vulnerability.GET("/main/index.php", controller.MockRefererContainsSensitiveData)
		// 47. 凭证长期生效(ElasticSearch根路径未授权访问)
		Vulnerability.GET("/", controller.MockElasticSearchRootPathUnauthorizedOrPermanentValidityOfCertificate)
		// 48. 危险方法未被禁用
		Vulnerability.DELETE("/api/user", controller.MockDangerousMethodsNotBanned)
		// // // 49. 存在空密码/空密钥问题
		// Vulnerability.POST("/WeakPasswd", controller.MockWeakPasswd)
		// // 50. 万能密码注入
		// Vulnerability.POST("/WeakPasswd", controller.MockWeakPasswd)
		// 51. MessageSolution 企业邮件归档管理系统 EEA 存在信息泄露漏洞
		Vulnerability.GET("/authenticationserverservlet", controller.MockSensitiveFileInformationDisclosure)
		// 52. FineReport管理平台privilege.xml文件信息泄露
		Vulnerability.GET("/WebReport/ReportServer", controller.MockPrivilegeXmlFileInformationDisclosure)
		// 53~57. Tomcat tomcat-users.xml 或 Apache Axis axis2.xml 或 Glassfish domain.xml 或 Jenkins config.xml 或 Zimbra conf/localconfig.xml信息泄露
		Vulnerability.GET("/a.php", controller.MockInformationDisclosureDetectedOnTomcatOrAxis2OrDomainOrConfigOrLocalconfigXml)
		// 58. F5 BIG-IP 认证绕过（CVE-2020-5902）
		Vulnerability.GET("/tmui/login.jsp/..;/tmui/locallb/workspace/a.php", controller.MockF5BIGIPAuthenticationBypass5902)
		// 59.69. F5 BIG-IP 认证绕过（CVE-2022-1388）或 F5 BIG-IP命令注入漏洞（CVE-2021-22986）
		Vulnerability.POST("/mgmt/tm/util/bash", controller.MockF5BIGIPAuthenticationBypassOrRceVulnerabilityAttack)
		// 60. 绿盟UTS综合威胁探针信息泄露
		Vulnerability.GET("/webapi/v1/system/accountmanage/account", controller.MockPageInformationDisclosureDetectedOnNSFOCUSUTSComprehensiveThreatProbe)
		// 61. Couchdb垂直权限绕过
		Vulnerability.PUT("/_users/org.couchdb.user:vulhub", controller.MockApacheCouchdbRemotePrivilegeEscalation)
		// 62. XXL-Job远程命令执行漏洞
		Vulnerability.POST("/run", controller.MockXXLJOBRemoteCodeExecutionVulnerability)
		// 63. Spring Cloud Gateway代码注入漏洞（CVE-2022-22947）
		Vulnerability.POST("/actuator/gateway/routes/hacktest", controller.MockSpringCloudGatewayActuatorAPISpELCodeInjectionVulnerabilityAttack)
		// 64. YApi未授权代码执行
		Vulnerability.POST("/api/project/up", controller.MockYapiConsoleRceVulnerabilityAttack)
		// 65. Coremail未授权访问（CNVD-2019-78549）
		Vulnerability.GET("/mailsms/s", controller.MockCoremailArbitraryFileReadVulnerabilityAttack)
		// 66. Druid敏感信息泄露
		Vulnerability.POST("/druid/sql.json", controller.MockDruidSensitiveInfomationUnauthorizedAccessAttack)
		// 67. Atlassian jira信息泄露漏洞(CVE-2019-8449)
		Vulnerability.GET("/rest/api/latest/groupuserpicker", controller.MockAtlassianJiraInformationDisclosureVulnerabilityAttack)
		// 68. spark未授权命令执行漏洞
		Vulnerability.POST("/v1/submissions/create", controller.MockSparkUnauthorizedAccessAttack)
		// 69. F5 BIG-IP命令注入漏洞（CVE-2021-22986）

		// 70. spring data rest 远程命令执行（CVE-2017-8046）
		Vulnerability.PATCH("/customers/1", controller.MockSpringDataRESTRemoteCodeExecutionVulnerabilityAttack)
		// 71. Apache APISIX 远程代码执行漏洞（CVE-2022-24112）
		Vulnerability.POST("/apisix/batch-requests", controller.MockApacheAPISIXRemoteCodeExecutionVulnerability)
		// 72. kylin远程命令执行漏洞（CVE-2020-1956）
		Vulnerability.PUT("/kylin/api/cubes/kylin_streaming_cube/learn_kylin/migrate", controller.MockApacheKylinCommandInjectionVulnerability)
		// 73. k8s Dashboard面板认证绕过漏洞（CVE-2018-18264）
		Vulnerability.GET("/api/v1/secret/kube-system/kubernetes-dashboard-certs", controller.MockKubernetesDashboardAuthorizationBypass)
		// 74. Cisco Data Center Network Manager 认证绕过漏洞（CVE-2019-15977）
		Vulnerability.POST("/DbAdminWSService/DbAdminWS", controller.MockCiscoDataCenterNetworkManagerAuthorizationBypass)
		// 75. Zimbra 组件漏洞信息泄露
		Vulnerability.POST("/service/soap", controller.MockTokenOfLeastPrivilegeAccounDisclosureDetectedOnZimbra)
		// 76. API接口文档未授权访问
		Vulnerability.GET("/openapi.json", controller.MockAPIDefinitionDocumentsCauseInterfaceInformationDisclosure)
		// // 77. CORS配置不当导致敏感信息泄露
		// r.GET("/get", controller.MockSensitiveInformationDisclosureDueToImproperCORSConfiguration)
		// 78. HTTP Basic Authorization弱口令
		Vulnerability.GET("/user", controller.MockHTTPBasicAuthorizationWeakPassword)
		// 79. 重要接口未授权访问
		Vulnerability.GET("/user/admin/infolog.sql", controller.MockUnauthorizedAccessToImportantInterfaces)

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
		Risk.POST("/SingleAccountReturnTooMuchSensitiveDataPeriod", controller.MockSingleAccountReturnTooMuchSensitiveDataPeriod)
		// 6. 单个IP一段时间内返回大量敏感数据
		Risk.GET("/SingleIPReturnTooMuchSensitiveDataPeriod", controller.MockSingleIPReturnTooMuchSensitiveDataPeriod)
		// 7. 单个账号单次返回大量敏感数据
		Risk.POST("/SingleAccountReturnTooMuchSensitiveDataOnce", controller.MockSingleAccountReturnTooMuchSensitiveDataOnce)
		// 8. 单个IP单次返回大量敏感数据
		Risk.GET("/SingleIPReturnTooMuchSensitiveDataOnce", controller.MockSingleIPReturnTooMuchSensitiveDataOnce)
		// 9. 单个账号单次返回敏感数据类型超过15种
		Risk.POST("/SingleAccountReturnTooManyKindsOfSensitiveDataOnce", controller.MockSingleAccountReturnTooManyKindsOfSensitiveDataOnce)
		// 10. 单个IP单次返回敏感数据类型超过15种
		Risk.GET("/SingleIPReturnTooManyKindsOfSensitiveDataOnce", controller.MockSingleIPReturnTooManyKindsOfSensitiveDataOnce)

		// 11. 单个账号单次返回新类型的敏感数据
		Risk.GET("/SingleAccountReturnNewTypeSensiDataOnce", controller.SingleAccountReturnNewTypeSensiDataOnce)
		//12. 单个IP单次返回新类型的敏感数据
		Risk.GET("/SingleIPReturnNewTypeSensiDataOnce", controller.SingleIPReturnNewTypeSensiDataOnce)
		// 13. 单个账号在一段时间内进行请求参数值遍历
		Risk.POST("/SingleAccountRequestTraversePeriod", controller.SingleAccountRequestTraversePeriod)
		// 14. 单个IP在一段时间内进行请求参数值遍历
		Risk.GET("/SingleIPRequestTraversePeriod", controller.SingleIPRequestTraversePeriod)
		// 15. 请求参数值出现新类型
		Risk.GET("/NewTypeInRequest", controller.NewTypeInRequest)
		// 16. 请求方法异常
		Risk.GET("/AbnormalRequestMethod", controller.AbnormalRequestMethod)
		// 17. 单个账号在一段时间内返回大量4XX
		Risk.POST("/SingleAccountReturnTooMuch4XXPeriod", controller.SingleAccountReturnTooMuch4XXPeriod)
		// 18. 单个IP在一段时间内返回大量4XX
		Risk.GET("/SingleIPReturnTooMuch4XXPeriod", controller.SingleIPReturnTooMuch4XXPeriod)
		// 19. 单个账号在一段时间内频繁访问同一API
		Risk.POST("/SingleAccountVisitSameAPIPeriod", controller.SingleAccountVisitSameAPIPeriod)
		// 20. 单个账号在异常时间段频繁访问同一API
		Risk.POST("/SingleAccountVisitSameAPIAbnormalPeriod", controller.SingleAccountVisitSameAPIAbnormalPeriod)
		// 21. 单个IP在一段时间内频繁访问同一API
		Risk.GET("/SingleIPVisitSameAPIPeriod", controller.SingleIPVisitSameAPIPeriod)
		// 22. 单个IP在异常时间段频繁访问同一API
		Risk.GET("/SingleIPVisitSameAPIAbnormalPeriod", controller.SingleIPVisitSameAPIAbnormalPeriod)
		// 23. 请求参数名缺失
		Risk.GET("/LackOfVarName", controller.LackOfVarName)
		// 24. 请求参数出现非预期的参数名
		Risk.GET("/UnexpectedVarName", controller.UnexpectedVarName)
		// 25. 单个账号在一段时间内进行路径遍历
		Risk.POST("/SingleAccountPathTraversePeriod/:name", controller.SingleAccountPathTraversePeriod)
		// 26. 单个IP在一段时间内进行路径遍历
		Risk.GET("/SingleIPPathTraversePeriod/:name", controller.SingleIPPathTraversePeriod)

		//dspm
		//个人访问任一应用的数据量超过历史基线
		Risk.POST("/PersonGetAppDataMoreThanBaseLine", controller.PersonGetAppDataMoreThanBaseLine)

		//ip短时间下载了大量的数据
		//IPReturnsLargeAmountSensitiveDataInShortTime
		Risk.POST("/IPReturnsLargeAmountSensitiveDataInShortTime", controller.IPReturnsLargeAmountSensitiveDataInShortTime)

		//dspm-账号爆破成功
		Risk.POST("/ParameterIterationLogin", controller.ParameterIteration)
		//dspm-账号共享
		Risk.POST("/AccountShare", controller.AccountShare)

		//dspm-应用的账号数超过历史基线
		Risk.POST("/AppAccountMoreThanBaseLine", controller.AppAccountMoreThanBaseLine)

		/***
			定制:以下是安全事件相关的业务接口
		***/

		//安全事件风险时间修改接口
		v1.GET("/custom/modifysystemdate", controller.SecurityEventsModifyTime)

		// 查询所有脆弱性与风险
		v1.GET("/custom/securityevents", controller.GetAllSecurityEvents)

		SecurityEvents := v1.Group("/custom/mock/securityevents")

		// 参数遍历获取大量敏感数据 或 频繁访问获取大量敏感数据
		SecurityEvents.GET("/RequestTraverseAndReturnTooMuchSensitiveData", controller.RequestTraverseAndReturnTooMuchSensitiveData)

		// 异常时间段频繁访问获取大量敏感数据

		// 发生探测攻击并通过参数遍历获取过量敏感数据

		// 发生探测攻击并通过频繁访问获取过量敏感数据

		// 发生探测攻击并在异常时间段频繁访问获取非预期敏感数据

		// 通过恶意构造请求窃取额外敏感数据

		// API接口遭遇渗透攻击

		// 通过恶意构造请求获取大量敏感数据

		// 账号失陷并下载了大量敏感数据
		SecurityEvents.GET("/OneAccountGetSensitiveData", controller.OneAccountGetSensitiveData)

	}
	return r
}
