package routes

import (
	"net/http"
	"simple-loan-app/controllers"

	"github.com/gin-gonic/gin"
)

func loanRoutes(superRoute *gin.RouterGroup) {
	superRoute.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	loanRouter := superRoute.Group("/loan")
	{
		loanRouter.POST("/create", controllers.CreateLoan)
	}
}
