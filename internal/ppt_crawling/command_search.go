package ppt_crawling

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func commandSearch(cfg *Config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("usage: search <start>")
	} else if len(args) == 1 {
		start := args[0]
		startInt, err := strconv.Atoi(start)
		if err != nil {
			return fmt.Errorf("start argument must be integer")
		}
		if startInt > 90 {
			return fmt.Errorf("it is not allowed to set start args more than 90")
		}
		cfg.StartPages = startInt
	}

	err := SearchRepl(cfg)
	if err != nil {
		return err
	}

	return nil
}

func SearchRepl(cfg *Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var lines []string
		fmt.Print("search > ")
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			if line == "exit" {
				cfg.StartPages = 1
				fmt.Println("search exit....")
				return nil
			}

			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading input in SearchRepl: %v", err)
		}

		if len(lines) == 0 {
			break
		}

		for _, line := range lines {
			words := cleanInput(line)

			if len(words) == 0 {
				continue
			}
			if len(words) == 1 && words[0] == "exit" {
				return nil
			}
			queryString := strings.Join(words, " ")

			urls, err := Search(cfg, queryString)
			if err != nil {
				if err.Error() == "no search result found in Search" {
					fmt.Printf("Search error in SearchRepl\n\n%v\n", err)
					continue
				}
				return err
			}
			if urls != nil {
				cfg.DownloadURL[queryString] = urls[0]
				fmt.Printf("Processing %s\n", queryString)
			} else {
				fmt.Printf("no url found in %s\n", queryString)
			}
			fmt.Println("===================================")
			fmt.Println()
			fmt.Println()

		}

	}

	return nil
}

func Search(cfg *Config, query string) ([]string, error) {
	urls, err := cfg.Client.GoogleSearch(query, cfg.StartPages)
	if err != nil {
		return nil, fmt.Errorf("error GoogleSearch in Search\n%v", err)

	}

	return urls, nil
}
