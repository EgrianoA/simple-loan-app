package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Loans struct {
	Loans []LoanObj `json:"loans"`
}

type LoanObj struct {
	Loan_id           string
	Name              string `json:"name"`
	KTP               string `json:"ktp"`
	Loan_amount       uint32 `json:"loan_amount"`
	Loan_period_month uint8  `json:"loan_period_month"`
	Loan_purpose      string `json:"loan_purpose"`
	DOB               string `json:"dob"`
	Sex               string `json:"sex"`
}

func CreateLoan(c *gin.Context) {
	var requestBody LoanObj
	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	// path, _ := os.Getwd()
	// fmt.Println(path)
	savedLoanJson, _ := os.Open("helpers/dummyLoanData.json")
	fmt.Println("The File is opened successfully...")
	defer savedLoanJson.Close()
	byteValue, _ := ioutil.ReadAll(savedLoanJson)

	var loans Loans

	json.Unmarshal(byteValue, &loans)

	mergedLoanArr := append(loans.Loans, requestBody)
	fmt.Println("loans: ", loans)
	fmt.Println("loans.loans: ", loans.Loans)
	fmt.Println("requestBody: ", requestBody)
	fmt.Println("mergedArr:", mergedLoanArr)
	// file, _ := json.MarshalIndent(mergedLoanArr, "", " ")
	// _ = ioutil.WriteFile("helpers/dummyLoanData.json", file, 0644)

	for i := 0; i < len(loans.Loans); i++ {
		fmt.Println(loans.Loans[i].Name)
	}

	fmt.Println(requestBody.Name)
	c.JSON(http.StatusOK, gin.H{"data": requestBody})
}
