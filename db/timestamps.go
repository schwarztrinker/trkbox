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
	UserID    uint
	Category  string `json:"category"`
}

func GetTimestampAndIndexByUUID(uuid uuid.UUID) (*Timestamp, int, error) {
	var ts *Timestamp

	maria.Where("uuid = ?", uuid).First(&ts)
	if ts != nil {
		return ts, int(ts.ID), nil
	}

	return nil, 0, errors.New("No timestamp Found")
}

// Getting Timestamps
func GetTimestampsFromUser(user User) []Timestamp {
	var ts []Timestamp
	maria.Where("user_id= ?", user.ID).Order("time asc").Find(&ts)

	return ts
}

func GetTimestampsByDay(user User, date string) ([]Timestamp, error) {
	var ts []Timestamp
	maria.Where("date = ? AND user_id = ? ", date, user.ID).Order("time asc").Find(&ts)

	return ts, nil
}

//Create Timestamps
func AddTimestamp(user User, ts Timestamp) (*Timestamp, error) {

	timestamp := ts
	timestamp.Uuid = uuid.New()
	timestamp.UserID = user.ID
	timestamp.Date = ts.Time.Format("2006-01-02")

	result := maria.Create(&timestamp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &timestamp, nil
}

// Delete Timestamps
func DeleteTimestampByUuid(user User, uuid uuid.UUID) (*Timestamp, error) {
	var ts *Timestamp

	result := maria.Where("uuid = ? AND user_id = ?", uuid, user.ID).Delete(&ts)
	if result.Error != nil {
		return ts, result.Error
	}

	return ts, nil
}

// Delete Timestamps
func DeleteTimestampById(user User, id uint) (*Timestamp, error) {
	var ts *Timestamp

	result := maria.Where("id = ? AND user_id = ?", id, user.ID).Delete(&ts)
	if result.RowsAffected == 0 {
		return ts, errors.New("No timestamp found or deleted")
	}

	return ts, nil
}
