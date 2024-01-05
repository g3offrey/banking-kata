package banking

import "errors"

var ErrNoMoney = errors.New("Can't withdraw")
var ErrAccountNotFound = errors.New("Account not found")
var ErrNegativeAmount = errors.New("Invalid amount (negative)")
