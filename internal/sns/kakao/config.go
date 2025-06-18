package kakao

import (
	"fmt"
	"os"
)

const (
	AuthURL = "https://kauth.kakao.com/oauth/authorize"
)

type Config struct {
	apiKey      string
	redirectURI string
}

func NewConfig() (*Config, error) {
	apiKey := os.Getenv("KAKAO_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("KAKAO_API_KEY is not set")
	}
	redirectURI := "http://localhost:8080/callback"
	return &Config{
		apiKey:      apiKey,
		redirectURI: redirectURI,
	}, nil
}
