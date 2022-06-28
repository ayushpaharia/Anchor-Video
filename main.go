package main

import (
	"fampay-youtube/config"
	"fampay-youtube/routes"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\nShutting down...")
		app.Shutdown()
	}()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	setupRoutes(app)

	// cron.StartYoutubeFetch()

	port := string(os.Getenv("PORT"))
	if err := app.Listen(":" + port); err != nil {
		log.Panic("Error app failed to start ", err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Fampay YouTube",
		})
	})

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "https://github.com/ayushpaharia/fampay-youtube",
		})
	})
	routes.YoutubeRoutes(api.Group("/youtube"))
}
