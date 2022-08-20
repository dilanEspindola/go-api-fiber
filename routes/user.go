package routes

import (
	"strconv"

	"github.com/dilanEspindola/restapiFiber/database"
	"github.com/dilanEspindola/restapiFiber/interfaces"
	"github.com/dilanEspindola/restapiFiber/models"
	"github.com/gofiber/fiber/v2"
)

func responseUserCreate(userModel models.User) interfaces.UserResonseCreate {
	return interfaces.UserResonseCreate{
		Id:       int(userModel.Id),
		Name:     userModel.Name,
		LastName: userModel.LastName,
		Message:  "user created",
	}
}

func responseErrorOrNotFound(msg string) interfaces.MessageUserDeleted {
	return interfaces.MessageUserDeleted{
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

func UpdateUser(c *fiber.Ctx) error {
	var newUser interfaces.UserEdit
	findUser := models.User{}
	paramsId := c.Params("id")
	strconv.ParseUint(paramsId, 0, 64)

	if err := c.BodyParser(&newUser); err != nil {
		response := responseErrorOrNotFound("invalid data")
		return c.Status(400).JSON(response)
	}

	if newUser.Name == "" && newUser.LastName == "" {
		response := responseErrorOrNotFound("you must be send at least one data")
		return c.Status(400).JSON(response)
	}

	if err := database.Database.Db.Find(&findUser, paramsId); err.Error != nil {
		response := responseErrorOrNotFound("internal server error")
		return c.Status(500).JSON(response)
	}

	if newUser.Name == "" {
		newUser.Name = findUser.Name
	}

	if newUser.LastName == "" {
		newUser.LastName = findUser.LastName
	}

	findUser.Name = newUser.Name
	findUser.LastName = newUser.LastName

	if err := database.Database.Db.Save(&findUser); err.Error != nil {
		response := responseErrorOrNotFound("internal server error")
		return c.Status(500).JSON(response)
	}

	return c.Status(200).JSON(findUser)
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
