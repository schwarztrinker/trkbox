package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schwarztrinker/trkbox/auth"
	"github.com/schwarztrinker/trkbox/handlers"
)

func Router() {
	r := mux.NewRouter()

	// Loading Timestamp DB
	handlers.LoadingTimestampsGlobalFromDB()

	// Check in and out
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")

	// Check in and out
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		auth.Authorized(w, r, handlers.HomePage)
	}).Methods("GET")

	// Check in and out
	r.HandleFunc("/stamp", handlers.StampHandler).Methods("POST")

	// Check in and out
	r.HandleFunc("/stamp/delete/{id}", handlers.DeleteStampHandler).Methods("POST")

	// Get summary for a specific day
	r.HandleFunc("/summary/day/{day}", handlers.GetSummaryForDay).Methods("GET")

	// Get summary for a specific day
	r.HandleFunc("/summary/week/{week}", func(rw http.ResponseWriter, r *http.Request) {}).Methods("GET")

	// get list of all timestampsGlobal
	r.HandleFunc("/list", handlers.ListAllTimestamps).Methods("GET")

	// Handling a test ping to the server
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// TODO save logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("pong")
	})

	http.ListenAndServe("0.0.0.0:13370", r)

}
