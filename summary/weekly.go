package summary

import (
	"github.com/schwarztrinker/trkbox/db"
)

type SummaryWeek struct {
	TimestampsWeek    []db.Timestamp `json:"timestamps"`
	TotalFloat        float32        `json:"differenceFloat"`
	TotalAbsoluteTime string         `json:"totalAbsoluteTime"`
	TotalWorkingDays  int            `json:"totalWorkingDays"`
	CategorySumWeek   []CategorySum
}

func GenerateSummaryByWeek(user *db.User, week string) (SummaryWeek, error) {
	ts, err := db.GetTimestampsByWeek(*user, week)
	if err != nil {
		return SummaryWeek{}, err
	}

	var categorySums = calculateCategorySumsForTimestamps(ts)

	return SummaryWeek{
		TimestampsWeek:  ts,
		CategorySumWeek: categorySums,
	}, nil

}
