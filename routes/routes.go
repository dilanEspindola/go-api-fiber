package routes

import "github.com/gofiber/fiber/v2"

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  byte   `json:"age"`
}

type JsonMessage struct {
	Msg string `json:"msg"`
}

var users []User

func Users() []User {
	users = append(users, User{
		Id:   "1",
		Name: "dilan",
		Age:  20,
	})
	users = append(users, User{
		Id:   "2",
		Name: "Andrea",
		Age:  18,
	})
	return users
}

// SendJSON
func sendJson(message string, status int) JsonMessage {
	newMessage := JsonMessage{
		Msg: message,
	}
	return newMessage
}

// GetUsers
func GetUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(Users())
}

// GetUser
func GetUser(c *fiber.Ctx) error {
	idParams := c.Params("id")
	for _, user := range Users() {
		if user.Id == idParams {
			return c.Status(200).JSON(user)
		}
	}
	msg := sendJson("user not found", 404)

	return c.Status(404).JSON(msg)
}
