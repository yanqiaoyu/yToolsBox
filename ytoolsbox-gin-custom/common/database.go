/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 14:26:36
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-06-22 19:16:11
 * @FilePath: \golang_web\common\database.go
 */

package common

import (
	"flag"
	"fmt"
	"main/model"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

// 工具盒业务的句柄
var DB *gorm.DB

// 泄密溯源的句柄
var DataLeakgeDB *gorm.DB

var RiskAndVulnerabilityList = []model.RiskAndVulnerability{
	{
		Name:          "数据识别的请求与响应", // 点击下拉框之后点击其中一个行业的名称，对应构建发送这个行业对应的http.json文件内容
		Type:          "数据识别",
		Desc:          "从浏览器发送http请求到POC工具的后台服务器。这里的报文包含了1个应用资产，30个API资产，合计1120次API请求。每个http请求对应一个特定的请求和响应报文，涵盖了用户、账户、交易、投资、贷款、合同、客户、公司、账单、报告、风险等多个数据类型。",
		Level:         "低危",
		TriggerMethod: "批量下发http请求，读取对应的http.json中的信息。",
	},
	{
		Name:          "url中存在密码信息",
		Type:          "脆弱性",
		Desc:          "接口认证方式不合理,登录功能采用GET方法传输账号密码,例如黑客可以通过中间人攻击等方式捕获用户的账号密码,导致账号密码泄露",
		Level:         "中危",
		TriggerMethod: "使用GET方法访问/api/auth/custom/mock/vulnerability/PswdInURL?password=123456,由于该URL携带了password=123456, 因此可以触发<url中存在密码信息>这个脆弱性",
	},
	{
		Name:          "响应数据存在密码信息",
		Type:          "脆弱性",
		Desc:          "接口响应不合理,响应内容中包含了密码,例如黑客可以通过中间人攻击等方式捕获用户的账号密码,导致密码泄露",
		Level:         "中危",
		TriggerMethod: "使用GET方法访问/api/auth/custom/mock/vulnerability/PswdInResp,由于该接口的响应中含有 password=123456,因此可以触发<响应数据存在密码信息>这个脆弱性",
	},
	{
		Name:          "cookie中存在密码信息",
		Type:          "脆弱性",
		Desc:          "认证规则不合理,将用户密码存储在cookie中,例如黑客可以通过中间人攻击等方式捕获用户的账号密码,导致账号被盗取",
		Level:         "中危",
		TriggerMethod: "使用GET方法访问/api/auth/custom/mock/vulnerability/PswdInCookie,由于该请求中含有cookie=123456,因此可以触发<cookie中存在密码信息>这个脆弱性",
	},
	{
		Name:          "请求数据存在明文密码信息",
		Type:          "脆弱性",
		Desc:          "接口传输不规范,传输过程中未对密码进行加密,例如黑客可以通过中间人攻击等方式捕获用户的账号密码,导致用户账号被窃取",
		Level:         "高危",
		TriggerMethod: "使用POST方法访问/api/auth/custom/mock/vulnerability/PswdInPlainText,由于请求体中含有password=123456,因此可以触发<请求数据存在明文密码信息>这个脆弱性",
	},
	{
		Name:          "登录弱密码",
		Type:          "脆弱性",
		Desc:          "密码设置不符合信息安全规范(比如为简单的数字字符组合),例如黑客可以通过口令爆破等方式猜解用户账号密码,导致用户账号被盗取",
		Level:         "高危",
		TriggerMethod: "使用POST方法访问/api/auth/custom/mock/vulnerability/WeakPasswd?reqType=WeakPassword,由于请求body中含有username=admin,password=admin,并且配置了相应的账号提取规则,因此可以触发<登录弱密码>这个脆弱性",
	},
	{
		Name:          "响应数据存在明文密码信息",
		Type:          "脆弱性",
		Desc:          "接口响应不规范,响应内容中包含了未加密的密码,例如黑客可以通过中间人劫持等方式获取用户的账号密码,导致密码泄露",
		Level:         "高危",
		TriggerMethod: "使用GET方法访问/api/auth/custom/mock/vulnerability/PswdInPlainTextResp,由于该接口的响应中含有 password=123456,因此可以触发<响应数据存在明文密码信息>这个脆弱性",
	},
	{
		Name:          "GET方式执行危险操作",
		Type:          "脆弱性",
		Desc:          "传输方法不合理,使用GET执行敏感操作,例如用户在未知情况下点击黑客构造的恶意url,导致自动进行危险操作",
		Level:         "低危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/DeleteInGet?delete=1,由于该URL携带了delete=1,因此可以触发<GET方式执行危险操作>这个脆弱性",
	},
	{
		Name:          "鉴权信息在url中",
		Type:          "脆弱性",
		Desc:          "权限认证方式不合理,采用GET方法传输鉴权信息,例如黑客可以通过中间人攻击等方式捕获用户的鉴权信息(token, sessionid等),导致账号被盗取",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/AuthInURL?itembox=&sessionReset=ST-1211-iBHOcPmJtfHGjrXMlx7c-sso01,由于该URL携带了sessionReset,因此可以触发<鉴权信息在url中>这个脆弱性",
	},
	{
		Name:          "敏感信息在url中",
		Type:          "脆弱性",
		Desc:          "接口传输不规范,传输过程中未对敏感数据进行加密,例如黑客可以通过中间人攻击等方式捕获用户的敏感信息(身份证、电话号码等),导致用户信息泄露",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/SensitiveDataInURL?phone=13687148956,因此可以触发<敏感信息在url中>这个脆弱性",
	},
	{
		Name:          "敏感接口未鉴权",
		Type:          "脆弱性",
		Desc:          "接口权限分配不规范,没有对接口进行权限限制,例如黑客通过低用户权限访问敏感接口,获取大量敏感数据,或者执行敏感操作",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/SensitiveAPINotSec,由于该请求中不包含任何鉴权信息,例如:token,session,authorize等字段,并且返回中含有敏感信息name=张三,因此可以触发<敏感接口未鉴权>该脆弱性",
	},
	{
		Name:          "非敏感接口未鉴权",
		Type:          "脆弱性",
		Desc:          "接口权限分配不规范,没有对接口进行权限限制,例如黑客通过低用户权限访问非敏感接口,获取接口信息",
		Level:         "低危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/NoneSensitiveAPINotSec,由于该请求中不包含任何鉴权信息,例如:token,session,authorize等字段,并且返回中不包含敏感信息,因此可以触发<非敏感接口未鉴权>该脆弱性",
	},
	{
		Name:          "敏感接口返回数据量可修改",
		Type:          "脆弱性",
		Desc:          "接口响应不规范,响应内容的数量可以被修改,例如黑客可以通过修改参数值等方式获取过量的敏感数据,造成大量数据泄露",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/DataAmountCanBeModified?query=&pagenum=1&pagesize=10,由于请求的URL中含有pagenum,pagesize等关键字,因此可以触发<敏感接口返回数据量可修改>该脆弱性",
	},
	{
		Name:          "单次访问数据量过大",
		Type:          "脆弱性",
		Desc:          "业务设计不合理,例如黑客可以通过sql注入等方式获取过量敏感数据,造成大量数据泄露",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/TooMuchDataInSingleRequest,由于返回的响应中存在200敏感数据,触发阈值为100,因此可以触发<单次访问数据量过大>该脆弱性",
	},
	{
		Name:          "单次访问敏感类型过多",
		Type:          "脆弱性",
		Desc:          "业务设计不合理,没有进行敏感数据过滤,例如黑客可以通过添加参数名等方式获取过多的敏感类型数据,导致敏感数据泄露",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/TooMuchTypeInSingleRequest,由于返回的响应中存在13种类型的敏感数据,触发阈值为5,因此可以触发<单次访问敏感类型过多>该脆弱性",
	},
	{
		Name:          "脱敏策略不一致",
		Type:          "脆弱性",
		Desc:          "接口响应不规范, 例如黑客可以通过中间人攻击等方式捕获用户的敏感信息,导致敏感信息被盗取",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/DifferentDesensePolicy,由于返回的响应中同时存在脱敏与未脱敏的数据,因此可以触发<脱敏策略不一致>该脆弱性",
	},
	{
		Name:          "Hadoop未授权访问",
		Type:          "脆弱性",
		Desc:          "攻击者未经授权获取Hadoop服务敏感信息和做敏感操作,且该页面存在系统命令执行漏洞风险。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
	},
	{
		Name:          "SonarQube未授权访问",
		Type:          "脆弱性",
		Desc:          "SonarQube系统在默认配置下,会将通过审计的源代码上传至SonarQube平台。由于SonarQube缺少对API接口访问的鉴权控制,攻击者利用该漏洞,可在未授权的情况下通过访问上述API接口,获取SonarQube平台上的程序源代码,构成项目源代码数据泄露风险。",
		Level:         "中危",
		TriggerMethod: "/api/settings/values",
	},
	{
		Name:          "ssh secret key信息泄露",
		Type:          "脆弱性",
		Desc:          "ssh私钥默认保存在ssh目录下的id_rsa文件中,如果该私钥泄露,黑客可以利用其攻陷主机。如果攻击者利用成功,将会有完整SSH访问权限,能够运行任何命令。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/fileDownload",
	},
	{
		Name:          "JDBC连接字符串信息泄露",
		Type:          "脆弱性",
		Desc:          "JDBC连接字符串含有数据库连接地址、数据库账号、密码信息。如果JDBC连接字符串泄露,黑客可利用其获得数据库权限。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/xxl-conf-admin/conf/find",
	},
	{
		Name:          "Alibaba Nacos未授权访问",
		Type:          "脆弱性",
		Desc:          "Nacos 官方github在2020年12月29日发布的issue中披露Alibaba Nacos 存在一个由于不当处理User-Agent导致的未授权访问漏洞。通过该漏洞,攻击者可以进行任意操作,包括创建新用户并进行登录后操作。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/nacos/v1/auth/users?username=yangy1&password=yangy",
	},
	// {
	// 	Name:		   "ElasticSearch根路径未授权访问",
	// 	Type:          "脆弱性",
	// 	Desc:          "该页面存在集群名称、集群版本等敏感信息,易造成信息泄漏或进一步的攻击行为。",
	// 	Level:         "高危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/?reqType=ElasticSearch",
	// },
	{
		Name:          "ElasticSearch查询数据接口未授权访问",
		Type:          "脆弱性",
		Desc:          "该ElasticSearch查询数据接口允许攻击者未授权访问,并将查询结果返回,造成数据泄漏。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/baizhi/_search/",
	},
	{
		Name:          "ElasticSearch节点目录未授权访问",
		Type:          "脆弱性",
		Desc:          "该ElasticSearch查询数据接口允许攻击者未授权访问，并将查询结果返回，造成数据泄漏。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/_nodes",
	},
	// {
	// 	Name:          "Docker部分路径未授权访问",
	// 	Type:          "脆弱性",
	// 	Desc:          "Docker部分常见路径未授权访问攻击成功,可通过该路径进行敏感操作",
	// 	Level:         "高危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/v1.40/images/json",
	// },
	{
		Name:          "SpringBoot部分常见路径未授权访问",
		Type:          "脆弱性",
		Desc:          "SpringBoot部分常见路径未授权访问攻击成功,可造成应用信息泄漏",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/actuator/info",
	},
	{
		Name:          "Jolokia未授权访问",
		Type:          "脆弱性",
		Desc:          "Jolokia允许对所有已注册的MBean进行HTTP访问,并对其进行敏感操作,可能造成数据泄露甚至远程代码执行",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/jolokia/list",
	},
	{
		Name:          "SonarQube未授权访问应用列表",
		Type:          "脆弱性",
		Desc:          "SonarQube系统在默认配置下,会将通过审计的源代码上传至SonarQube平台。由于SonarQube缺少对API接口访问的鉴权控制,攻击者利用该漏洞,可在未授权的情况下通过访问上述API接口,获取SonarQube平台上的程序源代码,构成项目源代码数据泄露风险。",
		Level:         "高危",
		TriggerMethod: "/api/webservices/list",
	},
	{
		Name:          "Swagger未授权访问",
		Type:          "脆弱性",
		Desc:          "如果生产环境中开启了Swagger功能且未开启认证功能,会导致API接口信息泄露,部分接口可能执行文件上传、查询用户信息等敏感操作,从而导致服务器被未授权访问或越权访问。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/v2/api-docs",
	},
	{
		Name:          "Solr未授权访问",
		Type:          "脆弱性",
		Desc:          "该漏洞允许攻击者查看敏感信息,造成包括系统环境信息、数据库配置信息、数据库数据等敏感信息泄漏,严重时可能造成远程代码执行",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/solr/admin/info/system",
	},
	{
		Name:          "云主机Accesskey或Secret key泄漏",
		Type:          "脆弱性",
		Desc:          "云主机的Accesskey或Secret key信息泄露,可能导致账户下所有的云主机被控制。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/GetDefualtOssInfo",
	},
	{
		Name:          "Kubernetes API Server未授权访问",
		Type:          "脆弱性",
		Desc:          "由于Kubernetes API Server鉴权配置不当,允许匿名用户以管理员权限向集群内部下发指令。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/v1/namespaces/pods",
	},
	// {
	// 	Name:          "Kubelet 未授权访问",
	// 	Type:          "脆弱性",
	// 	Desc:          "由于鉴权配置不当,允许匿名用户身份访问API,以管理员权限向该节点下发指令。",
	// 	Level:         "高危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/pods/",
	// },
	{
		Name:          "Kubernetes Dashboard面板未授权访问",
		Type:          "脆弱性",
		Desc:          "由于鉴权配置不当,从而使Dashboard允许匿名用户以管理员权限向集群内部下发指令。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/v1/settings/pinner",
	},
	{
		Name:          "GraphQL未授权访问",
		Type:          "脆弱性",
		Desc:          "由于GraphQL鉴权配置不当导致未授权访问,这将引入安全漏洞、敏感数据泄露、不安全的对象直接引用甚至SQL或NoSQL进行注入。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/graphql",
	},
	{
		Name:          "短信验证码出现在响应包中",
		Type:          "脆弱性",
		Desc:          "由于短信验证接口存在设计问题,服务器会直接返回短信验证码,并在前端做校验,从而导致验证方式失效。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/sendmessage",
	},
	// {
	// 	Name:          "Kong未授权访问漏洞（CVE-2020-11710）",
	// 	Type:          "脆弱性",
	// 	Desc:          "docker-kong是一款使用在Docker应用容器引擎中的API3网关产品。docker-kong(用于Kong) 2.0.3及之前版本中存在未授权访问漏洞,攻击者可利用该漏洞在127.0.0.1以外的接口上访问admin API端口。",
	// 	Level:         "中危",
	// 	TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/services",
	// },
	{
		Name:          "Weblogic未授权访问漏洞",
		Type:          "脆弱性",
		Desc:          "Weblogic 任意文件上传漏洞是通过ws_utc/config.do路径进行配置的过程中,在配置界面中会存在一个上传点,我们可以更改当前的工作路径,将路径改变成一个我们无需权限即可访问的路径,目的是我们可以不受限制的去访问我们接下来要上传的shell。接下来的步骤就是利用上传点上传Webshell,并在对应路径进行访问。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws_utc/resources/setting/keystore",
	},
	{
		Name:          "Apache APISIX Dashboard 认证绕过（CVE-2021-45232）",
		Type:          "脆弱性",
		Desc:          "Apache APISIX 是一个动态、实时、高性能的 API 网关, 提供负载均衡、动态上游、灰度发布、服务熔断、身份认证、可观测性等丰富的流量管理功能。Apache APISIX Dashboard 使用户可通过前端界面操作 Apache APISIX。该漏洞的存在是由于 Manager API 中的错误。Manager API 在 gin 框架的基础上引入了 droplet 框架,所有的 API 和鉴权中间件都是基于 droplet 框架开发的。但是有些 API 直接使用了框架 gin 的接口,从而绕过身份验证。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/apisix/admin/migrate/export",
	},
	{
		Name:          "Apache ShenYu Admin 身份验证绕过漏洞（CVE-2021-37580）",
		Type:          "脆弱性",
		Desc:          "Apache ShenYu(原名 Soul)是一个异步的、跨语言的、多协议的高性能响应式API 网关,并可应用于所有微服务场景。2021年 11 月 16日,Apache发布安全公告,公开了Apache ShenYu中的一个身份验证绕过漏洞(CVE-2021-37580),该漏洞的CVSS评分为9.8。由于ShenyuAdminBootstrap中JWT的错误使用,导致攻击者可以绕过身份验证,直接进入目标系统后台。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/dashboardUser",
	},
	{
		Name:          "Cookie没有添加HttpOnly属性",
		Type:          "脆弱性",
		Desc:          "Cookie常常用于保存用户的session凭证信息,如果Cookie没有添加HttpOnly属性,可能会造成凭证泄露的风险。",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/php/login.php",
	},
	// {
	// 	Name:          "开发环境涉敏API在公网暴露",
	// 	Type:          "脆弱性",
	// 	Desc:          "开发环境下涉敏API,这些API可能存在潜在风险(如敏感API未鉴权;即便配置了完善的鉴权模块,由于开发API较大可能存在弱口令问题,也会导致该敏感API被利用从而泄露敏感信息)。因此这类API暴露在公网的话易被攻击。",
	// 	Level:         "中危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/getUser",
	// },
	{
		Name:          "JWT认证错误地使用了空加密算法",
		Type:          "脆弱性",
		Desc:          "JWT可以使用不同的算法进行签名,但也可以不签名,在这种情况下,alg参数的值会为none,代表所谓的“不安全的JWT”。如JWT认证错误使用地了空加密算法,攻击者可以伪造任意用户身份登录。",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/islogin",
	},
	// {
	// 	Name:          "referer包含敏感信息",
	// 	Type:          "脆弱性",
	// 	Desc:          "跨域请求的referer中包含上一个url信息，其中包含敏感信息，比如用户密钥、token等，这样会导致业务的敏感数据泄露到第三方手中",
	// 	Level:         "中危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/main/index.php",
	// },
	{
		Name:          "凭证长期生效",
		Type:          "脆弱性",
		Desc:          "凭证生效期已大于7天，长期生效会增加系统的风险，攻击者可以利用被盗的凭证来绕过身份验证。",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/?reqType=Certificate",
	},
	{
		Name:          "危险方法未被禁用",
		Type:          "脆弱性",
		Desc:          "DELETE方法可以用于删除服务器上的资源，如果该方法未被禁用，可能会给系统带来风险。",
		Level:         "中危",
		TriggerMethod: "使用DELETE请求访问/api/auth/custom/mock/vulnerability/api/user",
	},
	{
		Name:          "存在空密码/空密钥问题",
		Type:          "脆弱性",
		Desc:          "存在空密码或空密钥问题会给系统安全带来很大的风险，攻击者可以利用这些漏洞来获取系统中的敏感信息或执行未经授权的操作。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/WeakPasswd?reqType=EmptyPassword",
	},
	{
		Name:          "万能密码注入",
		Type:          "脆弱性",
		Desc:          "万能密码注入是一种常见的安全漏洞，攻击者可以通过这种漏洞来绕过身份认证机制，获得系统权限，从而对系统进行攻击。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/WeakPasswd?reqType=UniversalPassword",
	},
	{
		Name:          "MessageSolution 企业邮件归档管理系统 EEA 存在信息泄露漏洞",
		Type:          "脆弱性",
		Desc:          "存在通用WEB信息泄漏，泄露Windows服务器administrator hash与web账号密码。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/authenticationserverservlet",
	},
	{
		Name:          "FineReport管理平台privilege.xml文件信息泄露",
		Type:          "脆弱性",
		Desc:          "检测发现FineReport管理平台privilege.xml文件存在信息泄露，Finereport 8.0管理平台可以在未授权的状态下进行配置文件读取，攻击者可以通过访问这些文件收集目标系统的敏感信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/WebReport/ReportServer?op=chart&cmd=get_geo_json&resourcepath=privilege.xml",
	},
	{
		Name:          "Tomcat tomcat-users.xml信息泄露",
		Type:          "脆弱性",
		Desc:          "tomcat-users.xml是Tomcat的用户配置文件、包含用户名密码等。它通常位于：/etc/tomcat/tomcat-users.xml，通过该文件，攻击者可以获取Tomcat的用户名密码等信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/a.php?file=tomcat-users.xml",
	},
	{
		Name:          "Apache Axis axis2.xml信息泄露",
		Type:          "脆弱性",
		Desc:          "Apache Axis信息泄露axis2.xml，该文件用与储存Apache axis的用户名密码，通过该文件，攻击者可以获取后台管理员账号密码。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/a.php?file=axis2.xml",
	},
	{
		Name:          "Glassfish domain.xml信息泄露",
		Type:          "脆弱性",
		Desc:          "Glassfish信息泄露domain.xml，该文件是Glassfish的主配置文件，包括认证信息配置，网络信息配置，线程池配置等，通过该文件，攻击者可以获取Glassfish主要配置信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/a.php?file=domain.xml",
	},
	{
		Name:          "Jenkins config.xml信息泄露",
		Type:          "脆弱性",
		Desc:          "/jenkins_home/config.xml 是Jenkins组件的配置文件, 里边包含Jenkins的一些重要配置，通过该文件，攻击者可以获得Jenkins的相关配置信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/a.php?file=config.xml",
	},
	{
		Name:          "Zimbra conf/localconfig.xml信息泄露",
		Type:          "脆弱性",
		Desc:          "conf/localconfig.xml是zimbra的配置文件，其中包含一些相关的配置，通过该文件，攻击者可以获得zimbra的重要配置信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/a.php?file=localconfig.xml",
	},
	{
		Name:          "F5 BIG-IP 认证绕过（CVE-2020-5902）",
		Type:          "脆弱性",
		Desc:          "在F5 BIG-IP产品的流量管理用户页面(TMUI)/配置程序的特定页面中存在一处认证绕过漏洞，导致可以未授权访问TMUI模块所有功能（包括未公开功能），漏洞影响范围包括执行任意系统命令、任意文件读取、任意文件写入、开启/禁用服务等。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/tmui/login.jsp/..;/tmui/locallb/workspace/a.php?fileName=/etc/passwd",
	},
	{
		Name:          "F5 BIG-IP 认证绕过（CVE-2022-1388）",
		Type:          "脆弱性",
		Desc:          "在F5 BIG-IP产品的流量管理用户页面(TMUI)/配置程序的特定页面中存在一处认证绕过漏洞，导致可以未授权访问TMUI模块所有功能（包括未公开功能），漏洞影响范围包括执行任意系统命令、任意文件读取、任意文件写入、开启/禁用服务等。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/mgmt/tm/util/bash?reqType=1388",
	},
	{
		Name:          "绿盟UTS综合威胁探针信息泄露",
		Type:          "脆弱性",
		Desc:          "检测发现绿盟UTS综合威胁探针信息泄露，攻击者可以获取管理员密码。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/webapi/v1/system/accountmanage/account",
	},
	{
		Name:          "Couchdb垂直权限绕过",
		Type:          "脆弱性",
		Desc:          "CVE-2017-12635是由于Erlang和JavaScript对JSON解析方式的不同，导致语句执行产生差异性导致的。这个漏洞可以让任意用户创建管理员，属于垂直权限绕过漏洞。",
		Level:         "高危",
		TriggerMethod: "使用PUT请求访问/api/auth/custom/mock/vulnerability/_users/org.couchdb.user:vulhub",
	},
	{
		Name:          "XXL-Job远程命令执行漏洞",
		Type:          "脆弱性",
		Desc:          "XXL-Job是一个分布式任务调度平台，分为admin和executor两端，executor默认没有配置认证，未授权攻击者可以通过API执行命令，易被攻击者获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/run",
	},
	{
		Name:          "Spring Cloud Gateway代码注入漏洞（CVE-2022-22947）",
		Type:          "脆弱性",
		Desc:          "Spring Cloud Gateway是基于Spring Framework和SpringBoot构建的API网关，当启用Actuator端点时，攻击者可以对使用Spring Cloud Gateway的应用程序执行任意代码，从而获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/actuator/gateway/routes/hacktest",
	},
	{
		Name:          "YApi未授权代码执行",
		Type:          "脆弱性",
		Desc:          "YApi是高效、易用、功能强大的API管理平台，攻击者可以经过注册后利用Mock功能远程执行任意代码，从而获取服务器权限",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/project/up",
	},
	{
		Name:          "Coremail未授权访问（CNVD-2019-78549）",
		Type:          "脆弱性",
		Desc:          "Coremail邮件系统是论客科技有限公司自主研发的大型企业邮件系统，攻击者利用该漏洞可在未授权的情况下访问大部分服务接口，造成信息泄露",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/mailsms/s?func=ADMIN:appState&dumpConfig=/",
	},
	{
		Name:          "Druid敏感信息泄露",
		Type:          "脆弱性",
		Desc:          "Druid是阿里巴巴数据库事业部出品，为监控而生的数据库连接池。当开发者配置不当时，druid存在未授权访问漏洞。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/druid/sql.json??orderBy=SQL&orderType=desc&page=1&perPageCount=1000000&",
	},
	{
		Name:          "Atlassian jira信息泄露漏洞(CVE-2019-8449)",
		Type:          "脆弱性",
		Desc:          "Atlassian Jira是澳大利亚Atlassian公司的一套缺陷跟踪管理系统。该系统主要用于对工作中各类问题、缺陷进行跟踪管理。 Atlassian Jira 8.4.0之前版本中的/rest/api/latest/groupuserpicker资源存在信息泄露漏洞。该漏洞源于网络系统或产品在运行过程中存在配置等错误。未授权的攻击者可利用漏洞获取受影响组件敏感信息。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/rest/api/latest/groupuserpicker",
	},
	{
		Name:          "spark未授权命令执行漏洞",
		Type:          "脆弱性",
		Desc:          "spark是专为大规模数据处理而设计的快速通用的计算引擎，未授权用户可以向管理节点提交一个恶意代码构成的应用，从而获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/v1/submissions/create",
	},
	{
		Name:          "F5 BIG-IP命令注入漏洞（CVE-2021-22986）",
		Type:          "脆弱性",
		Desc:          "F5 BIG-IP存在代码执行漏洞，该漏洞允许攻击者通过BIG-IP管理界面和自身IP地址对接口进行网络访问以执行任意命令，从而获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/mgmt/tm/util/bash?reqType=22986",
	},
	{
		Name:          "spring data rest 远程命令执行（CVE-2017-8046）",
		Type:          "脆弱性",
		Desc:          "攻击者通过构造恶意的PATCH请求提交给spring-data-rest服务器，使用特制的JSON数据来运行任意的Java代码，从而实现远程代码执行攻击，获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用PATCH请求访问/api/auth/custom/mock/vulnerability/customers/1",
	},
	{
		Name:          "Apache APISIX 远程代码执行漏洞（CVE-2022-24112）",
		Type:          "脆弱性",
		Desc:          "攻击者可以向batch-requests插件发送请求来绕过管理API的IP限制，造成远程代码执行漏洞，从而获取服务器权限。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/apisix/batch-requests",
	},
	{
		Name:          "kylin远程命令执行漏洞（CVE-2020-1956）",
		Type:          "脆弱性",
		Desc:          "Apache Kylin 是美国 Apache 软件基金会的一款开源的分布式分析型数据仓库，Apache Kylin 中的静态 API 存在安全漏洞。攻击者可借助特制输入利用该漏洞在系统上执行任意系统命令，从而获取。",
		Level:         "高危",
		TriggerMethod: "使用PUT请求访问/api/auth/custom/mock/vulnerability/kylin/api/cubes/kylin_streaming_cube/learn_kylin/migrate",
	},
	{
		Name:          "k8s Dashboard面板认证绕过漏洞（CVE-2018-18264）",
		Type:          "脆弱性",
		Desc:          "使用Kubernetes Dashboard v1.10及以前的版本有跳过用户身份认证，及使用Dashboard登录账号读取集群密钥信息的风险。",
		Level:         "高危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/api/v1/secret/kube-system/kubernetes-dashboard-certs",
	},
	{
		Name:          "Cisco Data Center Network Manager 认证绕过漏洞（CVE-2019-15977）",
		Type:          "脆弱性",
		Desc:          "服务器上执行的Cisco Data Center Network Manager (DCNM) 版本受到一个验证绕过漏洞登录web管理页面，从而导致信息泄露。",
		Level:         "高危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/DbAdminWSService/DbAdminWS",
	},
	{
		Name:          "Zimbra 组件漏洞信息泄露",
		Type:          "脆弱性",
		Desc:          "Zimbra 组件存在一个信息泄露漏洞，攻击者可以通过soap接口获取低权限token。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/service/soap",
	},
	{
		Name:          "API接口文档未授权访问",
		Type:          "脆弱性",
		Desc:          "API接口文档中包含了应用中声明的路径操作、请求参数、请求体安全性等的声明，现有多种框架以及工具可以生成API文档，例如Swagger-ui、Fastapi、Redoc等，若接口文档访问路径存在相关的配置缺陷，会导致接口信息泄漏问题，黑客可通过泄露接口进行安全测试。",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/openapi.json",
	},
	// {
	// 	Name:          "CORS配置不当导致敏感信息泄露",
	// 	Type:          "脆弱性",
	// 	Desc:          "跨域资源共享（CORS）是一种使浏览器可以在受控的情况下执行跨域请求的网络机制。如果目标站点CORS配置不当。黑客通过构造恶意网站，当用户访问恶意网站时，可窃取目标站点用户敏感信息。",
	// 	Level:         "中危",
	// 	TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/get",
	// },
	{
		Name:          "HTTP Basic Authorization弱口令",
		Type:          "脆弱性",
		Desc:          "HTTP基本认证（Basic Authorization弱口令）是一种用来允许网页浏览器或其他客户端程序在请求时提供用户名和口令形式的身份凭证的一种登录验证方式，发现目标站点HTTP基本认证存在弱口令登录问题，弱口令可能导致系统失陷以及数据库信息被窃取等问题。",
		Level:         "中危",
		TriggerMethod: "/api/auth/custom/mock/vulnerability/user",
	},
	{
		Name:          "重要接口未授权访问",
		Type:          "脆弱性",
		Desc:          "WEB应用对高危或核心接口功能未做鉴权，关键核心资源不安全的被访问，会导致系统受到恶意攻击或敏感信息泄露。",
		Level:         "中危",
		TriggerMethod: "/api/auth/custom/mock/vulnerability/user/admin/infolog.log",
	},

	// 风险
	{
		Name:          "账号多地访问",
		Type:          "风险",
		Desc:          "同一个账号在多个不同的地理位置上登录,可能存在风险",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/OneAccountWithMultiPlace,然后在后端调用流量回放服务,回放5次不同源地址的同一账号访问包",
	},
	{
		Name:          "账号多IP访问",
		Type:          "风险",
		Desc:          "同一个账号在多个不同的IP上登录,可能存在风险",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/OneAccountWithMultiIP,然后在后端调用流量回放服务,回放5次不同源IP的同一账号访问包",
	},
	{
		Name:          "境内IP有多个账号身份",
		Type:          "风险",
		Desc:          "同一个应用的多个账号使用境内的同一个IP访问,可能存在风险",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/LocalIPWithMultiAccount,然后在后端调用流量回放服务,回放5次同一境内IP的不同账号访问包",
	},
	{
		Name:          "境外IP有多个账号身份",
		Type:          "风险",
		Desc:          "同一个应用的多个账号使用境外的同一个IP访问,可能存在风险",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/ForeignIPWithMultiAccount,然后在后端调用流量回放服务,回放5次同一境外IP的不同账号访问包",
	},
	{
		Name:          "单个账号一段时间内返回大量敏感数据",
		Type:          "风险",
		Desc:          "单个账号一段时间内返回大量敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnTooMuchSensitiveDataPeriod,后端返回5个敏感数据",
	},
	{
		Name:          "单个IP一段时间内返回大量敏感数据",
		Type:          "风险",
		Desc:          "单个IP一段时间内返回大量敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPReturnTooMuchSensitiveDataPeriod,后端返回5个敏感数据",
	},
	{
		Name:          "单个账号单次返回大量敏感数据",
		Type:          "风险",
		Desc:          "单个账号单次返回大量敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnTooMuchSensitiveDataOnce,后端返回5个敏感数据",
	},
	{
		Name:          "单个IP单次返回大量敏感数据",
		Type:          "风险",
		Desc:          "单个IP单次返回大量敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPReturnTooMuchSensitiveDataOnce,后端返回5个敏感数据",
	},
	{
		Name:          "单个账号单次返回敏感数据类型超过15种",
		Type:          "风险",
		Desc:          "单个账号单次返回敏感数据类型超过15种",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnTooManyKindsOfSensitiveDataOnce,后端返回14种敏感数据类型",
	},
	{
		Name:          "单个IP单次返回敏感数据类型超过15种",
		Type:          "风险",
		Desc:          "单个IP单次返回敏感数据类型超过15种",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPReturnTooManyKindsOfSensitiveDataOnce,后端返回14种敏感数据类型",
	},
	{
		Name:          "单个账号单次返回新类型的敏感数据",
		Type:          "风险",
		Desc:          "单个账号单次返回新类型的敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnNewTypeSensiDataOnce, 返回新型敏感数据类型",
	},
	{
		Name:          "单个IP单次返回新类型的敏感数据",
		Type:          "风险",
		Desc:          "单个IP单次返回新类型的敏感数据",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPReturnNewTypeSensiDataOnce, 返回新型敏感数据类型",
	},
	{
		Name:          "单个账号在一段时间内进行请求参数值遍历",
		Type:          "风险",
		Desc:          "单个账号在一段时间内频繁变换请求参数值,疑似进行请求参数值遍历",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountRequestTraversePeriod,变换请求参数100次",
	},
	{
		Name:          "单个IP在一段时间内进行请求参数值遍历",
		Type:          "风险",
		Desc:          "单个IP在一段时间内频繁变换请求参数值,疑似进行请求参数值遍历",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPRequestTraversePeriod,变换请求参数100次",
	},
	{
		Name:          "请求参数值出现新类型",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求参数进行学习,构建了请求参数画像,当请求参数值的类型与画像不符合时,则为异常,疑似攻击者在构造异常参数",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/NewTypeInRequest,请求参数包含新类型",
	},
	{
		Name:          "请求方法异常",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求方法进行学习,构建了请求方法画像,当请求方法与画像不符合时,则为异常,疑似攻击者在构造异常参数,例如将HTTP中的请求方法由GET更改为DELETE,或由POST更改为DELETE,或由PUT更改为POST",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/AbnormalRequestMethod,请求方式异常",
	},
	{
		Name:          "单个账号在一段时间内返回大量4XX",
		Type:          "风险",
		Desc:          "单个账号在一段时间内返回大量400、401、403、404、405、408状态码",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnTooMuch4XXPeriod 100次,每次都返回400",
	},
	{
		Name:          "单个IP在一段时间内返回大量4XX",
		Type:          "风险",
		Desc:          "单个IP在一段时间内返回大量400、401、403、404、405、408状态码",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountReturnTooMuch4XXPeriod 100次,每次都返回401",
	},
	{
		Name:          "单个账号在一段时间内频繁访问同一API",
		Type:          "风险",
		Desc:          "单个账号在一段时间内频繁访问同一API",
		Level:         "中危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountVisitSameAPIPeriod 这一API 100次",
	},
	{
		Name:          "单个账号在异常时间段频繁访问同一API",
		Type:          "风险",
		Desc:          "单个账号在异常时间段频繁访问同一API",
		Level:         "中危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountVisitSameAPIAbnormalPeriod 这一API 100次",
	},
	{
		Name:          "单个IP在一段时间内频繁访问同一API",
		Type:          "风险",
		Desc:          "单个IP在一段时间内频繁访问同一API",
		Level:         "中危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPVisitSameAPIPeriod 这一API 100次",
	},
	{
		Name:          "单个IP在异常时间段频繁访问同一API",
		Type:          "风险",
		Desc:          "单个IP在异常时间段频繁访问同一API",
		Level:         "中危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPVisitSameAPIAbnormalPeriod 这一API 100次",
	},
	{
		Name:          "请求参数名缺失",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求参数进行学习,构建了请求参数画像,当请求参数名与画像不符合时,则为异常,疑似攻击者在构造异常参数,例如构造缺失的参数名",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/LackOfVarName,请求参数名缺失",
	},
	{
		Name:          "请求参数出现非预期的参数名",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求参数进行学习,构建了请求参数画像,当请求参数名与画像不符合时,则为异常,疑似攻击者在构造异常参数,例如构造没有出现过的参数名",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/UnexpectedVarName,请求参数出现非预期的参数名",
	},
	{
		Name:          "单个账号在一段时间内进行路径遍历",
		Type:          "风险",
		Desc:          "单个账号在一段时间内频繁变换路径,疑似进行路径遍历",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleAccountPathTraversePeriod/:name,频繁变换参数名",
	},
	{
		Name:          "单个IP在一段时间内进行路径遍历",
		Type:          "风险",
		Desc:          "单个IP在一段时间内频繁变换路径,疑似进行路径遍历",
		Level:         "高危",
		TriggerMethod: "访问/api/auth/custom/mock/risk/SingleIPPathTraversePeriod/:name,频繁变换参数名",
	},
}

var SecurityEventsList = []model.SecurityEvents{
	{
		Name:          "参数遍历获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "攻击者通过API请求参数遍历方式(比如ID类资源变换请求),获得了大量敏感数据",
		Level:         "中危",
		TriggerMethod: "同源IP/账号和同API,A和B中不同事件发生在相同时间段内,允许时间窗口包含关系,比如A发生在3:00 ~ 3:15,B发生在3:00 ~ 4:00,不限顺序",
	},
	{
		Name:          "频繁访问获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "攻击者通过频繁访问涉敏API接口,获得了大量敏感数据",
		Level:         "中危",
		TriggerMethod: "同源IP/账号和同API,A和B中不同事件发生在相同时间段内,允许时间窗口包含关系,比如A发生在3:00 ~ 3:15,B发生在3:00 ~ 4:00,不限顺序",
	},
	{
		Name:          "异常时间段频繁访问获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "攻击者在异常时间段频繁访问API,获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同源IP/账号和同API,A和B中不同事件发生在相同时间段内,允许时间窗口包含关系,比如A发生在3:00 ~ 3:15,B发生在3:00 ~ 4:00,不限顺序",
	},
	{
		Name:          "发生探测攻击并通过参数遍历获取过量敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段操作行为,首先进行扫描/探测类操作,寻找到可利用点之后通过参数遍历方式获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "根据相同源IP/账号关联,A和B和C需要在同一个时间窗口内发生;B和C发生在同一个API,且该api在A应用内",
	},
	{
		Name:          "发生探测攻击并通过频繁访问获取过量敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段攻击行为,首先进行扫描/探测类操作,寻找到可利用点之后通过频繁访问API的方式获取大量敏感数据",
		Level:         "中危",
		TriggerMethod: "根据相同源IP/账号关联,A和B和C需要在同一个时间窗口内发生;B和C发生在同一个API,且该api在A应用内",
	},
	{
		Name:          "发生探测攻击并在异常时间段频繁访问获取非预期敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段攻击行为,首先进行扫描/探测类操作,寻找到可利用点之后在异常时间段内频繁访问获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "根据相同源IP/账号关联,A和B和C需要在同一个时间窗口内发生;B和C发生在同一个API,且该api在A应用内",
	},
	{
		Name:          "通过恶意构造请求窃取额外敏感数据",
		Type:          "安全事件",
		Desc:          "疑似攻击者通过篡改请求内容,成功发现可利用点并获取了额外的敏感数据",
		Level:         "中危",
		TriggerMethod: "A和B发生在同一时间窗口内,且是同一api",
	},
	{
		Name:          "API接口遭遇渗透攻击",
		Type:          "安全事件",
		Desc:          "疑似攻击者对API接口实施渗透攻击,需确认实施者是授权正常操作还是恶意攻击",
		Level:         "高危",
		TriggerMethod: "A中出现事件和B中出现的事件需要同源IP/账号,并且A和B出现在同一时间窗,A和B发生在同一个应用",
	},
	{
		Name:          "通过恶意构造请求获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "疑似攻击者通过篡改请求内容,结合批量访问行为,在一段时间内成功获取了大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同源IP/账号和同API,A和B和C中不同事件发生在相同时间段内,允许时间窗口包含关系,比如A发生在3:00 ~ 3:15,B发生在3:00 ~ 4:00;",
	},
	{
		Name:          "账号失陷并下载了大量敏感数据",
		Type:          "安全事件",
		Desc:          "疑似用户账号被盗用后登录获取了大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同账号,A、B发生在同一时间窗口内;A和B发生在同一个应用",
	},
}

// 初始化业务句柄
func InitDB() *gorm.DB {
	var compileMode, host string

	flag.StringVar(&compileMode, "m", "test", "运行模式")
	flag.Parse()

	// 区分生产环境和测试环境
	if compileMode == "production" {
		host = viper.GetString("datasource.productionhost")
	} else {
		host = viper.GetString("datasource.testhost")
	}
	// fmt.Println("host is", host)
	// 一系列的读取配置操作

	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")

	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		username,
		password,
		database,
		port,
	)
	// 然后连接这个数据库
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect to postgres, error" + err.Error())
	}

	InitAllTables(db)

	DB = db
	return db
}

// 初始化泄密溯源的句柄
func InitDataLeakgeDB() *gorm.DB {
	var compileMode, host string

	// 区分生产环境和测试环境
	if compileMode == "production" {
		host = viper.GetString("datasource.productionhost")
	} else {
		host = viper.GetString("datasource.testhost")
	}
	// fmt.Println("host is", host)
	// 一系列的读取配置操作

	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.dataleakge_database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")

	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		username,
		password,
		database,
		port,
	)
	// 然后连接这个数据库
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect to postgres, error" + err.Error())
	}

	DataLeakgeDB = db
	return DataLeakgeDB
}

// 初始化所有表
func InitAllTables(db *gorm.DB) {
	InitUserTable(db)
	InitRightsTable(db)
	InitToolsTable(db)
	InitToolsConfigTable(db)
	InitTaskTable(db)
	InitCronTaskTable(db)
	InitCronTaskResultTable(db)
	InitRiskAndVulnerabilityTable(db)
	InitPOCConfigTable(db)
	InitRiskAndVulnerabilityLogTable(db)
	InitAgentInstallTable(db)
	InitSecurityEventsTable(db)
}

// 初始化工具基础信息表
func InitToolsTable(db *gorm.DB) {
	db.AutoMigrate(&model.Tool{})
}

// 初始化工具配置信息表
func InitToolsConfigTable(db *gorm.DB) {
	db.AutoMigrate(&model.ToolConfig{})
}

// 初始化用户表
func InitUserTable(db *gorm.DB) {
	UserList := []model.User{

		// 默认的超级管理员
		{
			UserName: "admin",
			Mobile:   "18578660000",
			Type:     1,
			Email:    "yqy1160058763@qq.com",
			MgState:  true, RoleName: "超级管理员",
			WorkNum: "10000颜桥宇", PassWord: "admin",
		},
		// 默认的访客
		{
			UserName: "guest",
			Mobile:   "18578660000",
			Type:     1,
			Email:    "yqy1160058763@qq.com",
			MgState:  true,
			RoleName: "访客",
			PassWord: "guest",
		},
	}

	db.Exec("truncate table users;")

	db.AutoMigrate(&model.User{})
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&UserList)
}

// 初始化权限表
func InitRightsTable(db *gorm.DB) {
	RightsList := []model.Rights{
		{AuthName: "首页", Level: 0, Pid: 0, Path: "home"},
		{AuthName: "任务", Level: 1, Pid: 0, Path: "dashboard"},
		{AuthName: "工具盒", Level: 1, Pid: 0, Path: "toolbox"},
		{AuthName: "全局配置", Level: 2, Pid: 0, Path: "config"},
		{AuthName: "用户管理", Level: 2, Pid: 4, Path: "users"},
		{AuthName: "权限管理", Level: 2, Pid: 4, Path: "rights"},
	}

	db.AutoMigrate(&model.Rights{})

	// 这里要先truncate,是我的问题
	// 本来可以用upsert来解决更新已有条目的问题
	// 但是我建表的时候没有选择主键,导致了现在更新出了点问题
	// 先用truncate暂时解决吧
	db.Exec("truncate table rights;")

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&RightsList)
}

// 初始化任务列表
func InitTaskTable(db *gorm.DB) {
	db.AutoMigrate(&model.Tasks{})
}

// 初始化定时任务列表
func InitCronTaskTable(db *gorm.DB) {
	db.AutoMigrate(&model.CronTasks{})
}

// 初始化定时任务结果列表
func InitCronTaskResultTable(db *gorm.DB) {
	db.AutoMigrate(&model.CronTasksResult{})
}

// 定制:初始化POC工具配置
func InitPOCConfigTable(db *gorm.DB) {
	db.AutoMigrate(&model.POCConfig{})
}

// 定制:初始化安装探针的配置的表
func InitAgentInstallTable(db *gorm.DB) {
	db.AutoMigrate(&model.AgentInstallConfig{})
}

// 定制:初始化脆弱性与风险的表
func InitRiskAndVulnerabilityTable(db *gorm.DB) {
	db.AutoMigrate(&model.RiskAndVulnerability{})
	// 这里要先truncate,是我的问题
	// 本来可以用upsert来解决更新已有条目的问题
	// 但是我建表的时候没有选择主键,导致了现在更新出了点问题
	// 先用truncate暂时解决吧
	db.Exec("truncate table risk_and_vulnerabilities;")

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&RiskAndVulnerabilityList)
}

// 定制:初始化脆弱性与风险日志
func InitRiskAndVulnerabilityLogTable(db *gorm.DB) {
	db.AutoMigrate(&model.RiskAndVulnerabilityLog{})
}

// 定制:初始化安全事件
func InitSecurityEventsTable(db *gorm.DB) {
	db.AutoMigrate(&model.SecurityEvents{})
	// 这里要先truncate,是我的问题
	// 本来可以用upsert来解决更新已有条目的问题
	// 但是我建表的时候没有选择主键,导致了现在更新出了点问题
	// 先用truncate暂时解决吧
	db.Exec("truncate table security_events;")

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&SecurityEventsList)
}

func GetDB() *gorm.DB {
	return DB
}

func GetDataLeakgeDB() *gorm.DB {
	return DataLeakgeDB
}
