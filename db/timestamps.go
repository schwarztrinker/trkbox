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
	User      User
	UserID    int
}

func DeleteTimestampByUuid(uuid uuid.UUID) (*Timestamp, error) {
	var ts *Timestamp

	maria.Where("uuid = ?", uuid).Delete(&ts)
	if ts != nil {
		return ts, nil
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

func GetTimestampsByDay(user User, day string) ([]Timestamp, error) {
	var ts []Timestamp
	maria.Where("date = ? AND user = ? ", day, user).Find(&ts)

	return ts, nil
}

func GetTimestampsFromUser(user User) []Timestamp {
	var ts []Timestamp
	maria.Where("user_id= ?", user.ID).Find(&ts)

	return ts
}

func AddTimestamp(user User, ts Timestamp) *Timestamp {
	timestamp := Timestamp{Uuid: uuid.New(), User: user, Time: ts.Time}

	result := maria.Create(&timestamp) // pass pointer of data to Create
	if result.Error != nil {
		panic(result.Error)
	}

	return &ts

}
