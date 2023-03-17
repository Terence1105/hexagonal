package out

import (
	"strconv"

	"github.com/Terence1105/hexagonal/account/application/port/out"
	"github.com/Terence1105/hexagonal/account/domain"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

type accountLevelDBAdapter struct {
	db *leveldb.DB
}

func NewAccountLevelDBAdapter() out.ITransferPort {
	db := &leveldb.DB{}

	a1 := domain.Account{
		Name:    "a1",
		Balance: 100,
	}
	a2 := domain.Account{
		Name:    "a2",
		Balance: 100,
	}
	db.Put([]byte(a1.Name), []byte(strconv.Itoa(a1.Balance)), nil)
	db.Put([]byte(a2.Name), []byte(strconv.Itoa(a2.Balance)), nil)

	return &accountLevelDBAdapter{
		db: db,
	}
}

func (a *accountLevelDBAdapter) GetAccount(name string) (*domain.Account, error) {
	balance, err := a.db.Get([]byte(name), nil)
	if err != nil {
		return nil, errors.Wrap(err, "account not found")
	}

	b, err := strconv.Atoi(string(balance))
	if err != nil {
		return nil, errors.Wrap(err, "balance convert failed")
	}

	return &domain.Account{
		Name:    name,
		Balance: b,
	}, nil
}

func (a *accountLevelDBAdapter) Transfer(fromAccountName, ToAccountName string, amount int) error {

	faBalance, err := a.db.Get([]byte(fromAccountName), nil)
	if err != nil {
		return errors.Wrap(err, "account not found")
	}

	fab, err := strconv.Atoi(string(faBalance))
	if err != nil {
		return errors.Wrap(err, "balance convert failed")
	}

	err = a.db.Put([]byte(fromAccountName), []byte(strconv.Itoa(fab-amount)), nil)
	if err != nil {
		return errors.Wrap(err, "update account failed")
	}

	taBalance, err := a.db.Get([]byte(ToAccountName), nil)
	if err != nil {
		return errors.Wrap(err, "account not found")
	}

	tab, err := strconv.Atoi(string(taBalance))
	if err != nil {
		return errors.Wrap(err, "balance convert failed")
	}

	err = a.db.Put([]byte(ToAccountName), []byte(strconv.Itoa(tab+amount)), nil)
	if err != nil {
		return errors.Wrap(err, "update account failed")
	}

	return nil
}
