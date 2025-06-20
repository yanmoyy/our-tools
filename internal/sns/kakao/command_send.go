package kakao

import (
	"fmt"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/msg"
)

func commandSend(cfg *Config, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("argument required")
	}
	target := args[0]
	content := strings.Join(args[1:], " ")
	err := msg.Send(cfg.auth.GetAccessToken(), target, content)
	if err != nil {
		return fmt.Errorf("failed to send message: %s", err)
	}
	return nil
}

func printSendHelp() {
	fmt.Println("Usage: send [target] [message]")
	fmt.Println("Target:")
	fmt.Println("  - me: send to me")
	fmt.Println("  - user: send to specific user")
	fmt.Println()
}
