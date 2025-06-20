package ppt_crawling

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandDownload(cfg *Config, args ...string) error {
	var filepath string
	if len(args) < 1 {
		filepath = "/mnt/c/Users/alstj/Downloads/"
	} else {
		filepath = args[0]
	}

	fmt.Println()
	fmt.Println("download")
	fmt.Println()

	for query, url := range cfg.DownloadURL {

		filepath := fmt.Sprintf("%s%s.ppt", filepath, query)
		err := DownloadFile(url, filepath)
		fmt.Printf("Downloading %s\n", filepath)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func DownloadFile(url string, filepath string) error {
	// Create the file
	fmt.Println("url: ", url)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
