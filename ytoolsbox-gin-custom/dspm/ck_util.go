package dspm

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
	"main/cons"
)

var CkDb *sqlx.DB

func GetCkConnection() {

	dspmPwdConfig := cons.DspmPwdConfig
	// 连接字符串
	var err error
	ckConnStr := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s",
		cons.DspmAddr, cons.CkPort, dspmPwdConfig.ClickhouseUsername, dspmPwdConfig.ClickhousePassword, "dspm")
	log.Printf("ckSql连接字符串：%v\r\n", ckConnStr)
	CkDb, err = sqlx.Open("clickhouse", ckConnStr)
	if err != nil {
		log.Println("成功连接到ck数据库失败:", err)
	}
	// 测试连接
	err = CkDb.Ping()
	if err != nil {
		log.Println("成功连接到ck数据库失败:", err)
	}
	log.Println("成功连接到ck数据库！")
}

func ClearCkModelData() {
	sql := fmt.Sprintf(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'dspm'
	AND (table_name LIKE '%s%%' or table_name LIKE '%s%%')`, "risk_baseline_", "risk_statistics_by_")

	if CkDb == nil {
		GetCkConnection()
	}
	// 执行查询
	rows, err := CkDb.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tableNames []string
	// 遍历结果
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}
	tableNames = append(tableNames, "risk_http")
	TruncateCkTable(tableNames)
}

func TruncateCkTable(tableNames []string) {

	if CkDb == nil {
		GetCkConnection()
	}
	for _, tableName := range tableNames {
		result, err := CkDb.Exec(fmt.Sprintf("truncate \"dspm\".\"%s\"", tableName))
		if err != nil {
			log.Printf("truncate %s status:%s,err:%v", tableName, result, err)
		}
		affected, err := result.RowsAffected()
		if err != nil {
			log.Printf("truncate %s status:%s,err:%v", tableName, result, err)
		}
		log.Printf("truncate %s status:%d", tableName, affected)
	}

}
