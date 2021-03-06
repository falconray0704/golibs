// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/admin (interfaces: Gotify)

// Package mock_admin is a generated GoMock package.
package mock_admin

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGotify is a mock of Gotify interface
type MockGotify struct {
	ctrl     *gomock.Controller
	recorder *MockGotifyMockRecorder
}

// MockGotifyMockRecorder is the mock recorder for MockGotify
type MockGotifyMockRecorder struct {
	mock *MockGotify
}

// NewMockGotify creates a new mock instance
func NewMockGotify(ctrl *gomock.Controller) *MockGotify {
	mock := &MockGotify{ctrl: ctrl}
	mock.recorder = &MockGotifyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGotify) EXPECT() *MockGotifyMockRecorder {
	return m.recorder
}

// Notify mocks base method
func (m *MockGotify) Notify(arg0 string, arg1 int, arg2 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Notify", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify
func (mr *MockGotifyMockRecorder) Notify(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockGotify)(nil).Notify), varargs...)
}

// Ping mocks base method
func (m *MockGotify) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockGotifyMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockGotify)(nil).Ping))
}
