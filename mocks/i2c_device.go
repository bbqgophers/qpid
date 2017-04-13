package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type I2cDevice struct {
	mock.Mock
}

// SetAddress provides a mock function with given fields: _a0
func (_m *I2cDevice) SetAddress(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
