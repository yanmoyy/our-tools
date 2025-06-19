package kakao

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	apiKey      string
	redirectURI string
	authCode    string
	mu          sync.Mutex
	// accessToken  string
	// refreshToken string
}

func NewConfig() (*Config, error) {
	apiKey := os.Getenv("KAKAO_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("KAKAO_API_KEY is not set")
	}
	redirectURI := "http://localhost:8080/oauth"
	return &Config{
		apiKey:      apiKey,
		redirectURI: redirectURI,
	}, nil
}

func (cfg *Config) StartMode() error {
	err := cfg.login()
	if err != nil {
		return err
	}
	return nil
}
