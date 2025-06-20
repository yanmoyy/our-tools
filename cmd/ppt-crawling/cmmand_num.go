package main

import (
	"fmt"
	"strings"
)

func commandNum(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: title <query>")
	}
	query := args[0]

	postfix := "장 ppt"

	fullQuery := fmt.Sprintf("%s%s", query, postfix)
	fmt.Println("fullQuery: ", fullQuery)

	response, err := cfg.client.GoogleSearch(fullQuery)
	if err != nil {
		return err
	}

	println("start")

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryNum) {
			link = item.Link
		}
	}
	println("end")

	str, err := getHTML(link)
	if err != nil {
		return err
	}
	err = makeHTMLFile(str)
	if err != nil {
		return err
	}

	fileblock, err := getFileBlocks(str)
	if err != nil {
		return err
	}

	var urls []string

	for _, fileblock := range fileblock {
		hrefs, _ := getURLFromBlock(fileblock)
		urls = append(urls, hrefs...)
	}

	for _, url := range urls {
		fmt.Println(url)
	}

	return nil
}
