package main

import (
	"time"

	crawling_api "github.com/yanmoyy/our-tools/internal/crawling-api"
)

func main() {
	pptClient := crawling_api.NewClient(5 * time.Second)
	cfg := &config{
		client: pptClient,
	}
	startRepl(cfg)

}
