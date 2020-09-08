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

func (rad ResAllocDirectDB) Insert(ra resAlloc.Record) (err error) {
	ra.CreatedAt = time.Now()

	// avoid slice to be nil
	if ra.AllocInfoBefore == nil {
		ra.AllocInfoBefore = make([]int64, 0)
	}
	if ra.AllocInfoAfter == nil {
		ra.AllocInfoAfter = make([]int64, 0)
	}
	if ra.AllocInfoChange == nil {
		ra.AllocInfoChange = make([]int64, 0)
	}

	return rad.theResAllocDB.Insert(ra)
}
