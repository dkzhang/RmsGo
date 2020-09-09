package resAllocDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"

type ResAllocDM interface {
	QueryByID(recordID int) (rar resAlloc.Record, err error)
	QueryByProjectID(projectID int) ([]resAlloc.Record, error)
	QueryAll() ([]resAlloc.Record, error)
	Insert(record resAlloc.Record) (err error)
}
