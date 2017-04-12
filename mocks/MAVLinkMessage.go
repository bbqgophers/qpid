package mocks

import "github.com/stretchr/testify/mock"

type MAVLinkMessage struct {
	mock.Mock
}

// Id provides a mock function with given fields:
func (_m *MAVLinkMessage) Id() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// Len provides a mock function with given fields:
func (_m *MAVLinkMessage) Len() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// Crc provides a mock function with given fields:
func (_m *MAVLinkMessage) Crc() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// Pack provides a mock function with given fields:
func (_m *MAVLinkMessage) Pack() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Decode provides a mock function with given fields: _a0
func (_m *MAVLinkMessage) Decode(_a0 []byte) {
	_m.Called(_a0)
}
