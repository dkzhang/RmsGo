package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/DbManage"
	"github.com/dkzhang/RmsGo/webapi"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("parameter input error. Expected at least 1 patameter.")
		return
	}

	switch os.Args[1] {
	case "create_all":
		DbManage.CreateAllTable()
	case "seed_all":
		DbManage.SeedAllTable()
	case "import_from_file":
		DbManage.ImportFromFile("", "")
	case "run":
		webapi.Serve()
	}
}
