package controllers

import (
	"onlineLoan/entity"
	"onlineLoan/services"

	"github.com/gin-gonic/gin"
)

// OnlineLoanController ...
type OnlineLoanController interface {
	AddApplication(ctx *gin.Context) int
	GetLoanInfo(string) (entity.LoanType, error)
}

type controller struct {
	services services.OnlineLoanServices
}

//New OnlineLoanController ...
func New(services services.OnlineLoanServices) OnlineLoanController {
	return &controller{
		services: services,
	}
}

func (c *controller) AddApplication(ctx *gin.Context) int {
	return c.services.AddApplication(entity.Application{})
}
func (c *controller) GetLoanInfo(id string) (entity.LoanType, error) {

	ent, err := c.services.GetLoanInfo(id)
	return ent, err

}
