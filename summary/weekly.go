package summary

import (
	"github.com/schwarztrinker/trkbox/db"
)

type SummaryWeek struct {
	TimestampsWeek    []db.Timestamp `json:"timestamps"`
	TotalFloat        float32        `json:"differenceFloat"`
	TotalAbsoluteTime string         `json:"totalAbsoluteTime"`
	TotalWorkingDays  int            `json:"totalWorkingDays"`
}

func GenerateSummaryByWeek(user *db.User, week string) (SummaryWeek, error) {
	ts, err := db.GetTimestampsByWeek(*user, week)
	if err != nil {
		return SummaryWeek{}, err
	}

	return SummaryWeek{
		TimestampsWeek: ts,
	}, nil

}
