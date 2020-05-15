package mostfrequentdays

import (
	"time"
)

func MostFrequentDays(year int) []string {
	beg := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()
	end := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).Weekday()
	if beg == end {
		return []string{beg.String()}
	}
	if beg == time.Sunday {
		beg, end = end, beg
	}
	return []string{beg.String(), end.String()}
}
