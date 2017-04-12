package mocks

import "github.com/briandowns/qpid/vendor/github.com/hybridgroup/gobot"
import "github.com/stretchr/testify/mock"

type Eventer struct {
	mock.Mock
}

// Events provides a mock function with given fields:
func (_m *Eventer) Events() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Event provides a mock function with given fields: name
func (_m *Eventer) Event(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AddEvent provides a mock function with given fields: name
func (_m *Eventer) AddEvent(name string) {
	_m.Called(name)
}

// DeleteEvent provides a mock function with given fields: name
func (_m *Eventer) DeleteEvent(name string) {
	_m.Called(name)
}

// Publish provides a mock function with given fields: name, data
func (_m *Eventer) Publish(name string, data interface{}) {
	_m.Called(name, data)
}

// Subscribe provides a mock function with given fields:
func (_m *Eventer) Subscribe() gobot.eventChannel {
	ret := _m.Called()

	var r0 gobot.eventChannel
	if rf, ok := ret.Get(0).(func() gobot.eventChannel); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(gobot.eventChannel)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: events
func (_m *Eventer) Unsubscribe(events gobot.eventChannel) {
	_m.Called(events)
}

// On provides a mock function with given fields: name, f
func (_m *Eventer) On(name string, f func(interface{})) error {
	ret := _m.Called(name, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(interface{})) error); ok {
		r0 = rf(name, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Once provides a mock function with given fields: name, f
func (_m *Eventer) Once(name string, f func(interface{})) error {
	ret := _m.Called(name, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(interface{})) error); ok {
		r0 = rf(name, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
