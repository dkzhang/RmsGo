package timeParse

import "math"

func HoursToDays(h int) (d int) {
	return int(math.Round(float64(h) / 24))
}
