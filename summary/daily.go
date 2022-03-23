package summary

import (
	"time"

	"github.com/schwarztrinker/trkbox/db"
)

type SummaryToday struct {
	TimestampsToday                []db.Timestamp `json:"timestamps"`
	DifferenceFloat                float32        `json:"differenceFloat"`
	TotalAbsoluteTime              time.Duration  `json:"totalAbsoluteTime"`
	TotalAbsoluteTimeHumanReadable string         `json:"totalAbsoluteTimeHR"`
	Percentage                     int            `json:"percentage"`
	IsComplete                     bool           `json:"isComplete"`
	Username                       string
	Date                           string
}

func GenerateSummaryByDate(user *db.User, date string) (SummaryToday, error) {
	timestamps, err := db.GetTimestampsByDay(*user, date)
	if err != nil {
		return SummaryToday{}, err
	}

	return SummaryToday{
		TimestampsToday: timestamps, Username: user.Username,
		Date: date, IsComplete: checkinIsAlternating(timestamps),
		TotalAbsoluteTime:              calculateTotalPresenceDuration(timestamps),
		TotalAbsoluteTimeHumanReadable: calculateTotalPresenceDuration(timestamps).Round(time.Hour).String()}, nil
}

func calculateTotalPresenceDuration(ts []db.Timestamp) time.Duration {
	var absoluteTime time.Duration

	if len(ts) > 1 {

		for i, _ := range ts {
			if i == 0 {
				continue
			}

			if ts[i].IsCheckin != ts[i-1].IsCheckin && ts[i-1].IsCheckin {
				absoluteTime += ts[i].Time.Sub(ts[i-1].Time)
			}

		}
	}

	if len(ts) == 1 {
		if ts[0].IsCheckin {
			absoluteTime = time.Now().Sub(ts[0].Time)
		}
	}
	return absoluteTime
}

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
