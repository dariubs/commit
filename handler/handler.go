package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/dariubs/commit/config"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	// Config : general config
	Config config.CONFIG

	// OauthStateString -
	OauthStateString string

	oauthConf = &oauth2.Config{}
)

func init() {
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
}

// HomepageHandler : '/'
func HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// LoginPageHandler -
func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// LoginActionHandler -
func LoginActionHandler(c *gin.Context) {
	url := oauthConf.AuthCodeURL(OauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// LoginVerificationHandler -
func LoginVerificationHandler(c *gin.Context) {
	state := c.PostForm("state")
	var NoContext = context.TODO()
	if state != OauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", OauthStateString, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(NoContext, "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
