package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Timestamp struct {
	gorm.Model
	Uuid      uuid.UUID `json:"uuid"`
	Time      time.Time `json:"time"`
	Date      string    `json:"date"`
	IsCheckin bool      `json:"isCheckin"`
	UserID    int
}

func DeleteTimestampByUuid(user User, uuid uuid.UUID) (*Timestamp, error) {
	var ts *Timestamp

	result := maria.Where("uuid = ? AND user_id = ?", uuid, user.ID).Delete(&ts)
	if result.Error != nil {
		return ts, result.Error
	}

	return ts, nil
}

// Getter
func GetTimestampAndIndexByUUID(uuid uuid.UUID) (*Timestamp, int, error) {
	var ts *Timestamp

	maria.Where("uuid = ?", uuid).First(&ts)
	if ts != nil {
		return ts, int(ts.ID), nil
	}

	return nil, 0, errors.New("No timestamp Found")
}

func GetTimestampsByDay(user User, date string) ([]Timestamp, error) {
	var ts []Timestamp
	maria.Where("date = ? AND user_id = ? ", date, user.ID).Find(&ts)

	return ts, nil
}

func GetTimestampsFromUser(user User) []Timestamp {
	var ts []Timestamp
	maria.Where("user_id= ?", user.ID).Find(&ts)

	return ts
}

func AddTimestamp(user User, ts Timestamp) (*Timestamp, error) {

	timestamp := Timestamp{Uuid: uuid.New(), UserID: int(user.ID), Time: ts.Time, Date: ts.Time.Format("2006-01-02"), IsCheckin: ts.IsCheckin}

	result := maria.Create(&timestamp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &timestamp, nil

}
