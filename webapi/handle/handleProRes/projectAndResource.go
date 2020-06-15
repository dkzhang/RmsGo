package handleProRes

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/gin-gonic/gin"
	"time"
)

type AppNewProResWA struct {
	ProjectID         int            `json:"project_id"`
	ApplicationID     int            `json:"application_id"`
	ProjectName       string         `json:"project_name"`
	TheResApplication ResApplication `json:"res_application"`
	Action            int            `json:"action"`
	ExtraInfo         string         `json:"extra_info"`
}

type ResApplication struct {
	resource.Resource
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func ApplyNewProRes(c *gin.Context) {

}

type AppExResWA struct {
}

type AppReComWA struct {
}

type AppReStoWA struct {
}

type AppMeterWA struct {
}
