/*
Package cmd/web implements commit service web app.
See https://github.com/dariubs/commit for more information.
*/
package main

import (
	"log"

	"github.com/dariubs/commit/config"
	"github.com/dariubs/commit/route"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error to load config")
	}

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	route.Routes(router)

	router.Run(config.Web.Port)
}
