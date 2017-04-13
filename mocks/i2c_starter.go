package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type I2cStarter struct {
	mock.Mock
}

// I2cStart provides a mock function with given fields: address
func (_m *I2cStarter) I2cStart(address int) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
