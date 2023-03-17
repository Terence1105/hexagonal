package in

import (
	"strconv"

	"github.com/Terence1105/hexagonal/account/application/port/in"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	transferUseCase in.ITransferUseCase
}

func (g *GinAdapter) Transfer(c *gin.Context) {
	fromAccountName := c.Param("fromAccountName")
	toAccountName := c.Param("toAccountName")
	amount := c.Param("amount")

	a, err := strconv.Atoi(amount)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	transferCommand, err := in.NewTransferCommand(fromAccountName, toAccountName, a)
	err = g.transferUseCase.TransferUseCase(transferCommand)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}
}

func (g *GinAdapter) Run() {
	r := gin.Default()

	r.GET("/transfer/:fromAccountName/:toAccountName/:amount", g.Transfer)

	r.Run(":8080")
}

func NewGinAdapter(transferUseCase in.ITransferUseCase) GinAdapter {
	return GinAdapter{
		transferUseCase: transferUseCase,
	}
}
