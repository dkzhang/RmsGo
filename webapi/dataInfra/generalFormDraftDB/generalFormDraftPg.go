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

func (atpg GeneralFormDraftPg) QueryGeneralFormDraftByOwner(userID int) (gfds []generalFormDraft.GeneralFormDraft, err error) {
	queryByOwner := `SELECT * FROM general_form_draft WHERE user_id=$1`
	err = atpg.db.Select(&gfds, queryByOwner, userID)
	if err != nil {
		return nil, fmt.Errorf("QueryGeneralFormDraftByOwner from db error: %v", err)
	}
	return gfds, nil
}

func (atpg GeneralFormDraftPg) QueryGeneralFormDraftByID(formID int) (gfd generalFormDraft.GeneralFormDraft, err error) {
	err = atpg.db.Get(&gfd, `SELECT * FROM general_form_draft WHERE form_id=$1`, formID)
	if err != nil {
		return generalFormDraft.GeneralFormDraft{}, fmt.Errorf("QueryGeneralFormDraftByID in db error: %v", err)
	}
	return gfd, nil
}

func (atpg GeneralFormDraftPg) InsertGeneralFormDraft(gfd generalFormDraft.GeneralFormDraft) (id int, err error) {
	err = atpg.db.Get(&id, `INSERT INTO general_form_draft (user_id, form_type, basic_content, extra_content) VALUES ($1, $2, $3, $4) RETURNING form_id`,
		gfd.UserID, gfd.FormType, gfd.BasicContent, gfd.ExtraContent)
	if err != nil {
		return -1, fmt.Errorf("InsertGeneralFormDraft in db error: %v", err)
	}
	return id, nil
}

func (atpg GeneralFormDraftPg) UpdateGeneralFormDraft(gfd generalFormDraft.GeneralFormDraft) error {
	_, err := atpg.db.NamedExec("UPDATE general_form_draft "+
		"SET user_id=:user_id, form_type=:form_type, "+
		"basic_content=:basic_content, extra_content=:extra_content "+
		"WHERE form_id=:form_id", gfd)
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE general_form_draft error: %v", err)
	}
	return nil
}

func (atpg GeneralFormDraftPg) DeleteGeneralFormDraft(formID int) error {
	deleteGeneralFormDraft := `DELETE FROM general_form_draft WHERE form_id=$1`

	result, err := atpg.db.Exec(deleteGeneralFormDraft, formID)
	if err != nil {
		return fmt.Errorf("db.Exec(deleteGeneralFormDraft, formID), userID = %d", formID)
	}
	fmt.Printf("DeleteGeneralFormDraft success: %v \n", result)
	return nil
}
