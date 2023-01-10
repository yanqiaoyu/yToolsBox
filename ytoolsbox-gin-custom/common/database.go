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
		Desc:          "密码设置不符合信息安全规范（比如为简单的数字字符组合）,例如黑客可以通过口令爆破等方式猜解用户账号密码,导致用户账号被盗取",
		Level:         "高危",
		TriggerMethod: "使用POST方法访问/api/auth/custom/mock/vulnerability/WeakPasswd,由于请求body中含有username=admin,password=admin,并且配置了相应的账号提取规则,因此可以触发<登录弱密码>这个脆弱性",
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
		Desc:          "权限认证方式不合理,采用GET方法传输鉴权信息,例如黑客可以通过中间人攻击等方式捕获用户的鉴权信息（token, sessionid等）,导致账号被盗取",
		Level:         "中危",
		TriggerMethod: "使用GET请求访问/api/auth/custom/mock/vulnerability/AuthInURL?sessionid=456,由于该URL携带了sessionid=456,因此可以触发<鉴权信息在url中>这个脆弱性",
	},
	{
		Name:          "敏感信息在url中",
		Type:          "脆弱性",
		Desc:          "接口传输不规范,传输过程中未对敏感数据进行加密,例如黑客可以通过中间人攻击等方式捕获用户的敏感信息（身份证、电话号码等）,导致用户信息泄露",
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
		Desc:          "攻击者未经授权获取Hadoop服务敏感信息和做敏感操作，且该页面存在系统命令执行漏洞风险。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
	},
	{
		Name:          "SonarQube未授权访问",
		Type:          "脆弱性",
		Desc:          "SonarQube系统在默认配置下，会将通过审计的源代码上传至SonarQube平台。由于SonarQube缺少对API接口访问的鉴权控制，攻击者利用该漏洞，可在未授权的情况下通过访问上述API接口，获取SonarQube平台上的程序源代码，构成项目源代码数据泄露风险。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
	},
	{
		Name:          "ssh secret key信息泄露",
		Type:          "脆弱性",
		Desc:          "ssh私钥默认保存在ssh目录下的id_rsa文件中，如果该私钥泄露，黑客可以利用其攻陷主机。如果攻击者利用成功，将会有完整SSH访问权限，能够运行任何命令。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
	},
	{
		Name:          "JDBC连接字符串信息泄露",
		Type:          "脆弱性",
		Desc:          "JDBC连接字符串含有数据库连接地址、数据库账号、密码信息。如果JDBC连接字符串泄露，黑客可利用其获得数据库权限。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
	},
	{
		Name:          "Alibaba Nacos未授权访问",
		Type:          "脆弱性",
		Desc:          "Nacos 官方github在2020年12月29日发布的issue中披露Alibaba Nacos 存在一个由于不当处理User-Agent导致的未授权访问漏洞。通过该漏洞，攻击者可以进行任意操作，包括创建新用户并进行登录后操作。",
		Level:         "中危",
		TriggerMethod: "使用POST请求访问/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application",
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
		TriggerMethod: "To Be Done",
	},
	{
		Name:          "单个IP单次返回新类型的敏感数据",
		Type:          "风险",
		Desc:          "单个IP单次返回新类型的敏感数据",
		Level:         "高危",
		TriggerMethod: "To Be Done",
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
		TriggerMethod: "To Be Done",
	},
	{
		Name:          "请求方法异常",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求方法进行学习,构建了请求方法画像,当请求方法与画像不符合时,则为异常,疑似攻击者在构造异常参数,例如将HTTP中的请求方法由GET更改为DELECT,或由POST更改为DELECT,或由PUT更改为POST",
		Level:         "高危",
		TriggerMethod: "To Be Done",
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
		TriggerMethod: "To Be Done",
	},
	{
		Name:          "请求参数出现非预期的参数名",
		Type:          "风险",
		Desc:          "基于过去一段时间内对每个API的请求参数进行学习,构建了请求参数画像,当请求参数名与画像不符合时,则为异常,疑似攻击者在构造异常参数,例如构造没有出现过的参数名",
		Level:         "高危",
		TriggerMethod: "To Be Done",
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
		Desc:          "攻击者通过API请求参数遍历方式（比如ID类资源变换请求），获得了大量敏感数据",
		Level:         "中危",
		TriggerMethod: "同源IP/账号和同API，A和B中不同事件发生在相同时间段内，允许时间窗口包含关系，比如A发生在3:00 ~ 3:15，B发生在3:00 ~ 4:00，不限顺序",
	},
	{
		Name:          "频繁访问获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "攻击者通过频繁访问涉敏API接口，获得了大量敏感数据",
		Level:         "中危",
		TriggerMethod: "同源IP/账号和同API，A和B中不同事件发生在相同时间段内，允许时间窗口包含关系，比如A发生在3:00 ~ 3:15，B发生在3:00 ~ 4:00，不限顺序",
	},
	{
		Name:          "异常时间段频繁访问获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "攻击者在异常时间段频繁访问API，获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同源IP/账号和同API，A和B中不同事件发生在相同时间段内，允许时间窗口包含关系，比如A发生在3:00 ~ 3:15，B发生在3:00 ~ 4:00，不限顺序",
	},
	{
		Name:          "发生探测攻击并通过参数遍历获取过量敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段操作行为，首先进行扫描/探测类操作，寻找到可利用点之后通过参数遍历方式获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "根据相同源IP/账号关联，A和B和C需要在同一个时间窗口内发生；B和C发生在同一个API，且该api在A应用内",
	},
	{
		Name:          "发生探测攻击并通过频繁访问获取过量敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段攻击行为，首先进行扫描/探测类操作，寻找到可利用点之后通过频繁访问API的方式获取大量敏感数据",
		Level:         "中危",
		TriggerMethod: "根据相同源IP/账号关联，A和B和C需要在同一个时间窗口内发生；B和C发生在同一个API，且该api在A应用内",
	},
	{
		Name:          "发生探测攻击并在异常时间段频繁访问获取非预期敏感数据",
		Type:          "安全事件",
		Desc:          "识别到高可疑攻击者多阶段攻击行为，首先进行扫描/探测类操作，寻找到可利用点之后在异常时间段内频繁访问获取大量敏感数据",
		Level:         "高危",
		TriggerMethod: "根据相同源IP/账号关联，A和B和C需要在同一个时间窗口内发生；B和C发生在同一个API，且该api在A应用内",
	},
	{
		Name:          "通过恶意构造请求窃取额外敏感数据",
		Type:          "安全事件",
		Desc:          "疑似攻击者通过篡改请求内容，成功发现可利用点并获取了额外的敏感数据",
		Level:         "中危",
		TriggerMethod: "A和B发生在同一时间窗口内，且是同一api",
	},
	{
		Name:          "API接口遭遇渗透攻击",
		Type:          "安全事件",
		Desc:          "疑似攻击者对API接口实施渗透攻击，需确认实施者是授权正常操作还是恶意攻击",
		Level:         "高危",
		TriggerMethod: "A中出现事件和B中出现的事件需要同源IP/账号，并且A和B出现在同一时间窗，A和B发生在同一个应用",
	},
	{
		Name:          "通过恶意构造请求获取大量敏感数据",
		Type:          "安全事件",
		Desc:          "疑似攻击者通过篡改请求内容，结合批量访问行为，在一段时间内成功获取了大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同源IP/账号和同API，A和B和C中不同事件发生在相同时间段内，允许时间窗口包含关系，比如A发生在3:00 ~ 3:15，B发生在3:00 ~ 4:00；",
	},
	{
		Name:          "账号失陷并下载了大量敏感数据",
		Type:          "安全事件",
		Desc:          "疑似用户账号被盗用后登录获取了大量敏感数据",
		Level:         "高危",
		TriggerMethod: "同账号，A、B发生在同一时间窗口内；A和B发生在同一个应用",
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
