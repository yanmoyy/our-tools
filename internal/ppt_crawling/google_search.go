package ppt_crawling

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func (c *Client) GoogleSearch(query string, start int) ([]string, error) {
	fmt.Println("===================================")
	fmt.Printf("Starting Searching... (%s)\n\n", query)
	_ = godotenv.Load()
	googleAPIKey := os.Getenv("GOOGLE_API_KEY")
	cseAPIKey := os.Getenv("CSE_API_KEY")

	url := "https://www.googleapis.com/customsearch/v1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error NewRequest in GoogleSearch\n%v", err)
	}

	q := req.URL.Query()
	q.Add("key", googleAPIKey)
	q.Add("cx", cseAPIKey)
	q.Add("q", query)

	q.Add("start", strconv.Itoa(start))

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error Do in GoogleSearch\n%v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code in GoogleSearch: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error ReadAll in GoogleSearch\n%v", err)
	}

	response := googleSearchResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error Unmarshal in GoogleSearch\n%v", err)
	}

	urls, err := getDownloadUrls(response)
	if err != nil {
		return nil, fmt.Errorf("error getDownloadUrls in GoogleSearch\n%v", err)
	}

	return urls, nil
}
