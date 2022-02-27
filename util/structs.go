package util

import (
	"time"

	"github.com/schwarztrinker/trkbox/db"
)

type SummaryToday struct {
	TimestampsToday   db.Timestamps `json:"timestamps"`
	DifferenceFloat   float32       `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration `json:"totalAbsoluteTime"`
	Percentage        int           `json:"percentage"`
	IsComplete        bool          `json:"isComplete"`
}

type SummaryWeek struct {
	TimestampsWeek    db.Timestamps `json:"timestamps"`
	TotalFloat        float32       `json:"differenceFloat"`
	TotalAbsoluteTime string        `json:"totalAbsoluteTime"`
	TotalWorkingDays  int           `json:"totalWorkingDays"`
}
