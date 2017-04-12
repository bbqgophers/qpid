package mocks

import "github.com/briandowns/qpid/vendor/github.com/hybridgroup/gobot/sysfs"
import "github.com/stretchr/testify/mock"

import "os"

type Filesystem struct {
	mock.Mock
}

// OpenFile provides a mock function with given fields: name, flag, perm
func (_m *Filesystem) OpenFile(name string, flag int, perm os.FileMode) (sysfs.File, error) {
	ret := _m.Called(name, flag, perm)

	var r0 sysfs.File
	if rf, ok := ret.Get(0).(func(string, int, os.FileMode) sysfs.File); ok {
		r0 = rf(name, flag, perm)
	} else {
		r0 = ret.Get(0).(sysfs.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, os.FileMode) error); ok {
		r1 = rf(name, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
