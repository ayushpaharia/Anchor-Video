package controllers

import (
	"context"
	"fampay-youtube/config"
	"fampay-youtube/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Something(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "Hello",
		"success": true,
	})
}

func GetVideosPaginated(c *fiber.Ctx) error {
	// page := c.Query("page")
	// limit := c.Query("limit")
	videosCollection := config.MI.DB.Collection("videos")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{}}

	cursor, err := videosCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "",
			"success": false,
			"error":   err.Error(),
		})
	}

	var videos models.Videos = make(models.Videos, 0)

	err = cursor.All(ctx, &videos)
	defer cursor.Close(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    videos,
		"success": true,
	})
}
