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
	err = atpg.db.Select(&apptemps, `SELECT * FROM application_temporary WHERE user_id=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("get all user info from db error: %v", err)
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
	err = atpg.db.Get(&id, `INSERT INTO application_temporary (user_id, app_type, basic_content, extra_content) VALUES ($1) RETURNING application_id`,
		app.UserID, app.AppType, app.BasicContent, app.ExtraContent)
	if err != nil {
		return -1, fmt.Errorf("InsertAppTemp in db error: %v", err)
	}
	return id, nil
}

func (atpg AppTempPg) UpdateAppTemp(app appTemp.AppTemp) error {

	return nil
}

func (atpg AppTempPg) DeleteAppTemp(appID int) error {
	return nil
}
