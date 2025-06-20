package kakao

import (
	"fmt"
)

func (cfg *Config) login() error {
	fmt.Println("Login to Kakao Talk...")
	code, err := cfg.getCode()
	if err != nil {
		return err
	}
	cfg.authCode = code
	token, err := cfg.getToken()
	if err != nil {
		return err
	}
	cfg.token = token
	fmt.Println("Login to Kakao Talk successfully!")
	return nil
}

func (cfg *Config) updateToken() error {
	// refreshed token
	token, err := cfg.getRefreshedToken()
	if err != nil {
		return err
	}
	if token.RefreshToken != "" {
		// refresh token is expired
		cfg.token = token
	} else {
		// refresh token is not expired
		cfg.token.AccessToken = token.AccessToken
		cfg.token.ExpiresIn = token.ExpiresIn
	}
	return nil
}
