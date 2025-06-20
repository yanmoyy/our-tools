package crawling_api

import "fmt"

func commandList(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("                URL                ")
	fmt.Println("===================================")
	fmt.Println()
	for query, url := range cfg.DownloadURL {
		fmt.Printf("%s: %s\n", query, url)
	}
	fmt.Println()

	return nil
}
