package util

import "time"

// TimeStamp struct
type Timestamps struct {
	Timestamps []Timestamp `json:"timestamps"`
}

// TimeStamp struct
type Timestamp struct {
	Id        int       `json:"id"`
	Date      time.Time `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
}

type SummaryToday struct {
	TimestampsToday   Timestamps    `json:"timestamps"`
	DifferenceFloat   float32       `json:"differenceFloat"`
	TotalAbsoluteTime time.Duration `json:"totalAbsoluteTime"`
	Percentage        int           `json:"percentage"`
	IsComplete        bool          `json:"isComplete"`
}

type SummaryWeek struct {
	TimestampsWeek    Timestamps `json:"timestamps"`
	TotalFloat        float32    `json:"differenceFloat"`
	TotalAbsoluteTime string     `json:"totalAbsoluteTime"`
	TotalWorkingDays  int        `json:"totalWorkingDays"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username     string     `json:"username"`
	PasswordHash string     `json:"passwordHash"`
	Password     string     `json:"password"`
	Salt         string     `json:"salt"`
	Timestamps   Timestamps `json:"timestamps"`
}
