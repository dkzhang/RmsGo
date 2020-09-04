package resScheduling

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
)

type ResScheduling struct {
	pdm projectDM.ProjectDM
}

func NewResScheduling(pdm projectDM.ProjectDM) ResScheduling {
	return ResScheduling{
		pdm: pdm,
	}
}

func (rs ResScheduling) Scheduling(projectID int, nodes []int) (err error) {

	// TODO
	return fmt.Errorf("not accomplished")
}
