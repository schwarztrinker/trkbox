package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/schwarztrinker/trkbox/util"
	"golang.org/x/crypto/bcrypt"
)

var UsersDB Users

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username     string          `json:"username"`
	PasswordHash []byte          `json:"passwordHash"`
	Password     string          `json:"password"`
	Timestamps   util.Timestamps `json:"timestamps"`
}

func (u *Users) CreateNewUser(user User) *Users {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}

	user.PasswordHash = hashedPassword
	//remove password before saving
	user.Password = ""
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it

	UsersDB.Users = append(UsersDB.Users, user)
	u.saveUserDB()
	return u
}

func (u *Users) saveUserDB() *Users {

	file, _ := json.MarshalIndent(u, "", " ")

	_ = ioutil.WriteFile("users.json", file, 0644)
	return u
}

func (u *Users) LoadUserDB() *Users {
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var users Users

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	UsersDB.Users = append(UsersDB.Users, users.Users...)
	return u
}
