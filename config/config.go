package config

// CONFIG type for app
type CONFIG struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	OauthState   string
}
