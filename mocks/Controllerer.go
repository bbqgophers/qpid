package mocks

import "github.com/stretchr/testify/mock"

import "github.com/bbqgophers/qpid"

type Controllerer struct {
	mock.Mock
}

// FoodMonitors provides a mock function with given fields:
func (_m *Controllerer) FoodMonitors() []qpid.Monitor {
	ret := _m.Called()

	var r0 []qpid.Monitor
	if rf, ok := ret.Get(0).(func() []qpid.Monitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]qpid.Monitor)
		}
	}

	return r0
}

// GrillMonitor provides a mock function with given fields:
func (_m *Controllerer) GrillMonitor() qpid.Monitor {
	ret := _m.Called()

	var r0 qpid.Monitor
	if rf, ok := ret.Get(0).(func() qpid.Monitor); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(qpid.Monitor)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *Controllerer) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Controllerer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *Controllerer) Status() (qpid.GrillStatus, error) {
	ret := _m.Called()

	var r0 qpid.GrillStatus
	if rf, ok := ret.Get(0).(func() qpid.GrillStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(qpid.GrillStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notifications provides a mock function with given fields:
func (_m *Controllerer) Notifications() chan qpid.Notification {
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

// Metrics provides a mock function with given fields:
func (_m *Controllerer) Metrics() chan qpid.GrillStatus {
	ret := _m.Called()

	var r0 chan qpid.GrillStatus
	if rf, ok := ret.Get(0).(func() chan qpid.GrillStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan qpid.GrillStatus)
		}
	}

	return r0
}

// Source provides a mock function with given fields:
func (_m *Controllerer) Source() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
