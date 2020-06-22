package generalFormDraftDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
	"github.com/jmoiron/sqlx"
)

type GeneralFormDraftPg struct {
	db *sqlx.DB
}

func NewGeneralFormDraftPg(db *sqlx.DB) GeneralFormDraftPg {
	return GeneralFormDraftPg{db: db}
}

func (atpg GeneralFormDraftPg) QueryGeneralFormDraftByOwner(userID int) (GeneralFormDrafts []generalFormDraft.GeneralFormDraft, err error) {
	queryByOwner := `SELECT * FROM general_form_draft WHERE user_id=$1`
	err = atpg.db.Select(&GeneralFormDrafts, queryByOwner, userID)
	if err != nil {
		return nil, fmt.Errorf("QueryGeneralFormDraftByOwner from db error: %v", err)
	}
	return GeneralFormDrafts, nil
}

func (atpg GeneralFormDraftPg) QueryGeneralFormDraftByID(appID int) (GeneralFormDraft generalFormDraft.GeneralFormDraft, err error) {
	err = atpg.db.Get(&GeneralFormDraft, `SELECT * FROM general_form_draft WHERE application_id=$1`, appID)
	if err != nil {
		return generalFormDraft.GeneralFormDraft{}, fmt.Errorf("QueryGeneralFormDraftByID in db error: %v", err)
	}
	return GeneralFormDraft, nil
}

func (atpg GeneralFormDraftPg) InsertGeneralFormDraft(app generalFormDraft.GeneralFormDraft) (id int, err error) {
	err = atpg.db.Get(&id, `INSERT INTO general_form_draft (user_id, app_type, basic_content, extra_content) VALUES ($1, $2, $3, $4) RETURNING application_id`,
		app.UserID, app.AppType, app.BasicContent, app.ExtraContent)
	if err != nil {
		return -1, fmt.Errorf("InsertGeneralFormDraft in db error: %v", err)
	}
	return id, nil
}

func (atpg GeneralFormDraftPg) UpdateGeneralFormDraft(app generalFormDraft.GeneralFormDraft) error {
	_, err := atpg.db.NamedExec("UPDATE general_form_draft "+
		"SET user_id=:user_id, app_type=:app_type, "+
		"basic_content=:basic_content, extra_content=:extra_content "+
		"WHERE application_id=:application_id", app)
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE general_form_draft: %v", err)
	}
	return nil
}

func (atpg GeneralFormDraftPg) DeleteGeneralFormDraft(appID int) error {
	deleteGeneralFormDraft := `DELETE FROM general_form_draft WHERE application_id=$1`

	result, err := atpg.db.Exec(deleteGeneralFormDraft, appID)
	if err != nil {
		return fmt.Errorf("db.Exec(deleteGeneralFormDraft, appID), userID = %d", appID)
	}
	fmt.Printf("DeleteGeneralFormDraft success: %v \n", result)
	return nil
}
