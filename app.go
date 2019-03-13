package main

import (
	"github.com/dariubs/commit/route"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	route.Routes(router)

	router.Run()
}
