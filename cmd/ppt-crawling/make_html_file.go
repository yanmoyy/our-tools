package main

import (
	"fmt"
	"os"
)

func makeHTMLFile(htmlContent string) error {
	filePath := "./cmd/ppt-crawling/output.html"

	// Write the HTML string to a file
	err := os.WriteFile(filePath, []byte(htmlContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("HTML file created successfully: %s\n", filePath)

	return nil
}
