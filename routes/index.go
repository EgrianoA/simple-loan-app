package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup) {
	loanRoutes(superRoute)
}
