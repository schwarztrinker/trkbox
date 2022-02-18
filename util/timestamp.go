package util

import "time"

// TimeStamp struct
type Timestamps struct {
	Timestamps []Timestamp `json:"timestamps"`
}

// TimeStamp struct
type Timestamp struct {
	Date      time.Time `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
}

type SummaryToday struct {
	Timestamps        Timestamps    `json:"timestamps"`
	DifferenceFloat   float32       `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration `json:"totalAbsoluteTime"`
}

type SummaryWeek struct {
	Timestamps        Timestamps `json:"timestamps"`
	TotalFloat        float32    `json:"differenceFloat"`
	TotalAbsoluteTime string     `json:"totalAbsoluteTime"`
	TotalWorkingDays  int        `json:"totalWorkingDays"`
}
