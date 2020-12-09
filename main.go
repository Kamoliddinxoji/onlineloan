package main

import (
	"fmt"
	"onlineLoan/controllers"
	"onlineLoan/entity"
	"onlineLoan/services"

	"github.com/gin-gonic/gin"
)

var (
	loanService    services.OnlineLoanServices      = services.New()
	loanController controllers.OnlineLoanController = controllers.New(loanService)
)

func main() {
	fmt.Println("kamoliddin")
	r := gin.Default()
	r.GET("/loans/:id", func(c *gin.Context) {
		ent, err := loanController.GetLoanInfo(c.Param("id"))
		if err != nil {
			c.Header("Content-Type", "application/json")
			c.JSON(200, entity.Response{"Error", -1, "Failed", ent})
		} else {
			c.Header("Content-Type", "application/json")
			c.JSON(200, entity.Response{"OK", 0, "Succes", ent})
		}

	})

	r.POST("/loans", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(200, loanController.AddApplication(c))

	})

	r.Run()

}
