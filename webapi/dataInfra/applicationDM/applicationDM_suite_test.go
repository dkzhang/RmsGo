package applicationDM_test

import (
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApplicationDM(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApplicationDM Suite")
}

var adb applicationDB.ApplicationDB
var adm applicationDM.ApplicationDM

var _ = BeforeSuite(func() {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := pgOpsSqlx.ConnectToDatabase()
	pgOpsSqlx.CreateAllTable(db)

	theLogMap := logMap.NewLogMap(`C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Parameter\logmap.yaml`)

	adb = applicationDB.NewApplicationPg(db, application.TableApp, application.TableAppOps)
	var err error
	adm, err = applicationDM.NewMemoryMap(adb, theLogMap)
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	adb.Close()
})
