package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func getURLFromImageBlock(imageblock *html.Node) ([]string, error) {
	urls := getDataFromKey(imageblock.FirstChild, "a", "href")
	if len(urls) > 0 {
		return urls, nil
	}

	return nil, fmt.Errorf("urls not found")
}
