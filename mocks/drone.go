package mocks

import "github.com/stretchr/testify/mock"

// *DO NOT EDIT* Auto-generated via mockery

type drone struct {
	mock.Mock
}

// TakeOff provides a mock function with given fields:
func (_m *drone) TakeOff() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Land provides a mock function with given fields:
func (_m *drone) Land() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Up provides a mock function with given fields: n
func (_m *drone) Up(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Down provides a mock function with given fields: n
func (_m *drone) Down(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Left provides a mock function with given fields: n
func (_m *drone) Left(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Right provides a mock function with given fields: n
func (_m *drone) Right(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Forward provides a mock function with given fields: n
func (_m *drone) Forward(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backward provides a mock function with given fields: n
func (_m *drone) Backward(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Clockwise provides a mock function with given fields: n
func (_m *drone) Clockwise(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CounterClockwise provides a mock function with given fields: n
func (_m *drone) CounterClockwise(n int) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *drone) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields:
func (_m *drone) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Video provides a mock function with given fields:
func (_m *drone) Video() chan []byte {
	ret := _m.Called()

	var r0 chan []byte
	if rf, ok := ret.Get(0).(func() chan []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan []byte)
		}
	}

	return r0
}

// StartRecording provides a mock function with given fields:
func (_m *drone) StartRecording() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopRecording provides a mock function with given fields:
func (_m *drone) StopRecording() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HullProtection provides a mock function with given fields: protect
func (_m *drone) HullProtection(protect bool) error {
	ret := _m.Called(protect)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(protect)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Outdoor provides a mock function with given fields: outdoor
func (_m *drone) Outdoor(outdoor bool) error {
	ret := _m.Called(outdoor)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(outdoor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
