package banking

type AccountRepository interface {
	FindAll() []Account
	Find(owner string) (result Account, found bool)
	Save(account Account) error
}
