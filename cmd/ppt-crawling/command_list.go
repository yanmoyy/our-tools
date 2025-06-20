package main

import "fmt"

func commandList(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("                URL                ")
	fmt.Println("===================================")
	fmt.Println()
	for query, url := range cfg.downloadURL {
		fmt.Printf("%s: %s\n", query, url)
	}
	fmt.Println()

	return nil
}
