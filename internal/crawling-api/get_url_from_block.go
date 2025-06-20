package crawling_api

import (
	"fmt"

	"golang.org/x/net/html"
)

func getURLFromBlock(block *html.Node) ([]string, error) {
	urls := getDataFromKey(block.FirstChild, "a", "href")
	if len(urls) > 0 {
		return urls, nil
	}

	return nil, fmt.Errorf("urls not found")
}
