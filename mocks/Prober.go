package mocks

import "github.com/stretchr/testify/mock"

import "github.com/bbqgophers/messages"
import "github.com/bbqgophers/qpid"

// *DO NOT EDIT* Auto-generated via mockery

type Prober struct {
	mock.Mock
}

// Target provides a mock function with given fields: temp
func (_m *Prober) Target(temp messages.Temp) (messages.Temp, error) {
	ret := _m.Called(temp)

	var r0 messages.Temp
	if rf, ok := ret.Get(0).(func(messages.Temp) messages.Temp); ok {
		r0 = rf(temp)
	} else {
		r0 = ret.Get(0).(messages.Temp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(messages.Temp) error); ok {
		r1 = rf(temp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Setpoint provides a mock function with given fields:
func (_m *Prober) Setpoint() (messages.Temp, error) {
	ret := _m.Called()

	var r0 messages.Temp
	if rf, ok := ret.Get(0).(func() messages.Temp); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(messages.Temp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HighThreshold provides a mock function with given fields: temp
func (_m *Prober) HighThreshold(temp messages.Temp) {
	_m.Called(temp)
}

// LowThreshold provides a mock function with given fields: temp
func (_m *Prober) LowThreshold(temp messages.Temp) {
	_m.Called(temp)
}

// Alerts provides a mock function with given fields:
func (_m *Prober) Alerts() chan qpid.Notification {
	ret := _m.Called()

	var r0 chan qpid.Notification
	if rf, ok := ret.Get(0).(func() chan qpid.Notification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan qpid.Notification)
		}
	}

	return r0
}

// Temperature provides a mock function with given fields:
func (_m *Prober) Temperature() (messages.Temp, error) {
	ret := _m.Called()

	var r0 messages.Temp
	if rf, ok := ret.Get(0).(func() messages.Temp); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(messages.Temp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Location provides a mock function with given fields:
func (_m *Prober) Location() messages.Location {
	ret := _m.Called()

	var r0 messages.Location
	if rf, ok := ret.Get(0).(func() messages.Location); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(messages.Location)
	}

	return r0
}

// Description provides a mock function with given fields:
func (_m *Prober) Description() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Source provides a mock function with given fields:
func (_m *Prober) Source() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
