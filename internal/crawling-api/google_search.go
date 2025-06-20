package crawling_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (c *Client) GoogleSearch(query string) (googleSearchResponse, error) {
	_ = godotenv.Load()
	googleAPIKey := os.Getenv("GOOGLE_API_KEY")
	cseAPIKey := os.Getenv("CSE_API_KEY")

	url := "https://www.googleapis.com/customsearch/v1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return googleSearchResponse{}, err
	}

	q := req.URL.Query()
	q.Add("key", googleAPIKey)
	q.Add("cx", cseAPIKey)
	q.Add("q", query)

	req.URL.RawQuery = q.Encode()
	println(req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return googleSearchResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return googleSearchResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return googleSearchResponse{}, err
	}

	response := googleSearchResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return googleSearchResponse{}, err
	}

	return response, nil
}
