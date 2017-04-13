package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type Commander struct {
	mock.Mock
}

// Command provides a mock function with given fields: _a0
func (_m *Commander) Command(_a0 string) func(map[string]interface{}) interface{} {
	ret := _m.Called(_a0)

	var r0 func(map[string]interface{}) interface{}
	if rf, ok := ret.Get(0).(func(string) func(map[string]interface{}) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(map[string]interface{}) interface{})
		}
	}

	return r0
}

// Commands provides a mock function with given fields:
func (_m *Commander) Commands() map[string]func(map[string]interface{}) interface{} {
	ret := _m.Called()

	var r0 map[string]func(map[string]interface{}) interface{}
	if rf, ok := ret.Get(0).(func() map[string]func(map[string]interface{}) interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]func(map[string]interface{}) interface{})
		}
	}

	return r0
}

// AddCommand provides a mock function with given fields: name, command
func (_m *Commander) AddCommand(name string, command func(map[string]interface{}) interface{}) {
	_m.Called(name, command)
}
