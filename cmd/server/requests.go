package main

import "net/http"

type CreateAccountRequest struct {
	Owner string `json:"owner"`
}

func (c CreateAccountRequest) Bind(r *http.Request) error {
	return nil
}

type DepositAccountRequest struct {
	Amount int `json:"amount"`
}

func (d DepositAccountRequest) Bind(r *http.Request) error {
	return nil
}

type WithdrawAccountRequest struct {
	Amount int `json:"amount"`
}

func (w WithdrawAccountRequest) Bind(r *http.Request) error {
	return nil
}

type TransferAccountRequest struct {
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

func (t TransferAccountRequest) Bind(r *http.Request) error {
	return nil
}
