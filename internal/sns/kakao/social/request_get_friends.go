package social

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/auth"
)

type Friend struct {
	UUID            string `json:"uuid"`
	ProfileNickname string `json:"profile_nickname"`
}

type getFriendsResponse struct {
	Elements   []Friend `json:"elements"`
	TotalCount int      `json:"total_count"`
	BeforeURL  string   `json:"before_url"`
	AfterURL   string   `json:"after_url"`
}

func requestGetFriends(token string) (getFriendsResponse, error) {
	req, err := http.NewRequest("GET", getFriendsURL, nil)
	if err != nil {
		return getFriendsResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", auth.GetBearerToken(token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return getFriendsResponse{}, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return getFriendsResponse{}, fmt.Errorf("failed to get friends (status: %s)", resp.Status)
	}
	var response getFriendsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return getFriendsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}
	return response, nil
}
