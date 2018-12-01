package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caominhtri91/micro/accountservice/dbclient"
	"github.com/caominhtri91/micro/accountservice/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest(http.MethodGet, "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be the 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetAccount(t *testing.T) {
	// Create a mock instance that implements the IBoltClient interface
	mockRepo := &dbclient.MockBoltClient{}

	mockRepo.On("QueryAccount", "123").Return(model.Account{ID: "123", Name: "Tri dap chai"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{ID: "456", Name: "Tri tai gioi"}, nil)

	DBClient = mockRepo

	Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest(http.MethodGet, "/accounts/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				So(account.ID, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Tri dap chai")
			})
		})
	})
}
