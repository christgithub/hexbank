package gordontester

import "github.com/christgithub/hexbank/bank"

type GordonTester struct{
	BalanceManager bank.BalanceManager
}

func New(balanceManager bank.BalanceManager) GordonTester {
	return GordonTester{BalanceManager: balanceManager}
}
