package db

import (
	"encoding/json"
	"io/ioutil"

	"github.com/schwarztrinker/trkbox/util"
)

var usersDB util.Users

func CreateNewUser(u util.User) {
	usersDB.Users = append(usersDB.Users, u)
	saveUserDB()
}

func saveUserDB() {
	//Sort all timestamps by Date before saving

	//save all the timestamps
	file, _ := json.MarshalIndent(usersDB, "", " ")

	_ = ioutil.WriteFile("db.json", file, 0644)
}
