package main

import (
	"time"

	"github.com/yanmoyy/our-tools/internal/ppt_crawling"
)

func main() {
	pptClient := ppt_crawling.NewClient(5 * time.Second)
	cfg := &ppt_crawling.Config{
		Client:      pptClient,
		DownloadURL: map[string]string{},
		StartPages:  1,
	}

	ppt_crawling.StartRepl(cfg)

}
