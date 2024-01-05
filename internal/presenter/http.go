package presenter

import (
	"net/http"

	"github.com/g3offrey/banking-kata/banking"
	"github.com/go-chi/render"
)

type AccountResponse struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

func (a AccountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type HTTPAccountPresenter struct {
	w http.ResponseWriter
	r *http.Request
}

func NewHTTPAccountPresenter(w http.ResponseWriter, r *http.Request) *HTTPAccountPresenter {
	return &HTTPAccountPresenter{
		w: w,
		r: r,
	}
}

func (h HTTPAccountPresenter) ShowAccounts(accounts []banking.AccountResponseModel) {
	responses := make([]render.Renderer, 0, len(accounts))
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			Owner:   account.Owner,
			Balance: account.Balance,
		})
	}

	_ = render.RenderList(h.w, h.r, responses)
}

func (h HTTPAccountPresenter) ShowAccount(account banking.AccountResponseModel) {
	_ = render.Render(h.w, h.r, AccountResponse{
		Owner:   account.Owner,
		Balance: account.Balance,
	})
}

func (h HTTPAccountPresenter) ShowError(err error) {
	render.Status(h.r, http.StatusBadRequest)
	render.PlainText(h.w, h.r, err.Error())
}
