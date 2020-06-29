package generalFormDraftDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
)

type GeneralFormDraftDB interface {
	QueryGeneralFormDraftByOwner(userID int) ([]generalFormDraft.GeneralFormDraft, error)
	QueryGeneralFormDraftByID(formID int) (generalFormDraft.GeneralFormDraft, error)

	InsertGeneralFormDraft(gfd generalFormDraft.GeneralFormDraft) (int, error)
	UpdateGeneralFormDraft(gfd generalFormDraft.GeneralFormDraft) error
	DeleteGeneralFormDraft(formID int) error
}
