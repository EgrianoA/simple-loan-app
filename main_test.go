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
	Data loans.LoanObj `json:data`
}

type LoanArrRes struct {
	Data []loans.LoanObj `json:data`
}

type LoanErrRes struct {
	Message string `json:message`
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
	var loan LoanObjRes
	json.Unmarshal(w.Body.Bytes(), &loan)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.IsType(t, LoanObjRes{}, loan)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Name is invalid", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Name is invalid. Name should have minimum two words", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Gender is invalid. Only accept these value: male, female", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The KTP Format is invalid", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The DOB is Invalid", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The DOB and NIK is not match", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Loan Amount is invalid. Loan Amount should between 1.000.000 and 10.000.000", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Maximum loan period is 240", loanErr.Message)
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
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Loan Purpose is invalid. Only accept these value: vacation, renovation, electronics, wedding, rent, car, investment", loanErr.Message)
}

func TestFindByLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-180822-7526", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var loan LoanObjRes
	json.Unmarshal(w.Body.Bytes(), &loan)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.IsType(t, LoanObjRes{}, loan)
}

func TestInvalidLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-18xyz2-7526", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Loan ID is Invalid", loanErr.Message)
}

func TestNotExistLoanId(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findById/:loan_id", loans.FindLoanById)
	req, _ := http.NewRequest("GET", "/api/loan/findById/LOAN-180822-0000", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "Loan ID is Not Found", loanErr.Message)
}

func TestFindByKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/3174052508971237", nil)
	w := httptest.NewRecorder()
	var loan LoanArrRes
	json.Unmarshal(w.Body.Bytes(), &loan)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.GreaterOrEqual(t, 1, len(loan.Data))
}

func TestInvalidFindKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/31740abc08971237", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The KTP Format is invalid", loanErr.Message)
}

func TestFindNotExistKTP(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/loan/findByKTP/:ktp", loans.FindLoadByKTP)
	req, _ := http.NewRequest("GET", "/api/loan/findByKTP/3174099999971237", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var loanErr LoanErrRes
	json.Unmarshal(w.Body.Bytes(), &loanErr)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "There's no loan with this KTP", loanErr.Message)
}
