package mocks

import "github.com/briandowns/qpid/vendor/github.com/prometheus/client_golang/api/prometheus"
import "github.com/stretchr/testify/mock"

import "time"
import "github.com/prometheus/common/model"
import "golang.org/x/net/context"

// *DO NOT EDIT* Auto-generated via mockery

type QueryAPI struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, query, ts
func (_m *QueryAPI) Query(ctx context.Context, query string, ts time.Time) (model.Value, error) {
	ret := _m.Called(ctx, query, ts)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) model.Value); ok {
		r0 = rf(ctx, query, ts)
	} else {
		r0 = ret.Get(0).(model.Value)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, query, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRange provides a mock function with given fields: ctx, query, r
func (_m *QueryAPI) QueryRange(ctx context.Context, query string, r prometheus.Range) (model.Value, error) {
	ret := _m.Called(ctx, query, r)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, prometheus.Range) model.Value); ok {
		r0 = rf(ctx, query, r)
	} else {
		r0 = ret.Get(0).(model.Value)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, prometheus.Range) error); ok {
		r1 = rf(ctx, query, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
