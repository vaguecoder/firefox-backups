// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// ChmodRevertFunc is an autogenerated mock type for the ChmodRevertFunc type
type ChmodRevertFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: filename
func (_m *ChmodRevertFunc) Execute(filename string) (fs.FileMode, error) {
	ret := _m.Called(filename)

	var r0 fs.FileMode
	if rf, ok := ret.Get(0).(func(string) fs.FileMode); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(fs.FileMode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewChmodRevertFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewChmodRevertFunc creates a new instance of ChmodRevertFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChmodRevertFunc(t mockConstructorTestingTNewChmodRevertFunc) *ChmodRevertFunc {
	mock := &ChmodRevertFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}