package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/schwarztrinker/trkbox/util"
	"golang.org/x/crypto/bcrypt"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//Implement Authentication Logic and User DB
	var userRequested util.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&userRequested)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser util.User
	if checkPasswordHash(userRequested.Password, dbUser.PasswordHash) {

	}

	// Genereate Token and send it to client
	validToken, err := GenerateJWT(dbUser)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(validToken)

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authorization Middleware
func Authorized(w http.ResponseWriter, r *http.Request, handler func(w http.ResponseWriter, r *http.Request)) {

	if r.Header["Authorization"] != nil && len(strings.Fields(r.Header["Authorization"][0])) > 1 {
		potentialToken := r.Header["Authorization"][0]
		fmt.Print(potentialToken)
		split := strings.Split(potentialToken, " ")

		token, err := jwt.Parse(split[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Printf(err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		if token.Valid {
			handler(w, r)
		}
	} else {

		fmt.Fprintf(w, "Not Authorized or token wrong")
	}

}

// Generate a token for the User
func GenerateJWT(user util.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
