package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type PwmWriter struct {
	mock.Mock
}

// PwmWrite provides a mock function with given fields: _a0, _a1
func (_m *PwmWriter) PwmWrite(_a0 string, _a1 byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
