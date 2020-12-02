package api

//Inbound adaptor
import "github.com/gorilla/mux"

type Router struct {
	Client *mux.Router
	Bank BalanceManager
}

func New(client *mux.Router, bank BalanceManager) Router {
	return Router{
		Client: client,
		Bank: bank,
	}
}

func(r Router) setUpRoutes() {
	r.Client.Handle("/account/{number}/balance", BalanceHandler{Manager: r.Bank})
}

