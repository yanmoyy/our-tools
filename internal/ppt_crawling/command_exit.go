package ppt_crawling

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the crwaler... Goodbye!")
	os.Exit(0)
	return nil
}
