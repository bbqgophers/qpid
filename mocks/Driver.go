package mocks

import "github.com/briandowns/qpid/vendor/github.com/hybridgroup/gobot"
import "github.com/stretchr/testify/mock"

type Driver struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *Driver) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Driver) Start() []error {
	ret := _m.Called()

	var r0 []error
	if rf, ok := ret.Get(0).(func() []error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// Halt provides a mock function with given fields:
func (_m *Driver) Halt() []error {
	ret := _m.Called()

	var r0 []error
	if rf, ok := ret.Get(0).(func() []error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// Connection provides a mock function with given fields:
func (_m *Driver) Connection() gobot.Connection {
	ret := _m.Called()

	var r0 gobot.Connection
	if rf, ok := ret.Get(0).(func() gobot.Connection); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(gobot.Connection)
	}

	return r0
}
