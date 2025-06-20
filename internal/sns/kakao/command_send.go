package kakao

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/msg"
)

func commandSend(cfg *Config, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("argument required")
	}
	target := args[0]
	content := strings.Join(args[1:], " ")
	authToken := cfg.auth.GetAccessToken()
	if authToken == "" {
		return fmt.Errorf("not logged in")
	}
	if target == "me" || target == "0" {
		err := msg.SendToMe(authToken, content)
		if err != nil {
			return fmt.Errorf("failed to send message to me: %s", err)
		}
		fmt.Println("Message sent to me")
		return nil
	}
	// send to specific user from friends list
	if cfg.friends == nil {
		return fmt.Errorf("friends list is empty (ls command required)")
	}
	for i, friend := range cfg.friends {
		// check idx, nickname
		if friend.ProfileNickname == target ||
			strconv.Itoa(i+1) == target {
			err := msg.SendToFriend(authToken, friend.UUID, content)
			if err != nil {
				return fmt.Errorf("failed to send message: %s", err)
			}
			fmt.Printf("Message sent to %s\n", friend.ProfileNickname)
			return nil
		}
	}
	return fmt.Errorf("user %s not found", target)
}

func printSendHelp() {
	fmt.Println("Usage: send [target] [message]")
	fmt.Println("Target (index | nickname):")
	fmt.Println("  - me: send to me")
	fmt.Println("  - user: send to specific user")
	fmt.Println()
}
