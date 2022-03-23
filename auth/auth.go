package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/schwarztrinker/trkbox/conf"
	"github.com/schwarztrinker/trkbox/db"
	"golang.org/x/crypto/bcrypt"
)

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

	user, err := db.GetUserByUsername(identity)
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
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Token generated for user", "data": t})
}

func CreateUser(c *fiber.Ctx) error {
	type UserInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var userInput UserInput

	if err := c.BodyParser(&userInput); err != nil {
		return err
	}

	if db.UserExists(userInput.Username) {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"status": "error", "message": "Username already exists"})
	}

	db.CreateNewUser(userInput.Username, userInput.Password)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully created"})
}

func CheckUserFromToken(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	userObj, err := db.GetUserByUsername(name)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Incorrect User in Token", "data": err})
	}

	c.Locals("username", userObj.Username)
	return c.Next()
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome " + name + "!")
}

func Whoami(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.JSON(fiber.Map{"status": "success", "message": "You are logged in as user", "data": name})
}
