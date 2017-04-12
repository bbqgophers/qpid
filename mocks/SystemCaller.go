package mocks

import "github.com/stretchr/testify/mock"

import "syscall"

type SystemCaller struct {
	mock.Mock
}

// Syscall provides a mock function with given fields: trap, a1, a2, a3
func (_m *SystemCaller) Syscall(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, uintptr, syscall.Errno) {
	ret := _m.Called(trap, a1, a2, a3)

	var r0 uintptr
	if rf, ok := ret.Get(0).(func(uintptr, uintptr, uintptr, uintptr) uintptr); ok {
		r0 = rf(trap, a1, a2, a3)
	} else {
		r0 = ret.Get(0).(uintptr)
	}

	var r1 uintptr
	if rf, ok := ret.Get(1).(func(uintptr, uintptr, uintptr, uintptr) uintptr); ok {
		r1 = rf(trap, a1, a2, a3)
	} else {
		r1 = ret.Get(1).(uintptr)
	}

	var r2 syscall.Errno
	if rf, ok := ret.Get(2).(func(uintptr, uintptr, uintptr, uintptr) syscall.Errno); ok {
		r2 = rf(trap, a1, a2, a3)
	} else {
		r2 = ret.Get(2).(syscall.Errno)
	}

	return r0, r1, r2
}
