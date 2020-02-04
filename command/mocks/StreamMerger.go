// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// StreamMerger is an autogenerated mock type for the StreamMerger type
type StreamMerger struct {
	mock.Mock
}

// CollectLines provides a mock function with given fields: onNewLine
func (_m *StreamMerger) CollectLines(onNewLine func(string)) error {
	ret := _m.Called(onNewLine)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(string)) error); ok {
		r0 = rf(onNewLine)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Merge provides a mock function with given fields: name, stream
func (_m *StreamMerger) Merge(name string, stream io.ReadCloser) {
	_m.Called(name, stream)
}
