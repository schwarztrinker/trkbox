package timestamp

import "time"

// TimeStamp struct
type Timestamp struct {
	Date time.Time `json:"date"`
	Type string    `json:"type"`
}
