package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/dbManage/pgManage"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("parameter input error. Expected at least 1 patameter.")
		return
	}

	switch os.Args[1] {
	case "create_all":
		fmt.Printf("删除所有表格并重建 \n")
		os.Setenv("DbConf", "./../Configuration/Security/database.yaml")
		PgManage.CreateAllTable()
	case "seed_all":
		fmt.Printf("无参：用测试数据初始化所有数据库表")
	case "import_from_file":
		fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	}
}
