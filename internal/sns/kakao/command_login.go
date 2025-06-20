package kakao

import "fmt"

func commandLogin(cfg *Config, args ...string) error {
	if cfg.auth.GetAccessToken() != "" {
		fmt.Println("Already logged in")
		return nil
	}
	fmt.Println("Login to Kakao Talk...")
	err := cfg.auth.Login()
	if err != nil {
		return err
	}
	fmt.Println("Login to Kakao Talk successfully!")
	return nil
}
