package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/schwarztrinker/trkbox/util"
)

var TimestampsDB util.Timestamps

func AddTimestampToDB(ts util.Timestamp) {
	TimestampsDB.Timestamps = append(TimestampsDB.Timestamps, ts)
	savingTimestampsGlobalFromDB()
}

func DeleteTimestampByID(id int) {
	TimestampsDB.Timestamps[id] = TimestampsDB.Timestamps[len(TimestampsDB.Timestamps)-1] // Copy last element to index i.
	//timestampsGlobal.Timestamps[len(timestampsGlobal.Timestamps)-1] = ""   // Erase last element (write zero value).
	TimestampsDB.Timestamps = TimestampsDB.Timestamps[:len(TimestampsDB.Timestamps)-1] // Truncate slice.
	savingTimestampsGlobalFromDB()
}

func LoadingTimestampsGlobalFromDB() {
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var timestampsGlobalStruct util.Timestamps

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &timestampsGlobalStruct)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	TimestampsDB.Timestamps = append(TimestampsDB.Timestamps, timestampsGlobalStruct.Timestamps...)
}

func savingTimestampsGlobalFromDB() {

	//Sort all timestamps by Date before saving
	sort.Slice(TimestampsDB.Timestamps, func(i, j int) bool {
		return TimestampsDB.Timestamps[i].Date.Before(TimestampsDB.Timestamps[j].Date)
	})

	for i := range TimestampsDB.Timestamps {
		TimestampsDB.Timestamps[i].Id = i
	}

	//save all the timestamps
	file, _ := json.MarshalIndent(TimestampsDB, "", " ")

	_ = ioutil.WriteFile("db.json", file, 0644)
}
