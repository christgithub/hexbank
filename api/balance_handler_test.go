package api_test

import (
	"github.com/christgithub/hexbank/api"
	mock_api "github.com/christgithub/hexbank/mocks/api"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("BalanceHandler", func() {
	When("handler requests a balance", func() {

		var mockController *gomock.Controller
		var mockBalanceManager *mock_api.MockBalanceManager

		mockController = gomock.NewController(GinkgoT())
		mockBalanceManager = mock_api.NewMockBalanceManager(mockController)


		It("gets a balance returns a http ok 200", func() {
			accountNumber := "123"
			mockBalanceManager.EXPECT().Retrieve(accountNumber).Return(float32(2345.45))

			balanceHandler := api.BalanceHandler{ Manager: mockBalanceManager}
			request := httptest.NewRequest("GET", "/account/{number}/balance", nil)
			request = mux.SetURLVars(request, map[string]string{"number":accountNumber})
			response := httptest.NewRecorder()
			balanceHandler.ServeHTTP(response, request)

			body, _ := ioutil.ReadAll(response.Body)

			Expect(string(body)).To(Equal("2345.45"))
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})
})
