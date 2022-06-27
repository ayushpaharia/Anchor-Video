package util

import (
	"context"
	"fampay-youtube/models"
	"flag"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Cricket", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func FetchVideos(APIKey string) (models.Videos, error) {
	flag.Parse()
	ctx := context.Background()

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
	fmt.Println("\nStoring videos")
	for _, video := range videos {
		fmt.Println(video.Title)
	}
}

func handleError(err error, s string) {
	if err != nil {
		log.Fatalf("%v: %v", s, err)
	}
}
