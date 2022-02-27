package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TimeStamp struct
type Timestamps struct {
	Timestamps []Timestamp `json:"timestamps"`
}

// TimeStamp struct
type Timestamp struct {
	Uuid      uuid.UUID `json:"uuid"`
	Date      time.Time `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
}

func (t *Timestamps) AppendTimestamp(ts Timestamp) *Timestamp {
	ts.Uuid = uuid.New()
	t.Timestamps = append(t.Timestamps, ts)

	UsersDB.SaveDB()
	return &ts
}

func (t *Timestamps) DeleteTimestampByUuid(uuid uuid.UUID) (*Timestamp, error) {
	ts, index, err := t.GetTimestampAndIndexByUUID(uuid)
	if err != nil {
		return nil, err
	}

	t.Timestamps[index] = t.Timestamps[len(t.Timestamps)-1]

	// Copy last element to index i.
	//timestampsGlobal.Timestamps[len(timestampsGlobal.Timestamps)-1] = ""   // Erase last element (write zero value).
	t.Timestamps = t.Timestamps[:len(t.Timestamps)-1] // Truncate slice.
	UsersDB.SaveDB()

	return ts, nil
}

// Getter
func (t *Timestamps) GetTimestampAndIndexByUUID(uuid uuid.UUID) (*Timestamp, int, error) {
	for index, ts := range t.Timestamps {
		if ts.Uuid == uuid {
			return &t.Timestamps[index], index, nil
		}
	}
	return nil, 0, errors.New("No timestamp Found")
}

func (t *Timestamps) GetTimestampsByDay(day string) (*Timestamps, error) {
	// Get all timestampsGlobal from the Day
	var out Timestamps
	for _, v := range t.Timestamps {
		if v.Date.Format("2006-01-02") == day {
			out.Timestamps = append(out.Timestamps, v)
		}
	}

	if len(out.Timestamps) == 0 {
		return nil, errors.New("No Timestamps found")
	}

	return &out, nil
}
