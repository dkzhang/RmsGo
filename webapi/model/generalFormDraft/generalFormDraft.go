package generalFormDraft

type GeneralFormDraft struct {
	FormID       int    `json:"form_id" db:"form_id"`
	UserID       int    `json:"user_id" db:"user_id"`
	FormType     int    `json:"form_type" db:"form_type"`
	BasicContent string `json:"basic_content" db:"basic_content"`
	ExtraContent string `json:"extra_content" db:"extra_content"`
}

var SchemaGeneralFormDraft = `
		CREATE TABLE general_form_draft (
    		form_id_id SERIAL PRIMARY KEY,
			user_id int,
			form_type int, 
			basic_content varchar(16384),			
			extra_content varchar(16384)
		);
		`

// 16K = 1024 * 16 = 16384
