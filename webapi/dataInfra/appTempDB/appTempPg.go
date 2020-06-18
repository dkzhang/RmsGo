package appTempDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/appTemp"
	"github.com/jmoiron/sqlx"
)

type AppTempPg struct {
	db *sqlx.DB
}

func NewAppTempPg(db *sqlx.DB) AppTempPg {
	return AppTempPg{db: db}
}

func (atpg AppTempPg) QueryAppTempByOwner(userID int) (apptemps []appTemp.AppTemp, err error) {
	queryByOwner := `SELECT * FROM application_temporary WHERE user_id=$1`
	err = atpg.db.Select(&apptemps, queryByOwner, userID)
	if err != nil {
		return nil, fmt.Errorf("QueryAppTempByOwner from db error: %v", err)
	}
	return apptemps, nil
}

func (atpg AppTempPg) QueryAppTempByID(appID int) (apptemp appTemp.AppTemp, err error) {
	err = atpg.db.Get(&apptemp, `SELECT * FROM application_temporary WHERE application_id=$1`, appID)
	if err != nil {
		return appTemp.AppTemp{}, fmt.Errorf("QueryAppTempByID in db error: %v", err)
	}
	return apptemp, nil
}

func (atpg AppTempPg) InsertAppTemp(app appTemp.AppTemp) (id int, err error) {
	err = atpg.db.Get(&id, `INSERT INTO application_temporary (user_id, app_type, basic_content, extra_content) VALUES ($1, $2, $3, $4) RETURNING application_id`,
		app.UserID, app.AppType, app.BasicContent, app.ExtraContent)
	if err != nil {
		return -1, fmt.Errorf("InsertAppTemp in db error: %v", err)
	}
	return id, nil
}

func (atpg AppTempPg) UpdateAppTemp(app appTemp.AppTemp) error {
	_, err := atpg.db.NamedExec("UPDATE application_temporary "+
		"SET user_id=:user_id, app_type=:app_type, "+
		"basic_content=:basic_content, extra_content=:extra_content "+
		"WHERE application_id=:application_id", app)
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE application_temporary: %v", err)
	}
	return nil
}

func (atpg AppTempPg) DeleteAppTemp(appID int) error {
	deleteAppTemp := `DELETE FROM application_temporary WHERE application_id=$1`

	result, err := atpg.db.Exec(deleteAppTemp, appID)
	if err != nil {
		return fmt.Errorf("db.Exec(deleteAppTemp, appID), userID = %d", appID)
	}
	fmt.Printf("DeleteAppTemp success: %v \n", result)
	return nil
}
