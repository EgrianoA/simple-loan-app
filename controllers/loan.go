package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoanObj struct {
	Loan_id           string `json:"loan_id"`
	Name              string `json:"name"`
	KTP               string `json:"ktp"`
	Loan_amount       uint32 `json:"loan_amount"`
	Loan_period_month uint8  `json:"loan_period_month"`
	Loan_purpose      string `json:"loan_purpose"`
	DOB               string `json:"dob"`
	Gender            string `json:"gender"`
}

func CreateLoan(c *gin.Context) {
	//Getting the request data
	var req LoanObj
	if err := c.BindJSON(&req); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	//Get the current directory
	// path, _ := os.Getwd()
	// fmt.Println(path)

	var loans []LoanObj
	newLoan := LoanObj{
		Loan_id:           "Loan-123",
		Name:              req.Name,
		KTP:               req.KTP,
		Loan_amount:       req.Loan_amount,
		Loan_period_month: req.Loan_period_month,
		Loan_purpose:      req.Loan_purpose,
		DOB:               req.DOB,
		Gender:            req.Gender,
	}

	//Opening the JSON data
	savedLoanJson, _ := os.Open("helpers/dummyLoanData.json")
	fmt.Println("The File is opened successfully...")
	defer savedLoanJson.Close()
	byteValue, _ := ioutil.ReadAll(savedLoanJson)
	json.Unmarshal(byteValue, &loans)
	// for i := 0; i < len(loans); i++ {
	// 	fmt.Println(loans[i].Name)
	// }

	//Merging the current Loan List with the Incoming Loan
	mergedLoanArr := append(loans, newLoan)
	fmt.Println("loans: ", loans)
	fmt.Println("loans.loans: ", loans)
	fmt.Println("requestBody: ", newLoan)
	fmt.Println("mergedArr:", mergedLoanArr)
	file, _ := json.MarshalIndent(mergedLoanArr, "", " ")
	_ = ioutil.WriteFile("helpers/dummyLoanData.json", file, 0644)

	c.JSON(http.StatusOK, gin.H{"data": newLoan})
}
