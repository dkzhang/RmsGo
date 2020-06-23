package statusActionMap

import "github.com/dkzhang/RmsGo/webapi/model/generalForm"

type StatusActionItem struct {
	Status    int
	Role      int
	Action    int
	ActionStr string
	Execute   func(gl generalForm.GeneralForm)
}

type StatusActionMap []StatusActionItem

func (sam StatusActionMap) Execute(status, role, action int, gl generalForm.GeneralForm) {
	for _, smi := range sam {
		if status == smi.Status && role == smi.Role && action == smi.Action {
			smi.Execute(gl)
			return
		}
	}
	// error: no match
}
