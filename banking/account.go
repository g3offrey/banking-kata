package banking

type Account struct {
	Owner   string
	Balance int
}

func NewAccount(owner string) *Account {
	account := Account{Owner: owner, Balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) error {
	if amount < 0 {
		return ErrNegativeAmount
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return ErrNegativeAmount
	}

	if a.Balance < amount {
		return ErrNoMoney
	}

	a.Balance -= amount
	return nil
}
