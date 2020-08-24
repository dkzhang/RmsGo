package pgOps

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var tableList = map[string]string{
	"user_info":               user.SchemaUser,
	"general_form_draft":      generalFormDraft.SchemaGeneralFormDraft,
	"application":             application.GetSchemaApplication("application"),
	"history_application":     application.GetSchemaApplication("history_application"),
	"app_ops_record":          application.GetSchemaAppOpsRecord("app_ops_record"),
	"history_app_ops_record":  application.GetSchemaAppOpsRecord("history_app_ops_record"),
	"project_static":          project.GetSchemaStatic("project_static"),
	"project_dynamic":         project.GetSchemaStatic("project_dynamic"),
	"history_project_static":  project.GetSchemaStatic("history_project_static"),
	"history_project_dynamic": project.GetSchemaStatic("history_project_dynamic"),
}

func CreateAllTable(db *sqlx.DB) {
	for name, scheme := range tableList {
		_, err := postgreOps.DropTable(db, name)
		if err != nil {
			logrus.Errorf("Drop table <%s> error: %v", name, err)
		} else {
			logrus.Infof("Drop table <%s> success", name)
		}

		_, err = postgreOps.CreateTable(db, scheme)
		if err != nil {
			logrus.Errorf("Create table <%s> error: %v", name, err)
		} else {
			logrus.Infof("Create table <%s> success", name)
		}
	}
}

func SeedAllTable(db *sqlx.DB) {
	seedUserTable(db)
}

func ImportFromFile(db *sqlx.DB) {

}

func seedUserTable(db *sqlx.DB) {
	theUserDB := userDB.NewUserInPostgre(db)
	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zhj007",
		ChineseName:    "张俊007",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试RoleController",
		Mobile:         "15383026353",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用调度员",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "jf-zhj002",
		ChineseName:    "张俊002",
		Department:     "计服中心",
		DepartmentCode: "jf",
		Section:        "测试RoleApprover",
		Mobile:         "15383026353",
		Role:           user.RoleApprover,
		Status:         user.StatusNormal,
		Remarks:        "测试用审批人",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "jf-zhj001",
		ChineseName:    "张俊001",
		Department:     "计服中心",
		DepartmentCode: "jf",
		Section:        "测试",
		Mobile:         "15383026353",
		Role:           user.RoleProjectChief,
		Status:         user.StatusNormal,
		Remarks:        "测试用项目长",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zxq",
		ChineseName:    "翟修齐",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试",
		Mobile:         "18699622740",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用账号for翟",
	})
}
