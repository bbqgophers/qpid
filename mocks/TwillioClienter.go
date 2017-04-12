package mocks

import "github.com/stretchr/testify/mock"

import "github.com/bbqgophers/qpid"

type TwillioClienter struct {
	mock.Mock
}

// Listen provides a mock function with given fields: a
func (_m *TwillioClienter) Listen(a chan qpid.Notification) {
	_m.Called(a)
}
