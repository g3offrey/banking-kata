package banking

type AccountPresenter interface {
	ShowAccounts(accounts []AccountResponseModel)
	ShowAccount(account AccountResponseModel)
	ShowError(err error)
}
