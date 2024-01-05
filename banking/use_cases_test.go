package banking_test

import (
	"testing"

	"github.com/g3offrey/banking-kata/banking"
	"github.com/g3offrey/banking-kata/internal/storage"
	"github.com/stretchr/testify/mock"
)

type PresenterSpy struct {
	mock.Mock
}

func (p *PresenterSpy) ShowAccount(account banking.AccountResponseModel) {
	_ = p.Called(account)
}

func (p *PresenterSpy) ShowAccounts(accounts []banking.AccountResponseModel) {
	_ = p.Called(accounts)
}

func (p *PresenterSpy) ShowError(err error) {
	_ = p.Called(err)
}

func TestCreateAccount(t *testing.T) {
	presenter := PresenterSpy{}
	presenter.On("ShowAccount", mock.Anything)
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), &presenter)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "John Doe",
	})

	presenter.AssertCalled(t, "ShowAccount", banking.AccountResponseModel{
		Owner: "John Doe",
	})
}

func TestDeposit(t *testing.T) {
	presenter := PresenterSpy{}
	presenter.On("ShowAccount", mock.Anything)
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), &presenter)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "John Doe",
	})
	uc.Deposit(banking.DepositRequestModel{
		Owner:  "John Doe",
		Amount: 100,
	})

	presenter.AssertCalled(t, "ShowAccount", banking.AccountResponseModel{
		Owner:   "John Doe",
		Balance: 100,
	})
}

func TestWithdraw(t *testing.T) {
	presenter := PresenterSpy{}
	presenter.On("ShowAccount", mock.Anything)
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), &presenter)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "John Doe",
	})
	uc.Deposit(banking.DepositRequestModel{
		Owner:  "John Doe",
		Amount: 100,
	})
	uc.Withdraw(banking.WithdrawRequestModel{
		Owner:  "John Doe",
		Amount: 50,
	})

	presenter.AssertCalled(t, "ShowAccount", banking.AccountResponseModel{
		Owner:   "John Doe",
		Balance: 50,
	})
}

func TestTransfer(t *testing.T) {
	presenter := PresenterSpy{}
	presenter.On("ShowAccount", mock.Anything)
	presenter.On("ShowAccounts", mock.Anything)
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), &presenter)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "John Doe",
	})
	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "Jane Doe",
	})
	uc.Deposit(banking.DepositRequestModel{
		Owner:  "John Doe",
		Amount: 100,
	})
	uc.Transfer(banking.TransferRequestModel{
		From:   "John Doe",
		To:     "Jane Doe",
		Amount: 50,
	})

	presenter.AssertCalled(t, "ShowAccounts", []banking.AccountResponseModel{
		{
			Owner:   "John Doe",
			Balance: 50,
		},
		{
			Owner:   "Jane Doe",
			Balance: 50,
		},
	})
}

func TestListAccounts(t *testing.T) {
	presenter := PresenterSpy{}
	presenter.On("ShowAccount", mock.Anything)
	presenter.On("ShowAccounts", mock.Anything)
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), &presenter)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "John Doe",
	})
	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: "Jane Doe",
	})
	uc.ListAccounts()

	presenter.AssertCalled(t, "ShowAccounts", []banking.AccountResponseModel{
		{
			Owner:   "John Doe",
			Balance: 0,
		},
		{
			Owner:   "Jane Doe",
			Balance: 0,
		},
	})
}
