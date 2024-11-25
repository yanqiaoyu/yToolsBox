package cons

/*
*
保存通用环境变量
*/
type DspmConfig struct {
	ClickhousePassword            string `json:"clickhouse_password"`
	ClickhouseUsername            string `json:"clickhouse_username"`
	ElasticsearchPassword         string `json:"elasticsearch_password"`
	ElasticsearchUsername         string `json:"elasticsearch_username"`
	EtcdPassword                  string `json:"etcd_password"`
	EtcdUsername                  string `json:"etcd_username"`
	KafkaBrokerPassword           string `json:"kafka_broker_password"`
	KafkaBrokerUsername           string `json:"kafka_broker_username"`
	KafkaClientPassword           string `json:"kafka_client_password"`
	KafkaClientUsername           string `json:"kafka_client_username"`
	KafkaZookeeperClientPassword  string `json:"kafka_zookeeper_client_password"`
	KafkaZookeeperClientUsername  string `json:"kafka_zookeeper_client_username"`
	MariadbPassword               string `json:"mariadb_password"`
	MariadbUsername               string `json:"mariadb_username"`
	MinioPassword                 string `json:"minio_password"`
	MinioUsername                 string `json:"minio_username"`
	MongodbPassword               string `json:"mongodb_password"`
	MongodbUsername               string `json:"mongodb_username"`
	PostgresqlPassword            string `json:"postgresql_password"`
	PostgresqlReplicationPassword string `json:"postgresql_replication_password"`
	PostgresqlReplicationUsername string `json:"postgresql_replication_username"`
	PostgresqlUsername            string `json:"postgresql_username"`
	PulsarPassword                string `json:"pulsar_password"`
	PulsarUsername                string `json:"pulsar_username"`
	RedisPassword                 string `json:"redis_password"`
	RedisUsername                 string `json:"redis_username"`
}

var DspmPwdConfig DspmConfig

var CkPort int = 30900
var PgPort int = 30432
var PulsarPort int = 30650
var FlinkPort int = 30081

var DspmAddr string = ""

var PulsarToken = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhZG1pbiJ9.XdHUncGOu4hQSIrIK3SxOiRVGS9mMVByukgGXGRUlaU"

var HttpSourceTopic = "persistent://10001001/default/http_log"

var PocAddr = ""

var FilePath = "/data/jsonFile"

var AccountSlice = []string{"xiaoqiao", "zhaosi", "xiaoli", "xiaoming", "zhangsan",
	"wukong", "zhuge", "xiaoqiang", "wanger", "lifei"}

var SensitiveStr = `{
	"username":[
		"王杰",
		"王梅",
		"牛伟",
		"萧想",
		"李鹏",
		"邹秀荣",
		"韩超",
		"宫俊",
		"何鹏",
		"黄丽丽"
	],
	"idCard":[
		"210283195711274627",
		"130827195009114940",
		"13102819641006540X",
		"140827198512022977",
		"330602197501144840",
		"441226196607162431",
		"130723194812145783",
		"140829199204081921",
		"513430193710240938",
		"420503198710218401"
	],
	"phone":[
		"13761417814",
		"13118329471",
		"18822784508",
		"15332917845",
		"13746639261",
		"14721698285",
		"18970522737",
		"14786733843",
		"13519652233",
		"15779378117"
	],
	"ip":[
		"189.9.182.25",
		"128.238.145.86",
		"212.147.197.100",
		"169.153.199.49",
		"157.193.187.101",
		"11.80.234.57",
		"188.165.185.204",
		"79.86.133.168",
		"9.147.208.123",
		"128.105.223.223"
	],
	"mac":[
		"00:0a:95:9d:68:16",
		"08:00:27:00:af:8c",
		"1c:1b:0d:5f:3e:9a",
		"40:8d:5c:2e:7f:bb",
		"78:4f:43:9a:21:ec",
		"78:4f:43:9a:21:e1",
		"78:4f:43:9a:21:e2",
		"78:4f:43:9a:21:e3",
		"78:4f:43:9a:21:e4",
		"78:4f:43:9a:21:e5"
	],
	"email":[
		"vgao@example.net",
		"oyan@example.net",
		"jing89@example.org",
		"yinfang@example.com",
		"rluo@example.com",
		"qdai@example.com",
		"qiangcheng@example.net",
		"gaoxia@example.net",
		"fliao@example.com",
		"yuanlei@example.org"
	],
	"phone_card":[
		"010-12345678",
		"021-23456789",
		"020-34567890",
		"0755-45678901",
		"028-56789012",
		"028-56789013",
		"028-56789014",
		"028-56789015"
	],
	"companyName":[
		"ABC公司",
		"XYZ有限公司",
		"123科技有限责任公司",
		"橙子网络科技股份有限公司",
		"绿叶食品集团有限公司",
		"绿叶食品集团有限公司1",
		"绿叶食品集团有限公司2",
		"绿叶食品集团有限公司3"
	]
}`
