package kakao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// get refreshed token from Kakao Talk.
// See more info:
// https://developers.kakao.com/docs/latest/en/kakaologin/rest-api#refresh-token
func (cfg *Config) getRefreshedToken() (token, error) {
	reqBody := url.Values{}
	reqBody.Add("client_id", cfg.apiKey)
	reqBody.Add("grant_type", "refresh_token")
	reqBody.Add("refresh_token", cfg.token.RefreshToken)

	req, err := http.NewRequest("POST", getTokenURL, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return token{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	if err != nil {
		return token{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return token{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return token{}, fmt.Errorf("failed to refresh token (status: %s)", resp.Status)
	}
	var t token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return token{}, err
	}
	return t, nil
}
