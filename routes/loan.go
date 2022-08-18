package routes

import (
	"net/http"
	loans "simple-loan-app/controllers"

	"github.com/gin-gonic/gin"
)

func loanRoutes(superRoute *gin.RouterGroup) {
	superRoute.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	loanRouter := superRoute.Group("/loan")
	{
		loanRouter.POST("/create", loans.CreateLoan)
		loanRouter.GET("/findById/:loan_id", loans.FindLoanById)
		loanRouter.GET("/findByKTP/:ktp", loans.FindLoadByKTP)
	}
}
