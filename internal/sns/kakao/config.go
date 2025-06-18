package kakao

const (
	AuthURL = "https://kauth.kakao.com/oauth/authorize"
)

type Config struct {
	apiKey      string
	redirectURI string
}

func NewConfig(apiKey, redirectURI string) *Config {
	return &Config{
		apiKey:      apiKey,
		redirectURI: redirectURI,
	}
}
