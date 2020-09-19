package meteringComputation

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"math"
	"time"
)

func Compute(records []resAlloc.Record, from, to time.Time) (amountInHours int, mis []metering.MeteringItem) {
	if len(records) < 1 {
		return 0, []metering.MeteringItem{}
	}

	amountInHours = 0

	// process index [0,1]~[n-2,n-1]
	for i := 0; i < len(records)-1; i++ {
		mi := metering.MeteringItem{
			Number:        records[i].NumAfter,
			AllocInfo:     records[i].AllocInfoAfterStr,
			StartDatetime: time.Time{},
			EndDatetime:   time.Time{},
			TheHours:      0,
			AmountInHours: 0,
		}

		if from.Before(records[i].CreatedAt) {
			mi.StartDatetime = records[i].CreatedAt
		} else if from.Before(records[i+1].CreatedAt) {
			mi.StartDatetime = from
		} else {
			continue
		}

		if to.After(records[i+1].CreatedAt) {
			mi.EndDatetime = records[i+1].CreatedAt
		} else if to.After(records[i].CreatedAt) {
			mi.EndDatetime = to
		} else {
			continue
		}

		mi.TheHours = int(math.Round(mi.EndDatetime.Sub(mi.StartDatetime).Hours()))
		mi.AmountInHours = mi.Number * mi.TheHours
		mis = append(mis, mi)
		amountInHours += mi.AmountInHours
	}

	// process index n-1
	lastRecord := records[len(records)-1]
	if lastRecord.NumAfter != 0 {
		// if lastRecord.NumAfter == 0, there is nothing to do

		if from.Before(lastRecord.CreatedAt) && to.After(lastRecord.CreatedAt) {
			// lastRecord.CreatedAt in the range [from, to]
			mi := metering.MeteringItem{
				Number:        lastRecord.NumAfter,
				AllocInfo:     lastRecord.AllocInfoAfterStr,
				StartDatetime: time.Time{},
				EndDatetime:   time.Time{},
				TheHours:      0,
				AmountInHours: 0,
			}

			mi.StartDatetime = lastRecord.CreatedAt
			mi.EndDatetime = to

			mi.TheHours = int(math.Round(mi.EndDatetime.Sub(mi.StartDatetime).Hours()))
			mi.AmountInHours = mi.Number * mi.TheHours
			mis = append(mis, mi)
			amountInHours += mi.AmountInHours
		}
	}

	return amountInHours, mis
}
