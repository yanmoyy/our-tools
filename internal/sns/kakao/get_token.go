package kakao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type token struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return token{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return token{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// print the status code and body
		fmt.Printf("Response status: %s\n", resp.Status)
		body, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("Response body: %s\n", body)
		return token{}, fmt.Errorf("failed to get token")
	}
	var t token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return token{}, err
	}
	return t, nil
}
