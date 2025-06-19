package kakao

import "fmt"

func (cfg *Config) login() error {
	fmt.Println("Login to Kakao Talk...")
	err := cfg.getAuthCode()
	if err != nil {
		return err
	}
	return nil
}
