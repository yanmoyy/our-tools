package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func commandTitle(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: title <count>")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	urls, err := titleRepl(cfg, count)
	if err != nil {
		return err
	}

	for _, url := range urls {
		fmt.Println(url)
	}

	return nil
}
func titleRepl(cfg *config, count int) ([]string, error) {

	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}
	for len(urls) < count {

		fmt.Print("title > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		joinedQuery := strings.Join(words, " ")
		postfix := " ppt"

		fullQuery := fmt.Sprintf("%s%s", joinedQuery, postfix)

		fmt.Println("fullQuery: ", fullQuery)

		url, err := titleSearch(cfg, fullQuery)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func titleSearch(cfg *config, query string) (string, error) {
	response, err := cfg.client.GoogleSearch(query)
	if err != nil {
		return "", err
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
		return "", err
	}
	fmt.Println("ImageBlocks start")

	imageblocks, err := getImageBlocks(str)
	if err != nil {
		return "", err
	}
	fmt.Println(len(imageblocks))

	var urls []string

	for _, imageblock := range imageblocks {
		hrefs, _ := getURLFromBlock(imageblock)
		urls = append(urls, hrefs...)
	}

	url := fmt.Sprintf("%s%s", urls[1], "?original")

	return url, nil
}
