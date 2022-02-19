package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/schwarztrinker/trkbox/util"
)

var timestampsGlobal util.Timestamps

func HomePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")

}

func StampHandler(w http.ResponseWriter, r *http.Request) {
	var t util.Timestamp

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return Timestamp to successfull confirm checkin
	w.Header().Set("Content-Type", "application/json")
	timestampsGlobal.Timestamps = append(timestampsGlobal.Timestamps, t)
	savingTimestampsGlobalFromDB()
	json.NewEncoder(w).Encode(&t)
}

func DeleteStampHandler(w http.ResponseWriter, r *http.Request) {
	arg := mux.Vars(r)
	stampId, err := strconv.Atoi(arg["id"])
	if err != nil {
		log.Fatal(err)
	}

	timestampsGlobal.Timestamps[stampId] = timestampsGlobal.Timestamps[len(timestampsGlobal.Timestamps)-1] // Copy last element to index i.
	//timestampsGlobal.Timestamps[len(timestampsGlobal.Timestamps)-1] = ""   // Erase last element (write zero value).
	timestampsGlobal.Timestamps = timestampsGlobal.Timestamps[:len(timestampsGlobal.Timestamps)-1] // Truncate slice.
	savingTimestampsGlobalFromDB()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("YES")
}

func GetSummaryForDay(w http.ResponseWriter, r *http.Request) {
	// getting router arg from mux
	arg := mux.Vars(r)
	var summary util.SummaryToday

	// Get all timestampsGlobal from the Day
	var out util.Timestamps
	for _, v := range timestampsGlobal.Timestamps {
		if v.Date.Format("2006-01-02") == arg["day"] {
			out.Timestamps = append(out.Timestamps, v)
		}
	}
	summary.TimestampsToday = out

	// Todo Calculate working hours
	// check if timestampsGlobal are correct
	summary.IsComplete = calculateIsComplete(out)
	summary.TotalAbsoluteTime = calculateTotalPresenceDuration(out).Round(time.Second)
	summary.DifferenceFloat = float32(summary.TotalAbsoluteTime.Hours())
	summary.Percentage = int((summary.DifferenceFloat / 8) * 100)

	fmt.Print(summary)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func ListAllTimestamps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timestampsGlobal)

}

func calculateTotalPresenceDuration(ts util.Timestamps) time.Duration {
	var absoluteTime time.Duration

	if len(ts.Timestamps) > 1 {

		for i, _ := range ts.Timestamps {
			if i == 0 {
				continue
			}

			if ts.Timestamps[i].IsCheckin != ts.Timestamps[i-1].IsCheckin && ts.Timestamps[i-1].IsCheckin {
				absoluteTime += ts.Timestamps[i].Date.Sub(ts.Timestamps[i-1].Date)
			}

		}
	}

	if len(ts.Timestamps) == 1 {
		if ts.Timestamps[0].IsCheckin {
			absoluteTime = time.Now().Sub(ts.Timestamps[0].Date)
		}
	}
	return absoluteTime
}

func calculateIsComplete(ts util.Timestamps) bool {
	if len(ts.Timestamps)%2 == 1 || len(ts.Timestamps) <= 1 || !checkinIsAlternating(ts) {
		return false
	}
	return true
}

func checkinIsAlternating(ts util.Timestamps) bool {
	var last bool
	for i := range ts.Timestamps {
		if last == ts.Timestamps[i].IsCheckin {
			return false
		}

		last = ts.Timestamps[i].IsCheckin
	}
	return true
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
	timestampsGlobal.Timestamps = append(timestampsGlobal.Timestamps, timestampsGlobalStruct.Timestamps...)
}

func savingTimestampsGlobalFromDB() {

	//Sort all timestamps by Date before saving
	sort.Slice(timestampsGlobal.Timestamps, func(i, j int) bool {
		return timestampsGlobal.Timestamps[i].Date.Before(timestampsGlobal.Timestamps[j].Date)
	})

	for i := range timestampsGlobal.Timestamps {
		timestampsGlobal.Timestamps[i].Id = i
	}

	//save all the timestamps
	file, _ := json.MarshalIndent(timestampsGlobal, "", " ")

	_ = ioutil.WriteFile("db.json", file, 0644)
}
