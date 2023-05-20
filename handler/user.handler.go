package handler

import (
	"log"
	"tayoh-backend/database"
	"tayoh-backend/model/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func Create(ctx *fiber.Ctx) error {

}