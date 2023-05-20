package route

import (
	"tayoh-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/api/users", handler.UserGetAll)
	r.Post("/api/user", handler.CreateUser)
}