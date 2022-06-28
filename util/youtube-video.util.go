package util

import (
	"context"
	"fampay-youtube/config"
	"fampay-youtube/models"
	"flag"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Facebook", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func FetchVideos(APIKey string) (models.Videos, error) {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	service, err := youtube.NewService(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	handleError(err, "")

	// Make Videos Object
	Videos := make(models.Videos, 0)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			Videos = append(Videos, models.Video{
				VideoId:      item.Id.VideoId,
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
				PublishedAt:  item.Snippet.PublishedAt,
			})
		}
	}

	return Videos, err
}

func StoreVideos(videos models.Videos) {
	videoCollection := config.MI.DB.Collection("videos")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := videoCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "videoId", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic(err)
	}

	for _, video := range videos {
		doc := bson.D{
			{Key: "videoId", Value: video.VideoId},
			{Key: "title", Value: video.Title},
			{Key: "description", Value: video.Description},
			{Key: "thumbnailUrl", Value: video.ThumbnailURL},
			{Key: "publishedAt", Value: video.PublishedAt},
		}

		video, err := videoCollection.InsertOne(ctx, doc, &options.InsertOneOptions{})
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Inserted", video, "videos")
	}
}

func handleError(err error, s string) {
	if err != nil {
		log.Fatalf("%v: %v", s, err)
	}
}
