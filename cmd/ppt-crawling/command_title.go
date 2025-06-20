package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandTitle(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: title")
	}

	urls, queries, err := titleRepl(cfg)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("                URL                ")
	fmt.Println("===================================")
	fmt.Println()

	if len(urls) == 0 {
		return fmt.Errorf("no url found\n")
	}

	for i, url := range urls {
		fmt.Printf("%s: %s\n", queries[i], url)
	}

	return nil
}
func titleRepl(cfg *config) ([]string, []string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}
	queries := []string{}
	for {

		fmt.Print("title > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		if len(words) == 1 && words[0] == "exit" {
			break
		}
		joinedQuery := strings.Join(words, " ")
		postfix := " ppt"

		queries = append(queries, joinedQuery)

		fullQuery := fmt.Sprintf("%s%s", joinedQuery, postfix)

		fmt.Println("fullQuery: ", fullQuery)

		url, err := titleSearch(cfg, fullQuery)
		if err != nil {
			if err.Error() == "no search result found" {
				fmt.Println(err)
				continue
			}
			return nil, nil, err
		}
		urls = append(urls, url)
	}

	return urls, queries, nil
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

	if link == "" {
		return "", fmt.Errorf("no search result found")
	}

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

	if len(urls) == 0 {
		return "", fmt.Errorf("no url found")
	}

	url := fmt.Sprintf("%s%s", urls[1], "?original")

	return url, nil
}
