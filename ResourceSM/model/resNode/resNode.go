package resNode

import (
	"fmt"
	"time"
)

type Node struct {
	ID            int64     `db:"node_id" json:"node_id"`
	Name          string    `db:"node_name" json:"node_name"`
	Status        int       `db:"node_status" json:"node_status"`
	Description   string    `db:"description" json:"description"`
	ProjectID     int       `db:"project_id" json:"project_id"`
	AllocatedTime time.Time `db:"allocated_time" json:"allocated_time"`
}

var SchemaInfo = `
		CREATE TABLE %s (
    		node_id int PRIMARY KEY,
			node_name varchar(256),
			node_status int,
			description varchar(256),
			project_id int,
			allocated_time TIMESTAMP WITH TIME ZONE
		);
		`
var TableName = "res_node_info"

func GetSchema() string {
	return fmt.Sprintf(SchemaInfo, TableName)
}

///////////////////////////////////////////////////////////////////////////////

const (
	StatusUnselected = 1
	StatusSelected   = 2
)
