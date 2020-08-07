package projectDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/jmoiron/sqlx"
)

type ProjectPg struct {
	DBInfo
}

func NewProjectPg(sqlxdb *sqlx.DB, stn string, dtn string) ProjectPg {
	return ProjectPg{
		DBInfo: DBInfo{
			TheDB:            sqlxdb,
			StaticTableName:  stn,
			DynamicTableName: dtn,
		},
	}
}

func (ppg ProjectPg) Close() {
	ppg.TheDB.Close()
}

func (ppg ProjectPg) QueryStaticInfoByID(projectID int) (project.StaticInfo, error) {

}
func (ppg ProjectPg) QueryDynamicInfoByID(projectID int) (project.DynamicInfo, error) {

}

func (ppg ProjectPg) QueryInfoByOwner(userID int) ([]project.StaticInfo, []project.DynamicInfo) {

}
func (ppg ProjectPg) QueryInfoByDepartmentCode(dc string) ([]project.StaticInfo, []project.DynamicInfo) {

}
func (ppg ProjectPg) QueryAllInfo() ([]project.StaticInfo, []project.DynamicInfo) {

}

///////////////////////////////////////////////////////////////////////////////
func (ppg ProjectPg) InsertAllInfo(project.StaticInfo, project.DynamicInfo) (projectID int, err error) {

}
func (ppg ProjectPg) UpdateStaticInfo(projectInfo project.StaticInfo) (err error) {

}
func (ppg ProjectPg) UpdateDynamicInfo(projectInfo project.DynamicInfo) (err error) {

}

func (ppg ProjectPg) InnerArchiveProject(stnHistory string, dtnHistory string, projectID int) (err error) {

}
