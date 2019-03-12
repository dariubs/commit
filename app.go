package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	// Config : general config
	Config CONFIG

	// OauthStateString -
	OauthStateString string

	oauthConf = &oauth2.Config{}
)

func main() {
	// load config
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
		return
	}

	oauthConf = &oauth2.Config{
		ClientID:     Config.ClientID,
		ClientSecret: Config.ClientSecret,
		Scopes:       []string{"read:user", "public_repo"},
		Endpoint:     githuboauth.Endpoint,
	}
	OauthStateString = Config.OauthState

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", HomepageHandler)

	user := router.Group("/user")
	{
		user.GET("/login", LoginPageHandler)
		user.GET("/login/do", LoginActionHandler)
		user.Any("/verify", LoginVerificationHandler)
	}

	router.Run()
}
