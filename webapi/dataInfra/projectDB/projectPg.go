package projectDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
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

func (ppg ProjectPg) QueryStaticInfoByID(projectID int) (psi project.StaticInfo, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Get(&psi, queryByID, projectID)
	if err != nil {
		return project.StaticInfo{},
			fmt.Errorf("QueryStaticInfoByID project.StaticInfo in TheDB error: %v", err)
	}
	return psi, nil
}
func (ppg ProjectPg) QueryDynamicInfoByID(projectID int) (pdi project.DynamicInfo, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE project_id=$1`, ppg.DynamicTableName)
	err = ppg.TheDB.Get(&pdi, queryByID, projectID)
	if err != nil {
		return project.DynamicInfo{},
			fmt.Errorf("QueryStaticInfoByID project.DynamicInfo in TheDB error: %v", err)
	}
	return pdi, nil
}

func (ppg ProjectPg) QueryInfoByOwner(userID int) (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE chief_id=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryByOwner, userID)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByOwner project.StaticInfo (userID=%d) from TheDB error: %v", err, userID)
	}

	queryByOwnerJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id) WHERE %s.chief_id=$1`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName, ppg.StaticTableName)
	err = ppg.TheDB.Select(&pdis, queryByOwnerJoin, userID)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByOwner project.DynamicInfo (userID=%d) from TheDB error: %v", err, userID)
	}

	return psis, pdis, nil
}

func (ppg ProjectPg) QueryInfoByDepartmentCode(dc string) (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryByDC := fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1`, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryByDC, dc)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByDepartmentCode project.StaticInfo (DepartmentCode=%s) from TheDB error: %v", err, dc)
	}

	queryByDCJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id) WHERE %s.department_code=$1`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName, ppg.StaticTableName)
	err = ppg.TheDB.Select(&pdis, queryByDCJoin, dc)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByDepartmentCode project.DynamicInfo (DepartmentCode=%s) from TheDB error: %v", err, dc)
	}

	return psis, pdis, nil
}
func (ppg ProjectPg) QueryAllInfo() (psis []project.StaticInfo, pdis []project.DynamicInfo, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s `, ppg.StaticTableName)
	err = ppg.TheDB.Select(&psis, queryAll)
	if err != nil {
		return nil, nil,
			fmt.Errorf("QueryInfoByOwner project.StaticInfo ALL from TheDB error: %v", err)
	}

	queryAllJoin := fmt.Sprintf(`SELECT %s.* FROM %s JOIN %s USING(project_id)`,
		ppg.DynamicTableName, ppg.StaticTableName, ppg.DynamicTableName)
	err = ppg.TheDB.Select(&pdis, queryAllJoin)
	if err != nil {
		return psis, nil,
			fmt.Errorf("QueryInfoByOwner project.DynamicInfo ALL from TheDB error: %v", err)
	}

	return psis, pdis, nil
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
