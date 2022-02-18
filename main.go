package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"

	util "github.com/schwarztrinker/trkbox/util"

	"github.com/gorilla/mux"
)

var timestampsGlobal util.Timestamps

func main() {
	r := mux.NewRouter()

	// Loading Timestamp DB
	loadingTimestampsGlobalFromDB()

	// Check in and out
	r.HandleFunc("/stamp", stampHandler).Methods("POST")

	// Get summary for a specific day
	r.HandleFunc("/summary/day/{day}", getSummaryForDay).Methods("GET")

	// Get summary for a specific day
	r.HandleFunc("/summary/week/{week}", func(rw http.ResponseWriter, r *http.Request) {}).Methods("GET")

	// get list of all timestampsGlobal
	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timestampsGlobal)
	}).Methods("GET")

	// Handling a test ping to the server
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("pong")
	})

	http.ListenAndServe(":13370", r)
}

func stampHandler(w http.ResponseWriter, r *http.Request) {
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

func getSummaryForDay(w http.ResponseWriter, r *http.Request) {
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

	summary.TotalAbsoluteTime = calculateTotalPresenceDuration(out)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func calculateTotalPresenceDuration(ts util.Timestamps) time.Duration {
	var absoluteTime time.Duration = 0

	if len(ts.Timestamps) > 1 {
		if len(ts.Timestamps)%2 == 0 {

		}

		for i, _ := range ts.Timestamps {

			if ts.Timestamps[i].IsCheckin && !ts.Timestamps[i-1].IsCheckin {
				absoluteTime += ts.Timestamps[i].Date.Sub(ts.Timestamps[i-1].Date)
			}

		}
	}
	return absoluteTime
}

func calculateIsComplete(timestamps util.Timestamps) bool {
	if len(timestamps.Timestamps)%2 == 0 && len(timestamps.Timestamps) > 0 && checkinIsAlternating(timestamps) {
		return false
	}
	return true
}

func checkinIsAlternating(timestamps util.Timestamps) bool {
	var last bool
	for i := range timestamps.Timestamps {
		if last == timestamps.Timestamps[i].IsCheckin {
			return false
		}

		last = timestamps.Timestamps[i].IsCheckin
	}
	return true
}

func loadingTimestampsGlobalFromDB() {
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
