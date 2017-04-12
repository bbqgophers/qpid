package mocks

import "github.com/stretchr/testify/mock"

type AnalogReader struct {
	mock.Mock
}

// AnalogRead provides a mock function with given fields: _a0
func (_m *AnalogReader) AnalogRead(_a0 string) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
