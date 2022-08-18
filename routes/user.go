package routes

import (
	"github.com/dilanEspindola/restapiFiber/database"
	"github.com/dilanEspindola/restapiFiber/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Message  string `json:"message"`
}

type MessageUserDeleted struct {
	Message string `json:"message"`
}

func responseUserCreate(userModel models.User) User {
	return User{
		Id:       int(userModel.Id),
		Name:     userModel.Name,
		LastName: userModel.LastName,
		Message:  "user created",
	}
}

func responseErrorOrNotFound(msg string) MessageUserDeleted {
	return MessageUserDeleted{
		Message: msg,
	}
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	if err := database.Database.Db.Find(&users); err.Error != nil {
		message := responseErrorOrNotFound("internal serve error")
		return c.Status(500).JSON(message)
	}

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	user := models.User{}

	if err := database.Database.Db.Find(&user, paramsId); err.Error != nil {
		message := responseErrorOrNotFound("internal serve error")
		return c.Status(500).JSON(message)
	}

	if user.Id == 0 {
		message := responseErrorOrNotFound("user not found")
		return c.Status(404).JSON(message)
	}

	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Create(&user); err.Error != nil {
		message := responseErrorOrNotFound("internal serve error")
		return c.Status(500).JSON(message)
	}
	responseUser := responseUserCreate(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	user := models.User{}

	if err := database.Database.Db.Find(&user, paramsId); err.Error != nil {
		message := responseErrorOrNotFound("internal serve error")
		return c.Status(500).JSON(message)
	}

	if user.Id > 0 {
		if err := database.Database.Db.Where("id = ?", paramsId).Delete(&user); err.Error != nil {
			message := responseErrorOrNotFound("internal serve error")
			return c.Status(500).JSON(message)
		}

		message := responseErrorOrNotFound("user deleted")

		return c.Status(200).JSON(message)
	}

	message := responseErrorOrNotFound("user not found")

	return c.Status(200).JSON(message)
}
