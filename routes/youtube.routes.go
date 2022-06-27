package routes

import (
	"fampay-youtube/controllers"

	"github.com/gofiber/fiber/v2"
)

func YoutubeRoutes(route fiber.Router) {
	route.Get("/", controllers.Something)
}
