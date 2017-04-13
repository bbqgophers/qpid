package mocks

import "github.com/stretchr/testify/mock"

import "net/http"
import "net/url"

import "golang.org/x/net/context"

// *DO NOT EDIT* Auto-generated via mockery

type Client struct {
	mock.Mock
}

// url provides a mock function with given fields: ep, args
func (_m *Client) url(ep string, args map[string]string) *url.URL {
	ret := _m.Called(ep, args)

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func(string, map[string]string) *url.URL); ok {
		r0 = rf(ep, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// do provides a mock function with given fields: _a0, _a1
func (_m *Client) do(_a0 context.Context, _a1 *http.Request) (*http.Response, []byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) []byte); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *http.Request) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
