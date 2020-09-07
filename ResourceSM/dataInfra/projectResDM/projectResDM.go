package projectResDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"

type ProjectResDM interface {
	QueryByID(projectID int) (pr projectRes.ResInfo, err error)
	QueryAll() ([]projectRes.ResInfo, error)
	Insert(pr projectRes.ResInfo) (err error)
	Update(pr projectRes.ResInfo) (err error)
	Delete(projectID int) (err error)
}
