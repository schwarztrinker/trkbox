package summary

import (
	"time"

	"github.com/schwarztrinker/trkbox/db"
)

type CategorySum struct {
	Category             string        `json:"category"`
	TimeDuration         time.Duration `json:"timeDuration"`
	TimeDurationReadable string        `json:"timeDurationReadable"`
	CategoryStampCounts  int           `json:"categoryStampCounts"`
	IsComplete           bool
}

// logic for calculating category sums
func calculateCategorySumsForTimestamps(ts []db.Timestamp) []CategorySum {
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
				CategoryStampCounts:  len(filterTimestampsByCategory(ts, timestamp.Category)),
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
