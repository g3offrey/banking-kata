package banking

type UseCases struct {
	repository AccountRepository
	presenter  AccountPresenter
}

func NewUseCases(repository AccountRepository, presenter AccountPresenter) *UseCases {
	return &UseCases{
		repository: repository,
		presenter:  presenter,
	}
}

type AccountResponseModel struct {
	Owner   string
	Balance int
}

func AccountResponseModelFromEntity(account *Account) AccountResponseModel {
	return AccountResponseModel{
		Owner:   account.Owner,
		Balance: account.Balance,
	}
}

type CreateAccountRequestModel struct {
	Owner string
}

func (u *UseCases) CreateAccount(model CreateAccountRequestModel) {
	account := NewAccount(model.Owner)

	err := u.repository.Save(*account)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	u.presenter.ShowAccount(AccountResponseModelFromEntity(account))
}

type DepositRequestModel struct {
	Owner  string
	Amount int
}

func (u *UseCases) Deposit(model DepositRequestModel) {
	account, found := u.repository.Find(model.Owner)
	if !found {
		u.presenter.ShowError(ErrAccountNotFound)
		return
	}

	err := account.Deposit(model.Amount)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	err = u.repository.Save(account)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	u.presenter.ShowAccount(AccountResponseModelFromEntity(&account))
}

type WithdrawRequestModel struct {
	Owner  string
	Amount int
}

func (u *UseCases) Withdraw(model WithdrawRequestModel) {
	account, found := u.repository.Find(model.Owner)
	if !found {
		u.presenter.ShowError(ErrAccountNotFound)
		return
	}

	err := account.Withdraw(model.Amount)
	if err != nil {
		u.presenter.ShowError(err)
	}

	err = u.repository.Save(account)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	u.presenter.ShowAccount(AccountResponseModelFromEntity(&account))
}

type TransferRequestModel struct {
	From   string
	To     string
	Amount int
}

func (u *UseCases) Transfer(model TransferRequestModel) {
	fromAccount, found := u.repository.Find(model.From)
	if !found {
		u.presenter.ShowError(ErrAccountNotFound)
		return
	}
	toAccount, found := u.repository.Find(model.To)
	if !found {
		u.presenter.ShowError(ErrAccountNotFound)
		return
	}

	err := fromAccount.Withdraw(model.Amount)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}
	err = toAccount.Deposit(model.Amount)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	err = u.repository.Save(fromAccount)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}
	err = u.repository.Save(toAccount)
	if err != nil {
		u.presenter.ShowError(err)
		return
	}

	u.presenter.ShowAccounts(
		[]AccountResponseModel{
			AccountResponseModelFromEntity(&fromAccount),
			AccountResponseModelFromEntity(&toAccount),
		},
	)
}

func (u *UseCases) ListAccounts() {
	accounts := u.repository.FindAll()

	responseModels := make([]AccountResponseModel, len(accounts))
	for i, account := range accounts {
		responseModels[i] = AccountResponseModelFromEntity(&account)
	}

	u.presenter.ShowAccounts(responseModels)
}
