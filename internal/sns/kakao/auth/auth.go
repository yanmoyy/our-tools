package auth

import (
	"fmt"
)

// #nosec G101
const (
	getAuthCodeURL  = "https://kauth.kakao.com/oauth/authorize"
	getTokenURL     = "https://kauth.kakao.com/oauth/token"
	getTokenInfoURL = "https://kapi.kakao.com/v1/user/access_token_info"
)

type Config struct {
	apiKey      string
	redirectURI string
	authCode    string
	token       token
}

func NewConfig(apiKey, redirectURI string) *Config {
	return &Config{
		apiKey:      apiKey,
		redirectURI: redirectURI,
	}
}

type token struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

func (cfg *Config) Login() error {
	fmt.Println("Login to Kakao Talk...")
	code, err := cfg.getCode()
	if err != nil {
		return err
	}
	cfg.authCode = code
	token, err := cfg.getToken()
	if err != nil {
		return err
	}
	cfg.token = token
	fmt.Println("Login to Kakao Talk successfully!")
	return nil
}

func (cfg *Config) UpdateToken() error {
	// refreshed token
	token, err := cfg.getRefreshedToken()
	if err != nil {
		return err
	}
	if token.RefreshToken != "" {
		// refresh token is expired
		cfg.token = token
	} else {
		// refresh token is not expired
		cfg.token.AccessToken = token.AccessToken
		cfg.token.ExpiresIn = token.ExpiresIn
	}
	return nil
}
