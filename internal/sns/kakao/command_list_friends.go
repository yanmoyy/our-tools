package kakao

import (
	"fmt"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/social"
)

func commandListFriends(cfg *Config, args ...string) error {
	if cfg.friends == nil {
		authToken := cfg.auth.GetAccessToken()
		if authToken == "" {
			return fmt.Errorf("not logged in")
		}
		friends, err := social.GetFriends(authToken)
		if err != nil {
			return err
		}
		cfg.friends = friends
	}

	fmt.Println("Friends:")
	fmt.Printf("%d. %s\n", 0, "me")
	for i, friend := range cfg.friends {
		fmt.Printf("%d. %s\n", i+1, friend.ProfileNickname)
	}
	return nil
}
