package main

import (
	"net/http"
	"net/http/httptest"
	loans "simple-loan-app/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// type LoanObjRes struct {
// 	data struct loans.LoanObj `json:"loan_id"`
// }

func TestFindByKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-180822-7526", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// var loan loans.LoanObj
	// fmt.Println(w.Body)
	// fmt.Println(json.Unmarshal(w.Body.Bytes(), &loan))
	// json.Unmarshal(w.Body.Bytes(), &loan)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)
}
