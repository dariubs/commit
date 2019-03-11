package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	Config           CONFIG
	oauthStateString string
	oauthConf        = &oauth2.Config{}
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
	oauthStateString = Config.OauthState

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

func HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginActionHandler(c *gin.Context) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func LoginVerificationHandler(c *gin.Context) {
	state := c.PostForm("state")
	var NoContext = context.TODO()
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
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

func tokenToJSON(token *oauth2.Token) (string, error) {
	if d, err := json.Marshal(token); err != nil {
		return "", err
	} else {
		return string(d), nil
	}
}

func tokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	var token oauth2.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
