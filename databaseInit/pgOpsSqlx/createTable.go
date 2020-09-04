package pgOpsSqlx

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOpsSqlx"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/generalFormDraft"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var tableList = map[string]string{
	"user_info":                 user.SchemaUser,
	"general_form_draft":        generalFormDraft.SchemaGeneralFormDraft,
	"application":               application.GetSchemaApp(),
	"history_application":       application.GetSchemaHistoryApp(),
	"app_ops_record":            application.GetSchemaAppOps(),
	"history_app_ops_record":    application.GetSchemaHistoryAppOps(),
	"project_info":              project.GetSchema(),
	"history_project_info":      project.GetSchemaHistory(),
	"res_node_info":             resNode.GetSchema(),
	"res_alloc_cpu":             resAlloc.GetSchemaCPU(),
	"res_alloc_gpu":             resAlloc.GetSchemaGPU(),
	"res_alloc_storage":         resAlloc.GetSchemaStorage(),
	"history_res_alloc_cpu":     resAlloc.GetSchemaHistoryCPU(),
	"history_res_alloc_gpu":     resAlloc.GetSchemaHistoryGPU(),
	"history_res_alloc_storage": resAlloc.GetSchemaHistoryStorage(),
}

func CreateAllTable(db *sqlx.DB) {
	for name, scheme := range tableList {
		_, err := postgreOpsSqlx.DropTable(db, name)
		if err != nil {
			logrus.Errorf("Drop table <%s> error: %v", name, err)
		} else {
			logrus.Infof("Drop table <%s> success", name)
		}

		_, err = postgreOpsSqlx.CreateTable(db, scheme)
		if err != nil {
			logrus.Errorf("Create table <%s> error: %v", name, err)
		} else {
			logrus.Infof("Create table <%s> success", name)
		}
	}
}
