package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

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
