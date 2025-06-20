package crawling_api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandTitle(cfg *Config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: title")
	}

	urls, err := titleRepl(cfg)
	if err != nil {
		return err
	}

	if len(urls) == 0 {
		return fmt.Errorf("no url found")
	}

	return nil
}

func titleRepl(cfg *Config) ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}
	for {
		var lines []string
		fmt.Print("title > ")
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" || line == "exit" {
				break
			}
			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading input: %v", err)
		}

		if len(lines) == 0 {
			break
		}

		postfix := " ppt"
		for _, line := range lines {
			words := cleanInput(line)
			if len(words) == 0 {
				continue
			}
			if len(words) == 1 && words[0] == "exit" {
				return urls, nil
			}
			queryString := strings.Join(words, " ")
			fullQuery := fmt.Sprintf("%s%s", queryString, postfix)

			url, err := titleSearch(cfg, fullQuery)
			if err != nil {
				if err.Error() == "no search result found" {
					fmt.Println(err)
					continue
				}
				return nil, err
			}
			cfg.DownloadURL[queryString] = url
			urls = append(urls, url)
			fmt.Printf("Processing %s\n", queryString)
		}

	}

	return urls, nil
}

func titleSearch(cfg *Config, query string) (string, error) {
	response, err := cfg.Client.GoogleSearch(query)
	if err != nil {
		return "", err
	}

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryTitle) {
			link = item.Link
		}
	}

	if link == "" {
		return "", fmt.Errorf("no search result found")
	}

	str, err := getHTML(link)
	if err != nil {
		return "", err
	}

	imageblocks, err := getImageBlocks(str)
	if err != nil {
		return "", err
	}

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
