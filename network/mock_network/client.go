// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/golibs/network (interfaces: Client)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	gomock "github.com/golang/mock/gomock"
	network "github.com/qdm12/golibs/network"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// DoHTTPRequest mocks base method
func (m *MockClient) DoHTTPRequest(arg0 *http.Request) (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoHTTPRequest", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DoHTTPRequest indicates an expected call of DoHTTPRequest
func (mr *MockClientMockRecorder) DoHTTPRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoHTTPRequest", reflect.TypeOf((*MockClient)(nil).DoHTTPRequest), arg0)
}

// GetContent mocks base method
func (m *MockClient) GetContent(arg0 string, arg1 ...network.GetContentSetter) ([]byte, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetContent indicates an expected call of GetContent
func (mr *MockClientMockRecorder) GetContent(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockClient)(nil).GetContent), varargs...)
}
