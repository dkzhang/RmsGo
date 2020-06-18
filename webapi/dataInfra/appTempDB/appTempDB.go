package appTempDB

import "github.com/dkzhang/RmsGo/webapi/model/appTemp"

type AppTempDB interface {
	QueryAppTempByOwner(userID int) (appTemp.AppTemp, error)
	QueryAppTempByID(appID int) (appTemp.AppTemp, error)

	InsertAppTemp(app appTemp.AppTemp) (int, error)
	UpdateAppTemp(app appTemp.AppTemp) error
	DeleteAppTemp(appID int) error
}
