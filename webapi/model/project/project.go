package project

import (
	"time"
)

type StaticInfo struct {
	ProjectID        int       `db:"project_id" json:"project_id"`
	ProjectName      string    `db:"project_name" json:"project_name"`
	ProjectCode      string    `db:"project_code" json:"project_code"`
	DepartmentCode   string    `db:"department_code" json:"department_code"`
	Department       string    `db:"department" json:"department"`
	ChiefID          int       `db:"chief_id" json:"chief_id"`
	ChiefChineseName string    `db:"chief_cn_name" json:"chief_cn_name"`
	ExtraInfo        string    `db:"extra_info" json:"extra_info"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

var SchemaStaticInfo = `
		CREATE TABLE user_info (
    		project_id SERIAL PRIMARY KEY,
			project_name varchar(256) ,
			project_code varchar(32) UNIQUE,
			department_code varchar(32),
			department varchar(256),
			chief_id int,
			chief_cn_name varchar(32), 
			basic_content varchar(16384),			
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE
		);
		`
