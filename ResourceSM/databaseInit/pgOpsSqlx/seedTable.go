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

	for ig := 287; ig <= 301; ig++ {
		for in := 1; in <= 16; in++ {
			cNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("C%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbC.Insert(cNode)
			if err != nil {
				logrus.Fatalf("Insert CpuNode (id=%d) error: %v", err)
			}
		}
	}

	{
		ig := 302
		for in := 1; in <= 12; in++ {
			cNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("C%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbC.Insert(cNode)
			if err != nil {
				logrus.Fatalf("Insert CpuNode (id=%d) error: %v", err)
			}
		}
	}

	///
	for ig := 303; ig <= 317; ig++ {
		for in := 1; in <= 16; in++ {
			cNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("C%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbC.Insert(cNode)
			if err != nil {
				logrus.Fatalf("Insert CpuNode (id=%d) error: %v", err)
			}
		}
	}

	{
		ig := 318
		for in := 1; in <= 12; in++ {
			cNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("C%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbC.Insert(cNode)
			if err != nil {
				logrus.Fatalf("Insert CpuNode (id=%d) error: %v", err)
			}
		}
	}

	////////////////////////////////////////////////////////////////////////////

	// GPU Node DB
	rndbG := resNodeDB.NewResNodePg(db, resNode.TableNameGPU)

	for ig := 283; ig <= 286; ig++ {
		for in := 1; in <= 16; in++ {
			gNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("G%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbG.Insert(gNode)
			if err != nil {
				logrus.Fatalf("Insert GpuNode (id=%d) error: %v", err)
			}
		}
	}

	for ig := 319; ig <= 322; ig++ {
		for in := 1; in <= 16; in++ {
			gNode := resNode.Node{
				ID:            int64(ig*100 + in),
				Name:          fmt.Sprintf("G%d", in),
				Status:        0,
				Description:   "",
				ProjectID:     0,
				AllocatedTime: time.Time{},
			}
			err := rndbG.Insert(gNode)
			if err != nil {
				logrus.Fatalf("Insert GpuNode (id=%d) error: %v", err)
			}
		}
	}
}
