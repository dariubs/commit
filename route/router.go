package route

import (
	"github.com/dariubs/commit/handler"
	"github.com/gin-gonic/gin"
)

// Routes : setup routers
func Routes(router *gin.Engine) {
	router.GET("/", handler.HomepageHandler)

	// user login manager
	user := router.Group("/user")
	{
		user.GET("/login", handler.LoginPageHandler)
		user.GET("/login/do", handler.LoginActionHandler)
		user.Any("/verify", handler.LoginVerificationHandler) // Callback endpoint
	}

	// agora : list of issues to solve
	agora := router.Group("/agora")
	{
		agora.GET("/")
		agora.GET("/language/:lang/")
	}

}
