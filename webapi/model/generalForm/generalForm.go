package generalForm

type GeneralForm struct {
	ProjectID          int         `json:"project_id"`
	FormID             int         `json:"form_id"`
	Action             int         `json:"action"`
	BasicContent       string      `json:"basic_content"`
	BasicContentStruct interface{} `json:"-"`
	ExtraContent       string      `json:"extra_content"`
}
