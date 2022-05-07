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
	IsComplete                bool           `json:"isComplete"`
	Username                  string
	Date                      string
	CurrentAttendanceTime     time.Duration `json:"currentAttendanceTime"`
	CategorySums              []CategorySum `json:"categorysums"`
}

type CategorySum struct {
	Category             string
	TimeDuration         time.Duration
	TimeDurationReadable string
	StampCounts          int
	IsComplete           bool
}

func GenerateSummaryByDate(user *db.User, date string) (SummaryToday, error) {
	timestamps, err := db.GetTimestampsByDay(*user, date)
	if err != nil {
		return SummaryToday{}, err
	}

	return SummaryToday{
		IsComplete:      calculateIsComplete(timestamps),
		Percentage:      int((calculateTotalPresenceDuration(timestamps).Hours() / 8) * 100),
		DifferenceFloat: float32(calculateTotalPresenceDuration(timestamps).Hours()),
		TimestampsToday: timestamps, Username: user.Username,
		Date:                      date,
		TotalTimeDurationTime:     calculateTotalPresenceDuration(timestamps),
		TotalTimeDurationReadable: calculateTotalPresenceDuration(timestamps).String(),
		CategorySums:              calculateCategorySumsToday(timestamps)}, nil

}

func calculateCategorySumsToday(ts []db.Timestamp) []CategorySum {
	var categorysums []CategorySum

	for _, timestamp := range ts {

		var categoryExists bool = false
		for _, categorysum := range categorysums {
			if categorysum.Category == timestamp.Category {
				categoryExists = true
			}
		}

		if !categoryExists {
			var returnCategorySum CategorySum
			var filteredTimestamps = filterTimestampsByCategory(ts, timestamp.Category)

			// calculate time for the category
			var timeCategory time.Duration

			if len(filteredTimestamps) > 1 {
				for i := range filteredTimestamps {
					if i == 0 {
						continue
					}

					if filteredTimestamps[i].IsCheckin != filteredTimestamps[i-1].IsCheckin && filteredTimestamps[i-1].IsCheckin {
						timeCategory += filteredTimestamps[i].Time.Sub(filteredTimestamps[i-1].Time)
					}

				}
			}

			returnCategorySum = CategorySum{
				Category:             timestamp.Category,
				TimeDuration:         timeCategory,
				TimeDurationReadable: timeCategory.String(),
				StampCounts:          len(filterTimestampsByCategory(ts, timestamp.Category)),
				IsComplete:           calculateIsComplete(filteredTimestamps),
			}

			categorysums = append(categorysums, returnCategorySum)
		}
	}
	return categorysums
}

// helper function to filter timestamps by categoy
func filterTimestampsByCategory(ts []db.Timestamp, category string) []db.Timestamp {
	var out []db.Timestamp
	for _, timestamp := range ts {
		if timestamp.Category == category {
			out = append(out, timestamp)
		}
	}
	return out
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
