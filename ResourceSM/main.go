package ResourceSM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/ResourceSM/gRpcService/server"

	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("parameter input error. Expected at least 1 patameter.")
		return
	}

	switch os.Args[1] {
	case "create_all":
		pgOpsSqlx.CmdCreateAllTable()
	case "seed_all":
		pgOpsSqlx.CmdSeedAllTable()
	case "import_from_file":
		pgOpsSqlx.CmdImportFromFile("", "")
	case "run":
		server.Serve()
	}
}
