package summary

import (
	"time"

	"github.com/schwarztrinker/trkbox/db"
)

type SummaryToday struct {
	TimestampsToday           []db.Timestamp `json:"timestamps"`
	DifferenceFloat           float32        `json:"differenceFloat"`
	TotalTimeDurationTime     time.Duration  `json:"totalTimeDurationTime"`
	TotalTimeDurationReadable string         `json:"totalTimeDurationReadable"`
	Percentage                int            `json:"percentage"`
	Username                  string
	Date                      string
	CurrentAttendanceTime     time.Duration `json:"currentAttendanceTime"`
	CategorySums              []CategorySum `json:"categorysums"`
}

func GenerateSummaryByDate(user *db.User, date string) (SummaryToday, error) {
	timestamps, err := db.GetTimestampsByDay(*user, date)
	if err != nil {
		return SummaryToday{}, err
	}

	return SummaryToday{
		Percentage:      int((calculateTotalPresenceDuration(timestamps).Hours() / 8) * 100),
		DifferenceFloat: float32(calculateTotalPresenceDuration(timestamps).Hours()),
		TimestampsToday: timestamps, Username: user.Username,
		Date:                      date,
		TotalTimeDurationTime:     calculateTotalPresenceDuration(timestamps),
		TotalTimeDurationReadable: calculateTotalPresenceDuration(timestamps).String(),
		CategorySums:              calculateCategorySumsForTimestamps(timestamps)}, nil

}

func calculateTotalPresenceDuration(ts []db.Timestamp) time.Duration {
	var absoluteTime time.Duration

	// Calculate if ts more than 1
	if len(ts) > 1 {
		for i := range ts {
			if i == 0 {
				continue
			}

			if ts[i].IsCheckin != ts[i-1].IsCheckin && ts[i-1].IsCheckin {
				absoluteTime += ts[i].Time.Sub(ts[i-1].Time)
			}

		}
	}
	return absoluteTime
}

// helper functions to check if calculation of timestamps could be completeted
func checkinIsAlternating(ts []db.Timestamp) bool {
	var last bool
	for i := range ts {
		if last == ts[i].IsCheckin || len(ts) <= 1 {
			return false
		}

		last = ts[i].IsCheckin
	}
	return true
}

func calculateIsComplete(ts []db.Timestamp) bool {
	if len(ts)%2 == 1 || len(ts) <= 1 || !checkinIsAlternating(ts) {
		return false
	}
	return true
}
