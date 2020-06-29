package generalFormDraftDB

import (
	"github.com/dkzhang/RmsGo/databaseInit"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGeneralFormDraftPg_InsertGeneralFormDraft(t *testing.T) {
	os.Setenv("DbSE", `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`)
	db := databaseInit.ConnectToDatabase()
	pgOps.CreateAllTable(db)

	atdm := NewGeneralFormDraftPg(db)

	Convey("TestGeneralFormDraftPg", t, func() {
		Convey("Insert one GeneralFormDraft", func() {
			id, err := atdm.InsertGeneralFormDraft(generalFormDraft.GeneralFormDraft{
				UserID:       1,
				FormType:     1,
				BasicContent: "BasicContent1",
				ExtraContent: "ExtraContent1",
			})
			So(err, ShouldBeNil)
			Printf("Insert one GeneralFormDraft success, id = %d \n", id)
		})
		Convey("Insert another GeneralFormDraft", func() {
			id, err := atdm.InsertGeneralFormDraft(generalFormDraft.GeneralFormDraft{
				UserID:       1,
				FormType:     1,
				BasicContent: "BasicContent1",
				ExtraContent: "ExtraContent1",
			})
			So(err, ShouldBeNil)
			Printf("Insert another GeneralFormDraft success, id = %d \n", id)
		})

		Convey("QueryGeneralFormDraftByID", func() {
			GeneralFormDraft1, err := atdm.QueryGeneralFormDraftByID(1)
			So(err, ShouldBeNil)
			Printf("QueryGeneralFormDraftByID 1 success, GeneralFormDraft = %v \n", GeneralFormDraft1)

			GeneralFormDraft2, err := atdm.QueryGeneralFormDraftByID(2)
			So(err, ShouldBeNil)
			Printf("QueryGeneralFormDraftByID 2 success, GeneralFormDraft = %v \n", GeneralFormDraft2)
		})

		Convey("QueryGeneralFormDraftByOwner", func() {
			GeneralFormDrafts, err := atdm.QueryGeneralFormDraftByOwner(1)
			So(err, ShouldBeNil)
			Printf("QueryGeneralFormDraftByID 1 success, GeneralFormDraft = %v \n", GeneralFormDrafts)
		})

		Convey("UpdateGeneralFormDraft", func() {
			err := atdm.UpdateGeneralFormDraft(generalFormDraft.GeneralFormDraft{
				FormID:       1,
				UserID:       5,
				FormType:     5,
				BasicContent: "BasicContent5",
				ExtraContent: "ExtraContent5",
			})
			So(err, ShouldBeNil)
			Printf("UpdateGeneralFormDraft 1 success \n")

			GeneralFormDraft1, err := atdm.QueryGeneralFormDraftByID(1)
			So(err, ShouldBeNil)
			Printf("QueryGeneralFormDraftByID 1 success, GeneralFormDraft = %v \n", GeneralFormDraft1)
		})

		Convey("DeleteGeneralFormDraft", func() {
			err := atdm.DeleteGeneralFormDraft(2)
			So(err, ShouldBeNil)
			Printf("UpdateGeneralFormDraft 1 success \n")

			_, err = atdm.QueryGeneralFormDraftByID(2)
			So(err, ShouldBeError)
			Printf("QueryGeneralFormDraftByID 2 with error since it has been deleted, error = %v \n", err)
		})
	})
}
