package main

import "fmt"

func main() {
	ops := ""

	switch ops {
	case "create_all":
		fmt.Printf("无参：删除所有表格并重建")
	case "seed_all":
		fmt.Printf("无参：用测试数据初始化所有数据库表")
	case "import_from_file":
		fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	}
}
