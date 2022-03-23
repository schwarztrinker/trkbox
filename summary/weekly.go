package summary

import "github.com/schwarztrinker/trkbox/db"

type SummaryWeek struct {
	TimestampsWeek    []db.Timestamp `json:"timestamps"`
	TotalFloat        float32        `json:"differenceFloat"`
	TotalAbsoluteTime string         `json:"totalAbsoluteTime"`
	TotalWorkingDays  int            `json:"totalWorkingDays"`
}
