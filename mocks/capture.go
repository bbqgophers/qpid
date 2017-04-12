package mocks

import "github.com/stretchr/testify/mock"

import cv "github.com/lazywei/go-opencv/opencv"

type capture struct {
	mock.Mock
}

// RetrieveFrame provides a mock function with given fields: _a0
func (_m *capture) RetrieveFrame(_a0 int) *cv.IplImage {
	ret := _m.Called(_a0)

	var r0 *cv.IplImage
	if rf, ok := ret.Get(0).(func(int) *cv.IplImage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cv.IplImage)
		}
	}

	return r0
}

// GrabFrame provides a mock function with given fields:
func (_m *capture) GrabFrame() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
