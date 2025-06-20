package main

import "fmt"

func commandDownload(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("download")
	fmt.Println()

	for query, url := range cfg.downloadURL {
		fmt.Printf("%s : %s\n", query, url)
	}

	return nil
}
