package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanRequestBody struct {
	Name string
}

func CreateLoan(c *gin.Context) {
	var requestBody LoanRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	fmt.Println(requestBody.Name)
	c.JSON(http.StatusOK, gin.H{"data": requestBody})
}
