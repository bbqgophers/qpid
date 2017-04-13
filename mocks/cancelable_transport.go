package mocks

import "github.com/stretchr/testify/mock"

import "net/http"

// *DO NOT EDIT* Auto-generated via mockery

type CancelableTransport struct {
	mock.Mock
}

// CancelRequest provides a mock function with given fields: req
func (_m *CancelableTransport) CancelRequest(req *http.Request) {
	_m.Called(req)
}
