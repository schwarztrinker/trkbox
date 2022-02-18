package timestamp

import "time"

// TimeStamp struct
type Timestamp struct {
	Date      time.Time `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
}

type SummaryToday struct {
	Timestamps        []Timestamp   `json:"timestamps"`
	DifferenceFloat   float32       `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration `json:"totalAbsoluteTime"`
}

type SummaryWeek struct {
	Timestamps        []Timestamp `json:"timestamps"`
	TotalFloat        float32     `json:"differenceFloat"`
	TotalAbsoluteTime string      `json:"totalAbsoluteTime"`
	TotalWorkingDays  int         `json:"totalWorkingDays"`
}
