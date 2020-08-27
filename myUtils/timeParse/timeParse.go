package timeParse

import (
	"fmt"
	"time"
)

func ParseWithLocation(timeStr string, timeLayout string, locationName string) (time.Time, error) {

	if l, err := time.LoadLocation(locationName); err != nil {
		return time.Time{}, fmt.Errorf("time.LoadLocation error: %v", err)
	} else {
		lt, _ := time.ParseInLocation(timeLayout, timeStr, l)

		//check result
		ltStr := lt.Format(timeLayout)
		if ltStr != timeStr {
			return time.Time{},
				fmt.Errorf("result check failed, expect %s but got %s", timeStr, ltStr)
		}
		return lt, nil
	}
}

func ParseDateShangHai(timeStr string) (time.Time, error) {
	timeLayout := "2006-01-02"
	locationName := "Asia/Shanghai"

	return ParseWithLocation(timeStr, timeLayout, locationName)
}

func ParseDateShangHaiS(timeStr string) (time.Time, error) {
	timeLayout := "2006-1-2"
	locationName := "Asia/Shanghai"

	return ParseWithLocation(timeStr, timeLayout, locationName)
}
