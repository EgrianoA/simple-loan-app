package main

import (
	"bytes"
	"encoding/json"
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

type LoanObjRes struct {
	data loans.LoanObj `json:"loan_id"`
}

func TestCreateLoan(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateLoanInvalidName(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User 12",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidMinName(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "User",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidGender(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "he",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidKTP(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "abc4072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidDOB(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "dd/mm/yyyy",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanDOBAndNIKNotMatch(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1500000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "29/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidLoanAmount(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       20000,
		Loan_period_month: 6,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidLoanPeriod(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1200000,
		Loan_period_month: 255,
		Loan_purpose:      "vacation",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLoanInvalidLoanPurpose(t *testing.T) {
	r := setUpRouter()
	r.POST("/api/loan/create", loans.CreateLoan)
	newLoanTest := loans.LoanObj{
		Name:              "Demo User",
		KTP:               "3174072808971238",
		Loan_amount:       1200000,
		Loan_period_month: 6,
		Loan_purpose:      "have fun",
		DOB:               "28/08/1997",
		Gender:            "male",
	}
	jsonValue, _ := json.Marshal(newLoanTest)
	req, _ := http.NewRequest("POST", "/api/loan/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindByLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-180822-7526", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInvalidLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-18xyz2-7526", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNotExistLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-180822-0000", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestFindByKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/3174052508971237", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInvalidFindKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/31740abc08971237", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindNotExistKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/3174072108971237", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
