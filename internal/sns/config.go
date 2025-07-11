package sns

import (
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/color"
	"github.com/yanmoyy/our-tools/internal/sns/kakao"
)

type snsType string

const (
	Default snsType = "default"
	Kakao   snsType = "kakao"
)

func (snsType snsType) Upper() string {
	return strings.ToUpper(string(snsType))
}

func (snsType snsType) ColorString(s string) string {
	switch snsType {
	case Default:
		return color.Green.ColorString(s)
	case Kakao:
		return color.Yellow.ColorString(s)
	default:
		return ""
	}
}

type config struct {
	kakao   *kakao.Config
	snsType snsType
}

func NewConfig() (*config, error) {
	return &config{
		snsType: Default,
	}, nil
}

// setMode sets mode and initializes config for the mode
func (cfg *config) setMode(snsType snsType) error {
	switch snsType {
	case Kakao:
		cfg.snsType = snsType
		if cfg.kakao == nil {
			config, err := kakao.NewConfig()
			if err != nil {
				return err
			}
			cfg.kakao = config
		}
	default:
		cfg.snsType = Default
	}
	return nil
}
