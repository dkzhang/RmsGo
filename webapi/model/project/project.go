package project

import "time"

type Project struct {
	ProjectID      int
	ProjectName    string
	ProjectChiefID int
	ApproverID     int
	ProjectInfo    string
	CreateAt       time.Time
	CloseAt        time.Time
}
