package sns

import (
	"fmt"

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
	return &config{
		snsType: Unknown,
	}, nil
}

// setMode sets mode and initializes config for the mode
func (cfg *config) setMode(snsType snsType) error {
	switch snsType {
	case Kakao:
		cfg.snsType = snsType
		if cfg.kakaoConfig == nil {
			config, err := kakao.NewConfig()
			if err != nil {
				return err
			}
			cfg.kakaoConfig = config
		}
	default:
		return fmt.Errorf("unknown sns type: %s", snsType)
	}
	return nil
}
