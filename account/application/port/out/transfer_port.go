package out

import "github.com/Terence1105/hexagonal/account/domain"

type ITransferPort interface {
	GetAccount(name string) (*domain.Account, error)
	Transfer(fromAccountName, ToAccountName string, amount int) error
}
