package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	util "github.com/schwarztrinker/trkbox/util"

	"github.com/gorilla/mux"
)

var timestamps util.Timestamps

func main() {
	r := mux.NewRouter()

	// Sample Dates for Testing
	//testDay := time.Now().AddDate(0, 0, -1)
	//timestamps = append(timestamps, util.Timestamp{Date: testDay, IsCheckin: true}, util.Timestamp{Date: testDay, IsCheckin: false})
	timestamps.Timestamps = append(timestamps.Timestamps, readTimestampsFromDB().Timestamps...)
	// Check in and out
	r.HandleFunc("/stamp", stampHandler).Methods("POST")

	// Get summary for a specific day
	r.HandleFunc("/summary/day/{day}", getSummaryForDay).Methods("GET")

	// Get summary for a specific day
	r.HandleFunc("/summary/week/{week}", func(rw http.ResponseWriter, r *http.Request) {}).Methods("GET")

	// get list of all timestamps
	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timestamps)
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
	var t util.Timestamps

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return Timestamp to successfull confirm checkin
	w.Header().Set("Content-Type", "application/json")
	timestamps.Timestamps = append(timestamps.Timestamps, t.Timestamps...)
	json.NewEncoder(w).Encode(&t)
}

func getSummaryForDay(w http.ResponseWriter, r *http.Request) {
	// getting router arg from mux
	arg := mux.Vars(r)
	var summary util.SummaryToday

	// Get all Timestamps from the Day
	var out util.Timestamps
	for _, v := range timestamps.Timestamps {
		if v.Date.Format("2006-01-02") == arg["day"] {
			out.Timestamps = append(out.Timestamps, v)
		}
	}
	summary.Timestamps = out

	// Todo Calculate working hours
	// check if timestamps are correct
	var absoluteTime time.Duration
	if len(out.Timestamps)%2 == 0 && len(out.Timestamps) > 0 && calculationPossible(summary.Timestamps.Timestamps) {

		for i := range out.Timestamps {
			if i%2 == 1 {
				absoluteTime += out.Timestamps[i].Date.Sub(out.Timestamps[i-1].Date)
			}

		}
	}
	summary.TotalAbsoluteTime = absoluteTime

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func calculationPossible(timestamps []util.Timestamp) bool {

	return true
}

func readTimestampsFromDB() util.Timestamps {
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// we initialize our Users array
	var timestamps util.Timestamps

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &timestamps)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timestamps)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return timestamps

}
