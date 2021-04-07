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
		UserName:       "roc-zhj",
		ChineseName:    "张俊二级审批人",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "15383026353",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心测试用审批人",
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
		UserName:       "roc-zxq",
		ChineseName:    "翟修齐二级审批人",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "18699622740",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心测试用审批人",
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
		UserName:       "roc-w",
		ChineseName:    "温铁民",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "13833221445",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心审核",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "roc-xu",
		ChineseName:    "徐卫明",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "13933223322",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心审核",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "roc-xj",
		ChineseName:    "熊健",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "13933223435",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心审核",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "roc-zym",
		ChineseName:    "赵玉梅",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "15930217912",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心助理审核",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "roc-jy",
		ChineseName:    "贾琰",
		Department:     "资源运营中心",
		DepartmentCode: "ROC",
		Section:        "RoleApprover2",
		Mobile:         "15230464216",
		Role:           user.RoleApprover2,
		Status:         user.StatusNormal,
		Remarks:        "资源运营中心助理审核",
	})

	//////////////////////////////////////////////////////

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-jy",
		ChineseName:    "贾琰",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "RoleController",
		Mobile:         "15230464216",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zym",
		ChineseName:    "赵玉梅",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "RoleController",
		Mobile:         "15930217912",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-gzz",
		ChineseName:    "苟正忠",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "RoleController",
		Mobile:         "13930266986",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zwh",
		ChineseName:    "张卫华",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "RoleController",
		Mobile:         "13930296721",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "调度员",
	})
}
