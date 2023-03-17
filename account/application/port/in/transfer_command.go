package in

import (
	"github.com/Terence1105/hexagonal/utils"
	"github.com/pkg/errors"
)

type TransferCommand struct {
	FromAccountName string `validate:"min=4, required"`
	ToAccountName   string `validate:"min=4, required"`
	Amount          int    `validate:"gte=1, required"`
}

func NewTransferCommand(fromAccountName string, toAccountName string, amount int) (*TransferCommand, error) {
	tran := TransferCommand{
		ToAccountName:   toAccountName,
		FromAccountName: fromAccountName,
		Amount:          amount,
	}

	validate := utils.GetValidator()
	if err := validate.Struct(tran); err != nil {
		return nil, errors.Wrap(err, "TransferCommand validation error")
	}

	return &tran, nil
}
