package storage

import "github.com/g3offrey/banking-kata/banking"

type AccountRepositoryInMemory struct {
	accounts map[string]banking.Account
}

func NewAccountRepositoryInMemory() *AccountRepositoryInMemory {
	return &AccountRepositoryInMemory{
		accounts: make(map[string]banking.Account),
	}
}

func (r *AccountRepositoryInMemory) FindAll() []banking.Account {
	accounts := make([]banking.Account, 0, len(r.accounts))

	for _, account := range r.accounts {
		accounts = append(accounts, account)
	}

	return accounts
}

func (r *AccountRepositoryInMemory) Find(owner string) (result banking.Account, found bool) {
	result, found = r.accounts[owner]
	return result, found
}

func (r *AccountRepositoryInMemory) Save(account banking.Account) error {
	r.accounts[account.Owner] = account

	return nil
}
