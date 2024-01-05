package main

import (
	"fmt"
	"net/http"

	"github.com/g3offrey/banking-kata/banking"
	"github.com/g3offrey/banking-kata/internal/presenter"
	"github.com/g3offrey/banking-kata/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	accountRepository := storage.NewAccountRepositoryInMemory()

	r.Get("/account", func(w http.ResponseWriter, r *http.Request) {
		useCases := banking.NewUseCases(accountRepository, presenter.NewHTTPAccountPresenter(w, r))
		useCases.ListAccounts()
	})

	r.Post("/account/{account}/deposit", func(w http.ResponseWriter, r *http.Request) {
		useCases := banking.NewUseCases(accountRepository, presenter.NewHTTPAccountPresenter(w, r))
		var owner = chi.URLParam(r, "account")
		var request = DepositAccountRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		useCases.Deposit(banking.DepositRequestModel{
			Owner:  owner,
			Amount: request.Amount,
		})
	})

	r.Post("/account/{account}/withdraw", func(w http.ResponseWriter, r *http.Request) {
		useCases := banking.NewUseCases(accountRepository, presenter.NewHTTPAccountPresenter(w, r))
		var owner = chi.URLParam(r, "account")
		var request = WithdrawAccountRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		useCases.Withdraw(banking.WithdrawRequestModel{
			Owner:  owner,
			Amount: request.Amount,
		})
	})

	r.Post("/account/{account}/transfer", func(w http.ResponseWriter, r *http.Request) {
		useCases := banking.NewUseCases(accountRepository, presenter.NewHTTPAccountPresenter(w, r))
		var from = chi.URLParam(r, "account")
		var request = TransferAccountRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		useCases.Transfer(banking.TransferRequestModel{
			From:   from,
			To:     request.To,
			Amount: request.Amount,
		})
	})

	r.Post("/account", func(w http.ResponseWriter, r *http.Request) {
		useCases := banking.NewUseCases(accountRepository, presenter.NewHTTPAccountPresenter(w, r))
		var request = CreateAccountRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		useCases.CreateAccount(banking.CreateAccountRequestModel{
			Owner: request.Owner,
		})
	})

	fmt.Println("Starting server on port 3000")
	_ = http.ListenAndServe(":3000", r)
}
