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

// // func DeleteTimestampByID(id int) {

// // }

// func LoadingTimestampsGlobalFromDB() {
// 	jsonFile, err := os.Open("db.json")
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var timestampsGlobalStruct util.Timestamps

// 	byteValue, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = json.Unmarshal(byteValue, &timestampsGlobalStruct)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// defer the closing of our jsonFile so that we can parse it later on
// 	defer jsonFile.Close()
// 	TimestampsDB.Timestamps = append(TimestampsDB.Timestamps, timestampsGlobalStruct.Timestamps...)
// }

// func savingTimestampsGlobalFromDB() {

// 	//Sort all timestamps by Date before saving
// 	sort.Slice(TimestampsDB.Timestamps, func(i, j int) bool {
// 		return TimestampsDB.Timestamps[i].Date.Before(TimestampsDB.Timestamps[j].Date)
// 	})

// 	for i := range TimestampsDB.Timestamps {
// 		TimestampsDB.Timestamps[i].Id = i
// 	}

// 	//save all the timestamps
// 	file, _ := json.MarshalIndent(TimestampsDB, "", " ")

// 	_ = ioutil.WriteFile("db.json", file, 0644)
// }
