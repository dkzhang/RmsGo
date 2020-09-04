package pgOpsSqlx

import (
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
)

func SeedAllTable(db *sqlx.DB) {
	seedUserTable(db)
}

func ImportFromFile(db *sqlx.DB) {

}

func seedUserTable(db *sqlx.DB) {
	theUserDB := userDB.NewUserInPostgre(db)
	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zhj",
		ChineseName:    "张俊项调度",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试RoleController",
		Mobile:         "15383026353",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "app-zhj",
		ChineseName:    "张俊审批人",
		Department:     "计服中心1",
		DepartmentCode: "jf1",
		Section:        "测试RoleApprover",
		Mobile:         "15383026353",
		Role:           user.RoleApprover,
		Status:         user.StatusNormal,
		Remarks:        "测试用审批人",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "pc-zhj",
		ChineseName:    "张俊项目长",
		Department:     "计服中心1",
		DepartmentCode: "jf1",
		Section:        "测试RoleProjectChief",
		Mobile:         "15383026353",
		Role:           user.RoleProjectChief,
		Status:         user.StatusNormal,
		Remarks:        "测试用项目长",
	})

	//////////////////////////////////////////////////////

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zxq",
		ChineseName:    "翟修齐调度",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试RoleController",
		Mobile:         "18699622740",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "app-zxq",
		ChineseName:    "翟修齐审批人",
		Department:     "计服中心2",
		DepartmentCode: "jf2",
		Section:        "测试RoleApprover",
		Mobile:         "18699622740",
		Role:           user.RoleApprover,
		Status:         user.StatusNormal,
		Remarks:        "测试用审批人",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "pc-zxq",
		ChineseName:    "翟修齐项目长",
		Department:     "计服中心2",
		DepartmentCode: "jf2",
		Section:        "测试RoleProjectChief",
		Mobile:         "18699622740",
		Role:           user.RoleProjectChief,
		Status:         user.StatusNormal,
		Remarks:        "测试用项目长",
	})
	//////////////////////////////////////////////////////

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-lhs",
		ChineseName:    "李华松调度",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试RoleController",
		Mobile:         "18617772241",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "app-lhs",
		ChineseName:    "李华松审批人",
		Department:     "计服中心3",
		DepartmentCode: "jf3",
		Section:        "测试RoleApprover",
		Mobile:         "18617772241",
		Role:           user.RoleApprover,
		Status:         user.StatusNormal,
		Remarks:        "测试用审批人",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "pc-lhs",
		ChineseName:    "李华松项目长",
		Department:     "计服中心3",
		DepartmentCode: "jf3",
		Section:        "测试RoleProjectChief",
		Mobile:         "18617772241",
		Role:           user.RoleProjectChief,
		Status:         user.StatusNormal,
		Remarks:        "测试用项目长",
	})
}
