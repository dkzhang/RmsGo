package pgOpsSqlx

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOpsSqlx"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func CmdCreateAllTable() {
	fmt.Printf("删除所有表格并重建 \n")
	db := ConnectToDatabase()
	CreateAllTable(db)
}

func CmdSeedAllTable() {
	fmt.Printf("用测试数据初始化所有数据库表")
	db := ConnectToDatabase()
	CreateAllTable(db)
	SeedAllTable(db)
}

func CmdImportFromFile(tableName, fileName string) {
	fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	db := ConnectToDatabase()
	CreateAllTable(db)
	ImportFromFile(db)
}

/////////////////////////////////////////////////////////////////////////////////////////

func ConnectToDatabase() (db *sqlx.DB) {
	theDbSecurity, err := databaseSecurity.LoadDbSecurity(os.Getenv("DbSE"))
	if err != nil {
		log.Fatalf("dbConfig.LoadDbSecurity error, ENV DbSE = %s, error = %v", os.Getenv("DbSE"), err)
		return
	}

	db, err = postgreOpsSqlx.ConnectToDatabase(theDbSecurity.ThePgSecurity)
	if err != nil {
		log.Fatalf("postgreSQL.ConnectToDatabase error,error = %v", err)
		return
	} else {
		log.Printf("postgreSQL.ConnectToDatabase success, db = %v.", db)
	}
	return db
}

func CloseDatabase(db *sqlx.DB) error {
	return db.Close()
}
