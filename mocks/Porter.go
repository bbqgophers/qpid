package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type Porter struct {
	mock.Mock
}

// Port provides a mock function with given fields:
func (_m *Porter) Port() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
