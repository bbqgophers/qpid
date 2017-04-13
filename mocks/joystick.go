package mocks

import "github.com/stretchr/testify/mock"

import "github.com/veandco/go-sdl2/sdl"

// *DO NOT EDIT* Auto-generated via mockery

type joystick struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *joystick) Close() {
	_m.Called()
}

// InstanceID provides a mock function with given fields:
func (_m *joystick) InstanceID() sdl.JoystickID {
	ret := _m.Called()

	var r0 sdl.JoystickID
	if rf, ok := ret.Get(0).(func() sdl.JoystickID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sdl.JoystickID)
	}

	return r0
}
