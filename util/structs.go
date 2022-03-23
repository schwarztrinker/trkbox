package util

import (
	"time"

	"github.com/schwarztrinker/trkbox/db"
)

type SummaryToday struct {
	TimestampsToday   []db.Timestamp `json:"timestamps"`
	DifferenceFloat   float32        `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration  `json:"totalAbsoluteTime"`
	Percentage        int            `json:"percentage"`
	IsComplete        bool           `json:"isComplete"`
}

type SummaryWeek struct {
	TimestampsWeek    []db.Timestamp `json:"timestamps"`
	TotalFloat        float32        `json:"differenceFloat"`
	TotalAbsoluteTime string         `json:"totalAbsoluteTime"`
	TotalWorkingDays  int            `json:"totalWorkingDays"`
}
