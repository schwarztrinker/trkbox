package main

import (
	"encoding/json"
	"net/http"
	"time"

	util "github.com/schwarztrinker/trkbox/util"

	"github.com/gorilla/mux"
)

var timestamps []util.Timestamp

func main() {
	r := mux.NewRouter()

	// Sample Dates for Testing
	testDay := time.Now().AddDate(0, 0, -1)
	timestamps = append(timestamps, util.Timestamp{Date: testDay, IsCheckin: true}, util.Timestamp{Date: testDay, IsCheckin: false})

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
	timestamps = append(timestamps, t)
	json.NewEncoder(w).Encode(&t)
}

func getSummaryForDay(w http.ResponseWriter, r *http.Request) {
	// getting router arg from mux
	arg := mux.Vars(r)
	var summary util.SummaryToday

	// Get all Timestamps from the Day
	var out []util.Timestamp
	for _, v := range timestamps {
		if v.Date.Format("2006-01-02") == arg["day"] {
			out = append(out, v)
		}
	}
	summary.Timestamps = out

	// Todo Calculate working hours

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
