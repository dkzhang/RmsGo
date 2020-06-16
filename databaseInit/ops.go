package databaseInit

import (
	"fmt"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func CreateAllTable() {
	fmt.Printf("删除所有表格并重建 \n")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
}

func SeedAllTable() {
	fmt.Printf("用测试数据初始化所有数据库表")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
	pgOps.SeedAllTable(db)
}

func ImportFromFile(tableName, fileName string) {
	fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
	pgOps.ImportFromFile(db)
}

func connectToDatabase() (db *sqlx.DB) {
	theDbSecurity, err := databaseSecurity.LoadDbSecurity(os.Getenv("DbSE"))
	if err != nil {
		log.Fatalf("dbConfig.LoadDbSecurity error, ENV DbSE = %s, error = %v", os.Getenv("DbSE"), err)
		return
	}

	db, err = postgreOps.ConnectToDatabase(theDbSecurity.ThePgSecurity)
	if err != nil {
		log.Fatalf("postgreSQL.ConnectToDatabase error,error = %v", err)
		return
	} else {
		log.Printf("postgreSQL.ConnectToDatabase success, db = %v.", db)
	}
	return db
}
