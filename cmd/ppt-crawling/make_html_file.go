package main

import (
	"os"
)

func makeHTMLFile(htmlContent string) error {
	filePath := "./cmd/ppt-crawling/output.html"

	// Write the HTML string to a file
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, []byte(htmlContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
