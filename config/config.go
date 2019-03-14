package config

// Config type for app
type Config struct {
	GitHub GitHub `toml:"github"`
}

// Github type
type GitHub struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	OauthState   string
}
