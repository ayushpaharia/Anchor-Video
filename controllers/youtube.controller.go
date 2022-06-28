package controllers

import (
	"context"
	"fampay-youtube/config"
	"fampay-youtube/models"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"title": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	if sort := c.Query("sort"); sort != "" {
		if sort == "asc" {
			findOptions.SetSort(bson.D{{Key: "publishedAt", Value: 1}})
		} else if sort == "dsc" {
			findOptions.SetSort(bson.D{{Key: "publishedAt", Value: -1}})
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	var perPage int64 = 5

	total, _ := videosCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)

	cursor, err := videosCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "",
			"success": false,
			"error":   err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var videos models.Videos = make(models.Videos, 0)

	for cursor.Next(ctx) {
		var video models.Video
		cursor.Decode(&video)
		videos = append(videos, video)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      videos,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total) / float64(perPage)),
		"success":   true,
	})
}
