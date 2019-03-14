package config

// Github type
type GitHub struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	OauthState   string
}
