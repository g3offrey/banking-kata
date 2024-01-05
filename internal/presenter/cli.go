package presenter

import (
	"fmt"

	"github.com/g3offrey/banking-kata/banking"
)

type CliAccountPresenter struct{}

func NewCliAccountPresenter() *CliAccountPresenter {
	return &CliAccountPresenter{}
}

func (c CliAccountPresenter) ShowAccounts(accounts []banking.AccountResponseModel) {
	for _, account := range accounts {
		fmt.Printf("- Account: %s, Balance: %d$\n", account.Owner, account.Balance)
	}
}

func (c CliAccountPresenter) ShowAccount(account banking.AccountResponseModel) {
	fmt.Printf("Account: %s, Balance: %d$\n", account.Owner, account.Balance)
}

func (c CliAccountPresenter) ShowError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}
