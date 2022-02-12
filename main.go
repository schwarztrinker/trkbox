package main

import (
    "fmt"
    "net/http"
    "time"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/checkin", func(w http.ResponseWriter, r *http.Request) {
        year, month, day := time.Now().Date()
        date := strconv.Itoa(year)+"-"+ strconv.Itoa(int(month))+"-"+ strconv.Itoa(day)

        stamp := timestamp {time: currentTime(), date: date} 
        //fmt.Fprintf(w, "You've checked in at %s", stamp)

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        marsh, _ := json.Marshal(stamp)
        w.Write(marsh)
    })

    r.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "You've checked out at %s", currentTime())
    })

    
    http.ListenAndServe(":80", r)
}

func currentTime() string{
    currTime := time.Now()
    hours, minutes, seconds :=  currTime.Clock()
    out := strconv.Itoa(hours)+":"+ strconv.Itoa(minutes)+":"+ strconv.Itoa(seconds)
    return out
}

type timestamp struct {
    time string
    date string
}