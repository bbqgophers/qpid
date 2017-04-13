package mocks

import "github.com/briandowns/qpid/log"
import "github.com/stretchr/testify/mock"

import "github.com/bbqgophers/qpid"

// *DO NOT EDIT* Auto-generated via mockery

type Logger struct {
	mock.Mock
}

// Listen provides a mock function with given fields: n
func (_m *Logger) Listen(n chan qpid.Notification) {
	_m.Called(n)
}
