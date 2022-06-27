package cron

import (
	"fampay-youtube/models"
	"fampay-youtube/util"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/robfig/cron.v2"
)

func StartYoutubeFetch() {
	apiKeyString := os.Getenv("YOUTUBE_API_KEYS")
	apiKeys := strings.Split(apiKeyString, ";")
	c := cron.New()

	keyCount := len(apiKeys)
	c.AddFunc("@every 10s", func() {
		if keyCount > 0 {
			videos, err := fetchWithAPIKey(c, apiKeys[len(apiKeys)-keyCount])
			if err != nil {
				log.Println(err)
				keyCount--
			}
			util.StoreVideos(videos)
		} else {
			log.Println("All API keys exhausted")
			c.Stop()
		}
	})

	c.Start()
}

func fetchWithAPIKey(c *cron.Cron, APIKey string) (models.Videos, error) {
	fmt.Printf("APIKey:%s", APIKey)
	videos, err := util.FetchVideos(APIKey)
	if err != nil {
		log.Printf("APIKey:[%s] Exausted", APIKey)
		c.Stop()
	}
	return videos, err
}
