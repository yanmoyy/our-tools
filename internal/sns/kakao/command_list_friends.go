package kakao

import (
	"fmt"

	"github.com/yanmoyy/our-tools/internal/sns/kakao/social"
)

func commandListFriends(cfg *Config, args ...string) error {
	fmt.Println("Friends:")
	friends, err := social.GetFriends(cfg.auth.GetAccessToken())
	if err != nil {
		return err
	}
	for _, friend := range friends {
		fmt.Println(friend.ID)
		fmt.Println(friend.UUID)
		fmt.Println(friend.ProfileNickname)
	}
	return nil
}
