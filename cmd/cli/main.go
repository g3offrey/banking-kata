package main

import (
	"fmt"
	"os"

	"github.com/g3offrey/banking-kata/banking"
	"github.com/g3offrey/banking-kata/internal/presenter"
	"github.com/g3offrey/banking-kata/internal/storage"
)

func main() {
	uc := banking.NewUseCases(storage.NewAccountRepositoryInMemory(), presenter.NewCliAccountPresenter())

	fmt.Println("Hello, welcome to the banking kata!")

	for {
		fmt.Println("Press C to create an account, W to withdraw, D to deposit, T to transfer, L to list accounts or Q to quit.")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "L":
			listAccounts(uc)
		case "C":
			createAccount(uc)
		case "W":
			withdraw(uc)
		case "D":
			deposit(uc)
		case "T":
			transfer(uc)
		case "Q":
			os.Exit(0)
		}
	}
}

func createAccount(uc *banking.UseCases) {
	fmt.Println("Please enter your name:")
	var name string
	fmt.Scanln(&name)

	uc.CreateAccount(banking.CreateAccountRequestModel{
		Owner: name,
	})
}

func deposit(uc *banking.UseCases) {
	fmt.Println("Please enter your name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Please enter the amount you want to deposit:")
	var amount int
	fmt.Scanln(&amount)

	uc.Deposit(banking.DepositRequestModel{
		Owner:  name,
		Amount: amount,
	})
}

func withdraw(uc *banking.UseCases) {
	fmt.Println("Please enter your name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Please enter the amount you want to withdraw:")
	var amount int
	fmt.Scanln(&amount)

	uc.Withdraw(banking.WithdrawRequestModel{
		Owner:  name,
		Amount: amount,
	})
}

func listAccounts(uc *banking.UseCases) {
	fmt.Println("Here are all the accounts:")
	uc.ListAccounts()
}

func transfer(uc *banking.UseCases) {
	fmt.Println("Please enter the name of the account you want to transfer from:")
	var from string
	fmt.Scanln(&from)

	fmt.Println("Please enter the name of the account you want to transfer to:")
	var to string
	fmt.Scanln(&to)

	fmt.Println("Please enter the amount you want to transfer:")
	var amount int
	fmt.Scanln(&amount)

	uc.Transfer(banking.TransferRequestModel{
		From:   from,
		To:     to,
		Amount: amount,
	})
}
