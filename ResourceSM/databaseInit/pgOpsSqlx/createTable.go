package pgOpsSqlx

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOpsSqlx"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var tableList = map[string]string{
	"res_node_cpu": resNode.GetSchemaCPU(),
	"res_node_gpu": resNode.GetSchemaGPU(),

	"res_alloc_cpu":             resAlloc.GetSchemaCPU(),
	"res_alloc_gpu":             resAlloc.GetSchemaGPU(),
	"res_alloc_storage":         resAlloc.GetSchemaStorage(),
	"history_res_alloc_cpu":     resAlloc.GetSchemaHistoryCPU(),
	"history_res_alloc_gpu":     resAlloc.GetSchemaHistoryGPU(),
	"history_res_alloc_storage": resAlloc.GetSchemaHistoryStorage(),

	"project_res_info": projectRes.GetSchema(),
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
