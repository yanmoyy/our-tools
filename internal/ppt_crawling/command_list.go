package ppt_crawling

import "fmt"

func commandList(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("                URL                ")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println()

	for query, url := range cfg.DownloadURL {
		print(fmt.Sprintf("%s : %s\n", query, url))
	}

	return nil
}
