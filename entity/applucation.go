package entity

import (
	"time"
)

//Application ...
type Application struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	LoanTypeID int       `json:"loantype_id"`
	LoanAmount float64   `json:"loan_amount"`
	LoanTerm   int       `json:"loan_term"`
	BeginDate  time.Time `json:"begin_date"`
	EndDate    time.Time `json:"pay_type"`
	PayType    string    `json:"end_date"`
	WritedDate time.Time `json:"writed_date"`
}
