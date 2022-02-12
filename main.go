package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// TimeStamp struct
type Timestamp struct {
	Time string `json:"time"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func main() {
	r := mux.NewRouter()

	var timestamps []Timestamp

	r.HandleFunc("/checkin", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		stamp := generateCurrentTimestamp("checkin")
		timestamps = append(timestamps, stamp)
		json.NewEncoder(w).Encode(stamp)

	})

	r.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		stamp := generateCurrentTimestamp("checkout")
		timestamps = append(timestamps, stamp)
		json.NewEncoder(w).Encode(stamp)
	})

	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timestamps)
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("pong")
	})

	http.ListenAndServe(":13370", r)
}

func getAllTimestamps() {
	panic("unimplemented")
}

// Returns a time string
//e.g. {12:18:12}
func currentTime() string {
	currTime := time.Now()
	hours, minutes, seconds := currTime.Clock()
	out := formatDigitTwoLetters(hours) + ":" + formatDigitTwoLetters(minutes) + ":" + formatDigitTwoLetters(seconds)
	return out
}

//Formats single digits for time and dates to two digit strings
//@d int
//@out string (two digits)
func formatDigitTwoLetters(d int) string {
	var out string
	if d < 10 {
		out = "0" + strconv.Itoa(d)
	} else {
		out = strconv.Itoa(d)
	}
	return out
}

//Generates the current Timestamp struct from time and date strings
func generateCurrentTimestamp(t string) Timestamp {
	year, month, day := time.Now().Date()
	date := strconv.Itoa(year) + "-" + formatDigitTwoLetters(int(month)) + "-" + formatDigitTwoLetters(day)

	stamp := Timestamp{Time: currentTime(), Date: date}
	stamp.Type = t
	return stamp
}
