package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var UsersDB Users

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username     string     `json:"username"`
	PasswordHash []byte     `json:"passwordHash"`
	Timestamps   Timestamps `json:"timestamps"`
}

func (u *Users) CreateNewUser(username string, password string) *Users {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	// Generate new user and insert data
	user := new(User)
	user.Init(username, hashedPassword)

	UsersDB.Users = append(UsersDB.Users, *user)
	u.SaveDB()
	return u
}

func (u *Users) SaveDB() *Users {

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

func (u *User) Init(username string, passwordHash []byte) {
	u.Username = username
	u.PasswordHash = passwordHash
}
