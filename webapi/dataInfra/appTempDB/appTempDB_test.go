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
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	atdm := NewAppTempPg(db)

	Convey("TestAppTempPg", t, func() {
		Convey("Insert one appTemp", func() {
			id, err := atdm.InsertAppTemp(appTemp.AppTemp{
				UserID:       1,
				AppType:      1,
				BasicContent: "BasicContent1",
				ExtraContent: "ExtraContent1",
			})
			So(err, ShouldBeNil)
			Printf("Insert one appTemp success, id = %d \n", id)
		})
		Convey("Insert another appTemp", func() {
			id, err := atdm.InsertAppTemp(appTemp.AppTemp{
				UserID:       1,
				AppType:      1,
				BasicContent: "BasicContent1",
				ExtraContent: "ExtraContent1",
			})
			So(err, ShouldBeNil)
			Printf("Insert another appTemp success, id = %d \n", id)
		})

		Convey("QueryAppTempByID", func() {
			apptemp1, err := atdm.QueryAppTempByID(1)
			So(err, ShouldBeNil)
			Printf("QueryAppTempByID 1 success, apptemp = %v \n", apptemp1)

			apptemp2, err := atdm.QueryAppTempByID(2)
			So(err, ShouldBeNil)
			Printf("QueryAppTempByID 2 success, apptemp = %v \n", apptemp2)
		})

		Convey("QueryAppTempByOwner", func() {
			apptemps, err := atdm.QueryAppTempByOwner(1)
			So(err, ShouldBeNil)
			Printf("QueryAppTempByID 1 success, apptemp = %v \n", apptemps)
		})

		Convey("UpdateAppTemp", func() {
			err := atdm.UpdateAppTemp(appTemp.AppTemp{
				ApplicationID: 1,
				UserID:        5,
				AppType:       5,
				BasicContent:  "BasicContent5",
				ExtraContent:  "ExtraContent5",
			})
			So(err, ShouldBeNil)
			Printf("UpdateAppTemp 1 success \n")

			apptemp1, err := atdm.QueryAppTempByID(1)
			So(err, ShouldBeNil)
			Printf("QueryAppTempByID 1 success, apptemp = %v \n", apptemp1)
		})

		Convey("DeleteAppTemp", func() {
			err := atdm.DeleteAppTemp(2)
			So(err, ShouldBeNil)
			Printf("UpdateAppTemp 1 success \n")

			_, err = atdm.QueryAppTempByID(2)
			So(err, ShouldBeError)
			Printf("QueryAppTempByID 2 with error since it has been deleted, error = %v \n", err)
		})
	})
}
