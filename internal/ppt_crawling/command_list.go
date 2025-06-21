package ppt_crawling

import "fmt"

func commandList(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("                URL                ")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println()

	count := 0

	for query, url := range cfg.DownloadURL {
		count++
		fmt.Printf("[%d] %s: \n", count, query)
		fmt.Println()
		print(fmt.Sprintf(" - %s\n", url))
		fmt.Println()
	}

	return nil
}
