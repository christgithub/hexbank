package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//go:generate mockgen -source=balance_handler.go -destination=../mocks/api/balance_handler.go
type BalanceManager interface {
	Retrieve(accountNumber string) float32
}

type BalanceHandler struct{
	Manager BalanceManager
}

func (b BalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountNumber := params["number"]
	balance := b.Manager.Retrieve(accountNumber)
	strBalance := fmt.Sprintf("%.2f", balance)
	w.Write([]byte(strBalance))
}
