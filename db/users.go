package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username"`
	PasswordHash []byte
	Timestamps   []Timestamp `json:"timestamps"`
}

func UserExists(u string) bool {
	var user []User

	maria.Where("username = ?", u).Find(&user)

	if len(user) > 0 {
		return true
	}

	return false
}

func GetUserByUsername(u string) (*User, error) {
	var user *User

	maria.Where("username = ?", u).First(&user)
	if user != nil {
		return user, nil
	}

	return nil, errors.New("No user found")
}

func CreateNewUser(username string, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	// Generate new user and insert data
	user := User{Username: username, PasswordHash: hashedPassword}

	result := maria.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		panic(result.Error)
	}

}
