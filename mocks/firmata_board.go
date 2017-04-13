package mocks

import "github.com/stretchr/testify/mock"

import "io"

import "github.com/hybridgroup/gobot/platforms/firmata/client"

// *DO NOT EDIT* Auto-generated via mockery

type firmataBoard struct {
	mock.Mock
}

// Connect provides a mock function with given fields: _a0
func (_m *firmataBoard) Connect(_a0 io.ReadWriteCloser) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.ReadWriteCloser) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *firmataBoard) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pins provides a mock function with given fields:
func (_m *firmataBoard) Pins() []client.Pin {
	ret := _m.Called()

	var r0 []client.Pin
	if rf, ok := ret.Get(0).(func() []client.Pin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Pin)
		}
	}

	return r0
}

// AnalogWrite provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) AnalogWrite(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPinMode provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) SetPinMode(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReportAnalog provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) ReportAnalog(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReportDigital provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) ReportDigital(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DigitalWrite provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) DigitalWrite(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// I2cRead provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) I2cRead(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// I2cWrite provides a mock function with given fields: _a0, _a1
func (_m *firmataBoard) I2cWrite(_a0 int, _a1 []byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// I2cConfig provides a mock function with given fields: _a0
func (_m *firmataBoard) I2cConfig(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServoConfig provides a mock function with given fields: _a0, _a1, _a2
func (_m *firmataBoard) ServoConfig(_a0 int, _a1 int, _a2 int) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Event provides a mock function with given fields: _a0
func (_m *firmataBoard) Event(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
