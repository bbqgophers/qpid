package mocks

import "github.com/stretchr/testify/mock"

import cv "github.com/lazywei/go-opencv/opencv"

type window struct {
	mock.Mock
}

// ShowImage provides a mock function with given fields: _a0
func (_m *window) ShowImage(_a0 *cv.IplImage) {
	_m.Called(_a0)
}
