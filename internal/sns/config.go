package sns

import (
	"fmt"
	"os"

	"github.com/yanmoyy/our-tools/internal/sns/kakao"
)

type snsType string

const (
	Unknown snsType = "unknown"
	Kakao   snsType = "kakao"
)

type config struct {
	kakaoConfig *kakao.Config
	snsType     snsType
}

func NewConfig() (*config, error) {
	kakaoApiKey := os.Getenv("KAKAO_API_KEY")
	if kakaoApiKey == "" {
		return nil, fmt.Errorf("KAKAO_API_KEY is not set")
	}
	kakaoCfg := kakao.NewConfig(kakaoApiKey, "http://localhost:8080/callback")
	return &config{
		kakaoConfig: kakaoCfg,
		snsType:     Unknown,
	}, nil
}
