package appTempDB

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/webapi/model/appTemp"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestAppTempPg_InsertAppTemp(t *testing.T) {
	Convey("Connect to database ", t, func() {
		os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
		db := databaseInit.ConnectToDatabase()
		pgOps.CreateAllTable(db)

		atdm := NewAppTempPg(db)

		Convey("Insert one appTemp", func() {
			id, err := atdm.InsertAppTemp(appTemp.AppTemp{
				UserID:       1,
				AppType:      1,
				BasicContent: "BasicContent1",
				ExtraContent: "ExtraContent1",
			})
			So(err, ShouldBeNil)
			Printf("Insert one appTemp success, id = %d", id)
		})
	})
}
