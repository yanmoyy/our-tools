package main

import (
	"fmt"
	"html"
	"strings"
)

func commandTitle(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: title <query>")
	}
	query := args[:]
	blank := html.EscapeString(" ")
	fullQuery := strings.Join(query, blank)

	fmt.Println("fullQuery: ", fullQuery)

	response, err := cfg.client.GoogleSearch(fullQuery)
	if err != nil {
		return err
	}

	println("start")

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryTitle) {
			link = item.Link
		}
	}
	fmt.Println(link)
	println("end")

	str, err := getHTML(link)
	if err != nil {
		return err
	}
	fmt.Println("ImageBlocks start")

	imageblocks, err := getImageBlocks(str)
	if err != nil {
		return err
	}
	fmt.Println(len(imageblocks))

	var urls []string

	for _, imageblock := range imageblocks {
		hrefs, _ := getURLFromImageBlock(imageblock)
		urls = append(urls, hrefs...)
	}

	fmt.Printf("%s?original\n", urls[1])

	return nil
}
