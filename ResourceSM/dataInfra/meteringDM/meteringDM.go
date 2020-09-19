package meteringDM

import "github.com/dkzhang/RmsGo/ResourceSM/model/metering"

type MeteringDM interface {
	MeteringReadOnlyDM

	QueryWithCreate(projectID int, mType int, typeInfo string) (ms metering.Statement, err error)

	//ComputeMonthly()
	//ComputeQuarterly()
	//ComputeAnnual()
	//ComputeAnyPeriod()

	ComputeSettlement(projectID int) (ms metering.Statement, err error)
}

type MeteringReadOnlyDM interface {
	Query(projectID int, mType int, typeInfo string) (ms metering.Statement, err error)
	QueryAll(projectID int, mType int) (mss []metering.Statement, err error)
}
