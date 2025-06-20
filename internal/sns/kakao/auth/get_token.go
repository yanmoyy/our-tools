package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// get auth token from Kakao Talk.
// See more info:
// https://developers.kakao.com/docs/latest/en/kakaologin/rest-api#request-code
func (cfg *Config) getToken() (token, error) {
	reqBody := url.Values{}
	reqBody.Add("client_id", cfg.apiKey)
	reqBody.Add("code", cfg.authCode)
	reqBody.Add("redirect_uri", cfg.redirectURI)
	reqBody.Add("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", getTokenURL, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return token{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return token{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return token{}, fmt.Errorf("failed to get token (status: %s)", resp.Status)
	}
	var t token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return token{}, err
	}
	return t, nil
}
