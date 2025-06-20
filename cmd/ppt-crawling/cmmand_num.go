package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandNum(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: num")
	}

	urls, queries, err := numRepl(cfg)
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
		fmt.Printf("%s: %s\n\n", queries[i], url)
	}

	return nil
}

func numRepl(cfg *config) ([]string, []string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}
	queries := []string{}
	for {

		fmt.Print("num > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if len(words) == 1 && words[0] == "exit" {
			break
		}
		postfix := "장 ppt"

		queries = append(queries, words[0])

		fullQuery := fmt.Sprintf("%s%s", words[0], postfix)
		fmt.Println("fullQuery: ", fullQuery)

		url, err := numSearch(cfg, fullQuery)
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

func numSearch(cfg *config, query string) (string, error) {

	response, err := cfg.client.GoogleSearch(query)
	if err != nil {
		return "", err
	}

	println("start")

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryNum) {
			link = item.Link
		}
	}

	fmt.Println("Link :", link)
	println("end")

	if link == "" {
		return "", fmt.Errorf("no search result found")
	}

	str, err := getHTML(link)
	if err != nil {
		return "", err
	}
	err = makeHTMLFile(str)
	if err != nil {
		return "", err
	}

	fileblock, err := getFileBlocks(str)
	if err != nil {
		return "", err
	}

	var urls []string

	for _, fileblock := range fileblock {
		hrefs, _ := getURLFromBlock(fileblock)
		urls = append(urls, hrefs...)
	}

	if len(urls) == 0 {
		return "", fmt.Errorf("no url found")
	}

	return urls[0], nil
}
