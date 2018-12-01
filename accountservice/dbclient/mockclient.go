package dbclient

import (
	"github.com/caominhtri91/micro/accountservice/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock implementation of a datastore
// client for testing perpose. Instead of the bolt.DB pointer
// we're just putting a generic mock object
type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) OpenBoltDB() {
	panic("not implemented")
}

// QueryAccount From here, we'll declare three functions that makes
// our MockBoltClient fulfill the interface IBoltClient
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) Seed() {
	panic("not implemented")
}
