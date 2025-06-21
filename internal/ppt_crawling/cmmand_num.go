package ppt_crawling

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandNum(cfg *Config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: num")
	}

	urls, err := numRepl(cfg)
	if err != nil {
		return err
	}

	if len(urls) == 0 {
		return fmt.Errorf("no url found")
	}

	return nil
}

func numRepl(cfg *Config) ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	urls := []string{}
	for {

		fmt.Print("num > ")
		var lines []string
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

		postfix := "장 ppt"
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

			urls, err := numSearch(cfg, fullQuery)
			if err != nil {
				if err.Error() == "no search result found" {
					fmt.Println(err)
					continue
				}
				return nil, err
			}

			cfg.DownloadURL[queryString] = urls
			fmt.Printf("Processing %s\n", queryString)
		}

	}

	return urls, nil
}

func numSearch(cfg *Config, query string) (string, error) {

	response, err := cfg.Client.GoogleSearch(query)
	if err != nil {
		return "", err
	}

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryNum) {
			link = item.Link
		}
	}

	if link == "" {
		return "", fmt.Errorf("no search result found")
	}

	str, err := getHTML(link)
	if err != nil {
		return "", fmt.Errorf("error get HTML: %v", err)
	}

	imageblocks, err := getBlocks(str, "span", "class", "imageblock")
	if err != nil {
		return "", fmt.Errorf("error get imageblock: %v", err)
	}

	var imageUrls []string

	for _, imageblock := range imageblocks {
		hrefs, _ := getURLFromBlock(imageblock)
		for _, href := range hrefs {
			imageUrls = append(imageUrls, fmt.Sprintf("%s?original", href))
		}
	}

	fileblocks, err := getBlocks(str, "figure", "class", "fileblock")
	if err != nil {
		return "", fmt.Errorf("error get fileblock: %v", err)
	}

	var fileUrls []string

	for _, fileblock := range fileblocks {
		hrefs, _ := getURLFromBlock(fileblock)
		fileUrls = append(fileUrls, hrefs...)

	}
	var url string

	if imageUrls == nil {
		if len(fileUrls) == 3 {
			url = fileUrls[1]
		} else {
			url = fileUrls[0]
		}
	} else {
		if len(imageUrls) == 3 {
			url = imageUrls[1]
		} else {
			url = imageUrls[0]
		}
	}

	return url, nil
}
