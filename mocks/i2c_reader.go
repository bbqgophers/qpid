package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type I2cReader struct {
	mock.Mock
}

// I2cRead provides a mock function with given fields: address, len
func (_m *I2cReader) I2cRead(address int, len int) ([]byte, error) {
	ret := _m.Called(address, len)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(int, int) []byte); ok {
		r0 = rf(address, len)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(address, len)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
