package handler

import (
	"log"
	"tayoh-backend/database"
	"tayoh-backend/model/entity"
	"tayoh-backend/model/request"

	"github.com/gofiber/fiber/v2"
)

func UserGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User {
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "Success to store user data",
		"data": newUser,
	})
}