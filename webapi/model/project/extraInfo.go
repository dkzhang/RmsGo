package project

import "time"

type ExtraInfoEx struct {
	ProjectID int `db:"project_id" json:"project_id"`
	ExtraInfo
}

type ExtraInfo struct {
	ExtraInfo string `db:"extra_info" json:"extra_info"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
