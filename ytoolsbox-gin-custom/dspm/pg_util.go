package dspm

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"main/cons"
)

var Pgdb *sqlx.DB

func GetPgConnection() {
	dspmPwdConfig := cons.DspmPwdConfig
	// 连接字符串
	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cons.DspmAddr, cons.PgPort, dspmPwdConfig.PostgresqlUsername, dspmPwdConfig.PostgresqlPassword, "dsc")
	log.Printf("sql连接字符串：%v\r\n", pgConnStr)
	// 打开数据库连接
	var err error
	Pgdb, err = sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Println("连接pg数据库失败:", err)
	}
	// 测试连接
	err = Pgdb.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("成功连接到pg数据库！")
}

func ClearPgModelData(mid string) {
	if Pgdb == nil {
		GetPgConnection()
	}
	// 检查遍历过程中是否有错误
	deleteSql := "delete from \"risk\".\"tb_model_task\" "
	if mid != "" {
		deleteSql = fmt.Sprintf("%s where \"modelId\" = '%s'", deleteSql, mid)
	}
	result, err := Pgdb.Exec(deleteSql)
	if err != nil {
		log.Printf("deleteSql:%s; status:%s,err:%v", deleteSql, result, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("deleteSql:%s; status:%s,err:%v", deleteSql, result, err)
	}
	log.Printf("deleteSql:%s;status:%d", deleteSql, affected)
}
