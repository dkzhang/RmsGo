package generalFormDraftDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
)

type GeneralFormDraftDB interface {
	QueryGeneralFormDraftByOwner(userID int) ([]generalFormDraft.GeneralFormDraft, error)
	QueryGeneralFormDraftByID(appID int) (generalFormDraft.GeneralFormDraft, error)

	InsertGeneralFormDraft(app generalFormDraft.GeneralFormDraft) (int, error)
	UpdateGeneralFormDraft(app generalFormDraft.GeneralFormDraft) error
	DeleteGeneralFormDraft(appID int) error
}
