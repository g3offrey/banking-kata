package banking_test

import (
	"testing"

	"github.com/g3offrey/banking-kata/banking"
	"github.com/stretchr/testify/assert"
)

func TestEmptyAccount(t *testing.T) {
	account := banking.NewAccount("john")

	assert.Equal(t, "john", account.Owner)
	assert.Equal(t, 0, account.Balance)
}

func TestDepositMoney(t *testing.T) {
	account := banking.NewAccount("john")

	err := account.Deposit(10)

	assert.Nil(t, err)
	assert.Equal(t, 10, account.Balance)
}

func TestWithdrawMoney(t *testing.T) {
	account := banking.NewAccount("john")
	_ = account.Deposit(10)

	err := account.Withdraw(5)

	assert.Nil(t, err)
	assert.Equal(t, 5, account.Balance)
}

func TestWithdrawTooMuchMoney(t *testing.T) {
	account := banking.NewAccount("john")
	_ = account.Deposit(10)

	err := account.Withdraw(15)

	assert.Equal(t, banking.ErrNoMoney, err)
	assert.Equal(t, 10, account.Balance)
}

func TestWithdrawNegativeAmount(t *testing.T) {
	account := banking.NewAccount("john")

	err := account.Withdraw(-5)

	assert.Equal(t, banking.ErrNegativeAmount, err)
	assert.Equal(t, 0, account.Balance)
}

func TestDepositNegativeAmount(t *testing.T) {
	account := banking.NewAccount("john")

	err := account.Deposit(-5)

	assert.Equal(t, banking.ErrNegativeAmount, err)
	assert.Equal(t, 0, account.Balance)
}
