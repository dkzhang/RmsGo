package pgOpsSqlx

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resNodeDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resNode"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

func SeedAllTable(db *sqlx.DB) {
	seedNodeTable(db)
}

func seedNodeTable(db *sqlx.DB) {
	// CPU Node DB
	rndbC := resNodeDB.NewResNodePg(db, resNode.TableNameCPU)
	for i := int64(1); i <= 256; i++ {
		ni := resNode.Node{
			ID:            i,
			Name:          fmt.Sprintf("CpuNode%d", i),
			Status:        0,
			Description:   "",
			ProjectID:     0,
			AllocatedTime: time.Time{},
		}
		err := rndbC.Insert(ni)
		if err != nil {
			logrus.Fatalf("Insert CpuNode (id=%d) error: %v", err)
		}
	}

	// GPU Node DB
	rndbG := resNodeDB.NewResNodePg(db, resNode.TableNameGPU)
	for i := int64(1); i <= 66; i++ {
		ni := resNode.Node{
			ID:            i,
			Name:          fmt.Sprintf("GpuNode%d", i),
			Status:        0,
			Description:   "",
			ProjectID:     0,
			AllocatedTime: time.Time{},
		}
		err := rndbG.Insert(ni)
		if err != nil {
			logrus.Fatalf("Insert GpuNode (id=%d) error: %v", err)
		}
	}

}
