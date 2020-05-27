package user

type UserInfo struct {
	//系统内ID，主键
	UserID int `db:"user_id" json:"user_id"`

	//登录名，含所属单位前缀，唯一
	UserName string `db:"user_name" json:"user_name"`

	ChineseName string `db:"chinese_name" json:"chinese_name"`

	Department     string `db:"department" json:"department"`
	DepartmentCode string `db:"department_code" json:"department_code"`
	Section        string `db:"section" json:"section"`

	Mobile string `db:"mobile" db:"mobile"`

	//项目长=1，审批人=2，调度员=9
	Role int `db:"role" json:"role"`

	//正常=1，删除=-1，停用=-2
	Status int `db:"status" json:"status"`

	Remarks string `db:"remarks" json:"remarks"`
}

var SchemaUser = `
		CREATE TABLE user_info (
    		user_id SERIAL PRIMARY KEY,
			user_name varchar(32) UNIQUE,
			chinese_name varchar(256), 
			department varchar(256),
			department_code varchar(32),
			section varchar(256),			
			mobile varchar(32),			
			role int,
			status int,			
			remarks varchar(256)
		);
		`

const (
	RoleProjectChief = 1
	RoleApprover     = 2
	RoleController   = 7
)

const (
	StatusNormal  = 1
	StatusDisable = -1
	StatusDelete  = -9
)
