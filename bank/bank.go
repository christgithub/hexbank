package bank

type Bank struct{}

type Account struct {}

func New() Bank {
	return Bank{}
}

type BalanceManager interface {
	Retrieve(accountNumber string) float32
}

type AccountCreator interface {
	CreateAccount() (Account, error)
}

func (b Bank) Retrieve(accountNumber string) float32 {
	switch (accountNumber) {
	case "123":
		return 2345.45
	default:
		return 0
	}
}

func (b Bank) CreateAccount() (Account, error) {
	//Interface call
	return Account{}, nil
}