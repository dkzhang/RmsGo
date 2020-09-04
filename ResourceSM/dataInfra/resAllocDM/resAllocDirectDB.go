package resAllocDM

import (
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDB"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"time"
)

type ResAllocDirectDB struct {
	theResAllocDB resAllocDB.ResAllocDB
}

func NewResAllocDirectDB(radb resAllocDB.ResAllocDB) ResAllocDirectDB {
	return ResAllocDirectDB{
		theResAllocDB: radb,
	}
}

func (rad ResAllocDirectDB) QueryByID(recordID int) (rar resAlloc.Record, err error) {
	return rad.theResAllocDB.QueryByID(recordID)
}

func (rad ResAllocDirectDB) QueryByProjectID(projectID int) ([]resAlloc.Record, error) {
	return rad.theResAllocDB.QueryByProjectID(projectID)
}
func (rad ResAllocDirectDB) QueryAll() ([]resAlloc.Record, error) {
	return rad.theResAllocDB.QueryAll()
}

func (rad ResAllocDirectDB) Insert(node resAlloc.Record) (err error) {
	node.CreatedAt = time.Now()
	return rad.theResAllocDB.Insert(node)
}
