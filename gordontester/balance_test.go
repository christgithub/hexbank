package gordontester_test

import (
	"github.com/christgithub/hexbank/bank"
	gordontester "github.com/christgithub/hexbank/gordontester"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BalanceManager", func() {
	When("Primary Actor Retrieve", func() {
		It("gets a balance", func() {
			bank := bank.New()
			tester := gordontester.New(bank)

			accountNumber := "123"
			balance := tester.BalanceManager.Retrieve(accountNumber)
			Expect(balance).To(Equal(float32(2345.45)))
		})
	})
})
