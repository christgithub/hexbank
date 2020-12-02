package hexbank

import (
	"github.com/christgithub/hexbank/api"
	"github.com/christgithub/hexbank/bank"
	"github.com/gorilla/mux"
)

func main() {
	bank := bank.New()
	client := mux.NewRouter()
	router := api.New(client, bank)
}
