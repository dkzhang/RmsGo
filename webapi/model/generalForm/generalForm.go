package generalForm

type GeneralForm struct {
	ProjectID    int    `json:"project_id"`
	FormID       int    `json:"form_id"`
	Type         int    `json:"form_type"`
	Action       int    `json:"action"`
	BasicContent string `json:"basic_content"`
	ExtraContent string `json:"extra_content"`
}
