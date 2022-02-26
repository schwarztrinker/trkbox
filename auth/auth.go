package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/db"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(u string) (*db.User, error) {
	database := db.UsersDB.Users

	for _, i := range database {
		if i.Username == u {
			return &i, nil
		}
	}
	return nil, errors.New("No user found")
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome " + name + "!")
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	identity := input.Identity
	password := input.Password

	user, err := GetUserByUsername(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Wrong username or password", "data": err})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(conf.Conf.JwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func CreateUser(c *fiber.Ctx) error {
	usr := new(db.User)

	if err := c.BodyParser(&usr); err != nil {
		return err
	}

	usr, err := GetUserByUsername(usr.Username)
	if usr != nil && err == nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"status": "error", "message": "Username already exists"})
	}

	db.UsersDB.CreateNewUser(*usr)

	return c.JSON(usr)
}
