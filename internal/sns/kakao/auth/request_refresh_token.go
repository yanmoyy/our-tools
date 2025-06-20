package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/header"
)

// get refreshed token from Kakao Talk.
// See more info:
// https://developers.kakao.com/docs/latest/en/kakaologin/rest-api#refresh-token
func requestRefreshToken(apiKey, refreshToken string) (token, error) {
	reqBody := url.Values{}
	reqBody.Add("client_id", apiKey)
	reqBody.Add("grant_type", "refresh_token")
	reqBody.Add("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", getTokenURL, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return token{}, err
	}
	req.Header.Set("Content-Type", header.UrlEncoded)
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
