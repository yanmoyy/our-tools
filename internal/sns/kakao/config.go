package kakao

import (
	"fmt"
	"os"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/auth"
)

type Config struct {
	auth *auth.Config
}

func NewConfig() (*Config, error) {
	apiKey := os.Getenv("KAKAO_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("KAKAO_API_KEY is not set")
	}
	redirectURI := "http://localhost:8080/oauth"
	return &Config{
		auth: auth.NewConfig(apiKey, redirectURI),
	}, nil
}

func (cfg *Config) StartMode() error {
	err := cfg.auth.Login()
	if err != nil {
		return err
	}
	return nil
}
