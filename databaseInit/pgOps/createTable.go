package pgOps

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
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
	"application":             application.GetSchemaApp(),
	"history_application":     application.GetSchemaHistoryApp(),
	"app_ops_record":          application.GetSchemaAppOps(),
	"history_app_ops_record":  application.GetSchemaHistoryAppOps(),
	"project_static":          project.GetSchemaStatic(),
	"project_dynamic":         project.GetSchemaDynamic(),
	"history_project_static":  project.GetSchemaHistoryStatic(),
	"history_project_dynamic": project.GetSchemaHistoryDynamic(),
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
