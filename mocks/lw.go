package mocks

import "github.com/stretchr/testify/mock"

import "C"

type lw struct {
	mock.Mock
}

// digitalWrite provides a mock function with given fields: _a0, _a1
func (_m *lw) digitalWrite(_a0 uint8, _a1 uint8) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, uint8) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// pinMode provides a mock function with given fields: _a0, _a1
func (_m *lw) pinMode(_a0 uint8, _a1 uint8) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, uint8) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// pwmInit provides a mock function with given fields:
func (_m *lw) pwmInit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// pwmStop provides a mock function with given fields:
func (_m *lw) pwmStop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// pwmUpdateCompare provides a mock function with given fields: _a0, _a1
func (_m *lw) pwmUpdateCompare(_a0 uint8, _a1 uint8) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, uint8) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// pwmUpdatePrescaler provides a mock function with given fields: _a0
func (_m *lw) pwmUpdatePrescaler(_a0 uint) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// servoInit provides a mock function with given fields:
func (_m *lw) servoInit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// servoUpdateLocation provides a mock function with given fields: _a0, _a1
func (_m *lw) servoUpdateLocation(_a0 uint8, _a1 uint8) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, uint8) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// error provides a mock function with given fields:
func (_m *lw) error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
