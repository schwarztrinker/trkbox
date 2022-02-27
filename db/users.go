package db

import (
	"encoding/json"
	"errors"
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

func (p *Users) GetUserByUsername(u string) (*User, error) {
	for index, i := range UsersDB.Users {
		if i.Username == u {
			return &UsersDB.Users[index], nil
		}
	}
	return nil, errors.New("No user found")
}

func (u *Users) AddTimestampToUser(username string, ts Timestamp) *Users {
	for _, i := range UsersDB.Users {
		if i.Username == username {
			i.Timestamps.Timestamps = append(i.Timestamps.Timestamps, ts)
		}
	}

	UsersDB.SaveDB()

	return u
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

	UsersDB = users
	return u
}

func (u *User) Init(username string, passwordHash []byte) {
	u.Username = username
	u.PasswordHash = passwordHash
}
