package service

import (
	"github.com/Terence1105/hexagonal/account/application/port/in"
	"github.com/Terence1105/hexagonal/account/application/port/out"
	"github.com/pkg/errors"
)

type transferService struct {
	transferPort out.ITransferPort
}

func (t *transferService) TransferUseCase(transferCommand *in.TransferCommand) error {

	to, err := t.transferPort.GetAccount(transferCommand.FromAccountName)
	if err != nil {
		return errors.Wrap(err, "GetAccountPort error")
	}

	if to.Balance < transferCommand.Amount {
		return errors.New("insufficient balance")
	}

	err = t.transferPort.Transfer(transferCommand.FromAccountName, transferCommand.ToAccountName, transferCommand.Amount)
	if err != nil {
		return errors.Wrap(err, "TransferPort error")
	}

	return nil
}

func NewTransferService(transferPort out.ITransferPort) in.ITransferUseCase {
	return &transferService{
		transferPort: transferPort,
	}
}
