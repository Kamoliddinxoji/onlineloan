package services

import (
	"database/sql"
	"errors"
	"fmt"
	"onlineLoan/entity"
	"time"

	// _ ....
	_ "github.com/lib/pq"
)

//OnlineLoanServices ...
type OnlineLoanServices interface {
	AddApplication(entity.Application) int
	GetLoanInfo(string) (entity.LoanType, error)
}

type onlineLoanServices struct {
}

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "xoji"
	dbname   = "onlineloan"
)

//New OnlineLoanServices
func New() OnlineLoanServices {
	return &onlineLoanServices{}

}

func (service *onlineLoanServices) AddApplication(entity.Application) int {

	return 1
}

func (service *onlineLoanServices) GetLoanInfo(id string) (entity.LoanType, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	defer db.Close()

	sql := fmt.Sprintf("Select * from loantype where id =%s", id)

	row, errr := db.Query(sql)
	var r entity.LoanType
	fmt.Println(row)
	fmt.Println(errr)
	if errr == nil {
		for row.Next() {
			err = row.Scan(&r.ID, &r.Name, &r.MaxAmount, &r.MinAmount, &r.Percent, &r.CalcType)
		}
		if r.ID == 0 {
			fmt.Println(err)
			return r, errors.New("error")
		}
		return r, nil
	}
	return r, errr
}
