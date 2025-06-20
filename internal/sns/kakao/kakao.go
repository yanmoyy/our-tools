package kakao

import (
	"fmt"
	"os"

	"github.com/yanmoyy/our-tools/internal/sns/cli"
	"github.com/yanmoyy/our-tools/internal/sns/kakao/auth"
	"github.com/yanmoyy/our-tools/internal/sns/kakao/social"
)

type Config struct {
	auth    *auth.Config
	friends []social.Friend
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

func GetCommands() map[string]cli.Command[Config] {
	return map[string]cli.Command[Config]{
		"login": {
			Name:        "login",
			Description: "Login to Kakao Talk",
			Callback:    commandLogin,
		},
		"send": {
			Name:        "send",
			Description: "Send message to Kakao Talk",
			Callback:    commandSend,
			Helper:      printSendHelp,
		},
		"ls": {
			Name:        "ls",
			Description: "List friends",
			Callback:    commandListFriends,
		},
	}
}
