// Code generated by mockery v2.33.1. DO NOT EDIT.

package proxytracker

import (
	limiter "github.com/hashicorp/consul/agent/grpc-external/limiter"
	mock "github.com/stretchr/testify/mock"
)

// MockSessionLimiter is an autogenerated mock type for the SessionLimiter type
type MockSessionLimiter struct {
	mock.Mock
}

// BeginSession provides a mock function with given fields:
func (_m *MockSessionLimiter) BeginSession() (limiter.Session, error) {
	ret := _m.Called()

	var r0 limiter.Session
	var r1 error
	if rf, ok := ret.Get(0).(func() (limiter.Session, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() limiter.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(limiter.Session)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockSessionLimiter creates a new instance of MockSessionLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSessionLimiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSessionLimiter {
	mock := &MockSessionLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
