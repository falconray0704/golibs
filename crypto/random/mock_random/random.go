// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/crypto/random (interfaces: Random)

// Package mock_random is a generated GoMock package.
package mock_random

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRandom is a mock of Random interface
type MockRandom struct {
	ctrl     *gomock.Controller
	recorder *MockRandomMockRecorder
}

// MockRandomMockRecorder is the mock recorder for MockRandom
type MockRandomMockRecorder struct {
	mock *MockRandom
}

// NewMockRandom creates a new mock instance
func NewMockRandom(ctrl *gomock.Controller) *MockRandom {
	mock := &MockRandom{ctrl: ctrl}
	mock.recorder = &MockRandomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRandom) EXPECT() *MockRandomMockRecorder {
	return m.recorder
}

// GenerateRandomAlphaNum mocks base method
func (m *MockRandom) GenerateRandomAlphaNum(arg0 uint64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomAlphaNum", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateRandomAlphaNum indicates an expected call of GenerateRandomAlphaNum
func (mr *MockRandomMockRecorder) GenerateRandomAlphaNum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomAlphaNum", reflect.TypeOf((*MockRandom)(nil).GenerateRandomAlphaNum), arg0)
}

// GenerateRandomBytes mocks base method
func (m *MockRandom) GenerateRandomBytes(arg0 int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomBytes", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRandomBytes indicates an expected call of GenerateRandomBytes
func (mr *MockRandomMockRecorder) GenerateRandomBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomBytes", reflect.TypeOf((*MockRandom)(nil).GenerateRandomBytes), arg0)
}

// GenerateRandomInt mocks base method
func (m *MockRandom) GenerateRandomInt(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomInt", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GenerateRandomInt indicates an expected call of GenerateRandomInt
func (mr *MockRandomMockRecorder) GenerateRandomInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomInt", reflect.TypeOf((*MockRandom)(nil).GenerateRandomInt), arg0)
}

// GenerateRandomInt63 mocks base method
func (m *MockRandom) GenerateRandomInt63() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomInt63")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GenerateRandomInt63 indicates an expected call of GenerateRandomInt63
func (mr *MockRandomMockRecorder) GenerateRandomInt63() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomInt63", reflect.TypeOf((*MockRandom)(nil).GenerateRandomInt63))
}

// GenerateRandomNum mocks base method
func (m *MockRandom) GenerateRandomNum(arg0 uint64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandomNum", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateRandomNum indicates an expected call of GenerateRandomNum
func (mr *MockRandomMockRecorder) GenerateRandomNum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandomNum", reflect.TypeOf((*MockRandom)(nil).GenerateRandomNum), arg0)
}
