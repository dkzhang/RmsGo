package workflow

import "github.com/dkzhang/RmsGo/webapi/model/generalForm"

type GeneralWorkflow interface {
	Apply(form generalForm.GeneralForm) (err error)
	Process(form generalForm.GeneralForm) (err error)
}
