package ppt_crawling

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	resp, err := http.Get(rawURL)
	if resp.StatusCode > 399 && resp.StatusCode < 499 {
		return "", fmt.Errorf("error status code in getHTML: %d", resp.StatusCode)
	}
	if err != nil {
		return "", fmt.Errorf("error Get in getHTML\n%v", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response in getHTML: %s", contentType)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error ReadAll in getHTML\n%v", err)
	}

	return string(data), nil
}
