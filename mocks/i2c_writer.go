package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type I2cWriter struct {
	mock.Mock
}

// I2cWrite provides a mock function with given fields: address, buf
func (_m *I2cWriter) I2cWrite(address int, buf []byte) error {
	ret := _m.Called(address, buf)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []byte) error); ok {
		r0 = rf(address, buf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
