package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type tokenInfo struct {
	ID        int64 `json:"id"` // service user ID
	ExpiresIn int   `json:"expires_in"`
	AppID     int   `json:"app_id"`
}

// get token info from Kakao Talk.
// See more info:
// https://developers.kakao.com/docs/latest/en/kakaologin/rest-api#get-token-info
func (cfg *Config) getTokenInfo() (tokenInfo, error) {
	req, err := http.NewRequest("GET", getTokenInfoURL, nil)
	if err != nil {
		return tokenInfo{}, err
	}
	req.Header.Set("Authorization", getBearerToken(cfg.token.AccessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return tokenInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return tokenInfo{}, fmt.Errorf("failed to get token info (status: %s)", resp.Status)
	}
	var t tokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return tokenInfo{}, err
	}
	return t, nil
}

func getBearerToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
