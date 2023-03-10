// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	bookmark "github.com/vaguecoder/firefox-backups/pkg/bookmark"

	mock "github.com/stretchr/testify/mock"
)

// Encoder is an autogenerated mock type for the Encoder type
type Encoder struct {
	mock.Mock
}

// Encode provides a mock function with given fields: _a0
func (_m *Encoder) Encode(_a0 []bookmark.Bookmark) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]bookmark.Bookmark) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Filename provides a mock function with given fields:
func (_m *Encoder) Filename() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Encoder) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewEncoder interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncoder creates a new instance of Encoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncoder(t mockConstructorTestingTNewEncoder) *Encoder {
	mock := &Encoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
