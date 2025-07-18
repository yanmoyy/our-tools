package ppt_crawling

import "fmt"

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the parise crawler!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
