package loans

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

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

func FindFromArrStr(arr []string, params string) (found bool) {
	for i := range arr {
		if arr[i] == params {
			found = true
			break
		}
	}
	return
}

func CreateLoan(c *gin.Context) {
	//Initialize enum
	loan_purpose_enum := [7]string{"vacation", "renovation", "electronics", "wedding", "rent", "car", "investment"}
	gender_enum := [2]string{"male", "female"}

	//Getting the request data
	var req LoanObj
	if err := c.BindJSON(&req); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	//Get the current directory
	// path, _ := os.Getwd()
	// fmt.Println(path)

	//Generate Loan_ID
	currentTime := time.Now().Format("02-01-2006")
	reformatedCurrentTime := currentTime[0:2] + currentTime[3:5] + currentTime[8:10]
	unixTime := strconv.FormatInt(time.Now().Unix(), 10)

	var loan_id string = "LOAN-" + reformatedCurrentTime + "-" + unixTime[len(unixTime)-4:]

	var loans []LoanObj
	newLoan := LoanObj{
		Loan_id:           loan_id,
		Name:              req.Name,
		KTP:               req.KTP,
		Loan_amount:       req.Loan_amount,
		Loan_period_month: req.Loan_period_month,
		Loan_purpose:      req.Loan_purpose,
		DOB:               req.DOB,
		Gender:            req.Gender,
	}

	//Validating the request
	var nameRegex, _ = regexp.Compile(`^[a-zA-Z ]{4,}$`)
	var nameValid = nameRegex.MatchString(newLoan.Name)
	if !nameValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name is invalid"})
		return
	}

	splittedName := strings.Fields(newLoan.Name)
	if len(splittedName) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name is invalid. Name should have minimum two words"})
		return
	}

	gender_valid := FindFromArrStr(gender_enum[:], newLoan.Gender)
	if !gender_valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gender is invalid. Only accept these value: male, female"})
		return
	}

	var ktpRegex, _ = regexp.Compile(`^[0-9]{16}$`)
	var ktpFormatValid = ktpRegex.MatchString(newLoan.KTP)
	if !ktpFormatValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The KTP Format is invalid"})
		return
	}

	var dateRegex, _ = regexp.Compile(`^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`)
	var dateFormatValid = dateRegex.MatchString(newLoan.DOB)
	if !dateFormatValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The DOB is Invalid"})
		return
	}

	var nikBirthDate = newLoan.KTP[6:12]
	var birthDate = newLoan.DOB[0:2]

	if newLoan.Gender == "female" {
		var birthDateInt, _ = strconv.Atoi(birthDate)
		birthDateInt += 40
		birthDate = strconv.Itoa(birthDateInt)
	}
	var reformatedDOB = birthDate + newLoan.DOB[3:5] + newLoan.DOB[8:10]
	fmt.Println("reformatedDOB: ", reformatedDOB)
	if nikBirthDate != reformatedDOB {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The DOB and NIK is not match"})
		return
	}

	if newLoan.Loan_amount < 1000000 || newLoan.Loan_amount > 10000000 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Loan Amount is invalid. Loan Amount should between 1.000.000 and 10.000.000"})
		return
	}

	if newLoan.Loan_period_month > 240 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Maximum loan period is 240"})
		return
	}

	loan_purpose_valid := FindFromArrStr(loan_purpose_enum[:], newLoan.Loan_purpose)
	if !loan_purpose_valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Loan Purpose is invalid. Only accept these value: vacation, renovation, electronics, wedding, rent, car, investment"})
		return
	}

	//Opening the JSON data
	savedLoanJson, _ := os.Open("helpers/dummyLoanData.json")
	fmt.Println("The File is opened successfully...")
	defer savedLoanJson.Close()
	byteValue, _ := ioutil.ReadAll(savedLoanJson)
	json.Unmarshal(byteValue, &loans)

	//Merging the current Loan List with the Incoming Loan
	mergedLoanArr := append(loans, newLoan)
	file, _ := json.MarshalIndent(mergedLoanArr, "", " ")
	_ = ioutil.WriteFile("helpers/dummyLoanData.json", file, 0644)

	c.JSON(http.StatusOK, gin.H{"data": newLoan})
}

func FindLoanById(c *gin.Context) {
	//Getting the request data
	requestedLoanId := c.Params.ByName("loan_id")

	//Validating the input
	var loanIdRegex, _ = regexp.Compile(`^LOAN-[0-9]{6}-[0-9]{4}$`)
	var loanIdValid = loanIdRegex.MatchString(requestedLoanId)
	if !loanIdValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Loan ID is Invalid"})
		return
	}

	//Opening the JSON data
	var loans []LoanObj
	savedLoanJson, _ := os.Open("helpers/dummyLoanData.json")
	fmt.Println("The File is opened successfully...")
	defer savedLoanJson.Close()
	byteValue, _ := ioutil.ReadAll(savedLoanJson)
	json.Unmarshal(byteValue, &loans)
	var loanIdx int = -1
	for i := range loans {
		if loans[i].Loan_id == requestedLoanId {
			loanIdx = i
		}
	}
	if loanIdx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Loan ID is Not Found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": loans[loanIdx]})
	}
}

func FindLoadByKTP(c *gin.Context) {
	//Getting the request data
	requestedKTP := c.Params.ByName("ktp")

	//Validating the input
	var ktpRegex, _ = regexp.Compile(`^[0-9]{16}$`)
	var ktpFormatValid = ktpRegex.MatchString(requestedKTP)
	if !ktpFormatValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The KTP Format is invalid"})
		return
	}
	//Opening the JSON data
	var loans []LoanObj
	savedLoanJson, _ := os.Open("helpers/dummyLoanData.json")
	fmt.Println("The File is opened successfully...")
	defer savedLoanJson.Close()
	byteValue, _ := ioutil.ReadAll(savedLoanJson)
	json.Unmarshal(byteValue, &loans)

	var foundLoanList []LoanObj
	for i := range loans {
		if loans[i].KTP == requestedKTP {
			foundLoanList = append(foundLoanList, loans[i])
		}
	}
	if len(foundLoanList) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "There's no loan with this KTP"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": foundLoanList})
	}

}
