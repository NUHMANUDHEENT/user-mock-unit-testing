package database

import (
    "github.com/stretchr/testify/mock"
)

type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) CreateUser(username, password string) error {
    args := m.Called(username, password)
    return args.Error(0)
}
type Database interface {
    CreateUser(username, password string) error
}
 