// Code generated by mockery v1.0.0. DO NOT EDIT.

package helpers

import mock "github.com/stretchr/testify/mock"

// mockTestingT is an autogenerated mock type for the testingT type
type mockTestingT struct {
	mock.Mock
}

// Errorf provides a mock function with given fields: format, args
func (_m *mockTestingT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *mockTestingT) FailNow() {
	_m.Called()
}