package kakao

import (
	"fmt"
)

func (cfg *Config) login() error {
	fmt.Println("Login to Kakao Talk...")
	code, err := cfg.getAuthCode()
	if err != nil {
		return err
	}
	cfg.authCode = code
	token, err := cfg.getToken()
	if err != nil {
		return err
	}
	cfg.token = token
	return nil
}
