package appTemp

type AppTemp struct {
	ApplicationID int    `json:"application_id" db:"application_id"`
	UserID        int    `json:"user_id" 		db:"user_id"`
	AppType       int    `json:"app_type" 		db:"app_type"`
	BasicContent  string `json:"basic_content" 	db:"basic_content"`
	ExtraContent  string `json:"extra_content" 	db:"extra_content"`
}

var SchemaAppTemp = `
		CREATE TABLE application_temporary (
    		application_id SERIAL PRIMARY KEY,
			user_id int,
			app_type int, 
			basic_content varchar(16384),			
			extra_content varchar(16384)
		);
		`

// 16K = 1024 * 16 = 16384
