// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/build-index/tagclient (interfaces: Client)

// Package mocktagclient is a generated GoMock package.
package mocktagclient

import (
	gomock "github.com/golang/mock/gomock"
	core "github.com/uber/kraken/core"
	reflect "reflect"
	time "time"
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

// DuplicatePut mocks base method
func (m *MockClient) DuplicatePut(arg0 string, arg1 core.Digest, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DuplicatePut", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DuplicatePut indicates an expected call of DuplicatePut
func (mr *MockClientMockRecorder) DuplicatePut(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicatePut", reflect.TypeOf((*MockClient)(nil).DuplicatePut), arg0, arg1, arg2)
}

// DuplicateReplicate mocks base method
func (m *MockClient) DuplicateReplicate(arg0 string, arg1 core.Digest, arg2 core.DigestList, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DuplicateReplicate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DuplicateReplicate indicates an expected call of DuplicateReplicate
func (mr *MockClientMockRecorder) DuplicateReplicate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateReplicate", reflect.TypeOf((*MockClient)(nil).DuplicateReplicate), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockClient) Get(arg0 string) (core.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(core.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0)
}

// Has mocks base method
func (m *MockClient) Has(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockClientMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClient)(nil).Has), arg0)
}

// List mocks base method
func (m *MockClient) List(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClient)(nil).List), arg0)
}

// ListRepository mocks base method
func (m *MockClient) ListRepository(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepository", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepository indicates an expected call of ListRepository
func (mr *MockClientMockRecorder) ListRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepository", reflect.TypeOf((*MockClient)(nil).ListRepository), arg0)
}

// Origin mocks base method
func (m *MockClient) Origin() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Origin")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Origin indicates an expected call of Origin
func (mr *MockClientMockRecorder) Origin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MockClient)(nil).Origin))
}

// Put mocks base method
func (m *MockClient) Put(arg0 string, arg1 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockClientMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0, arg1)
}

// PutAndReplicate mocks base method
func (m *MockClient) PutAndReplicate(arg0 string, arg1 core.Digest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutAndReplicate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutAndReplicate indicates an expected call of PutAndReplicate
func (mr *MockClientMockRecorder) PutAndReplicate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutAndReplicate", reflect.TypeOf((*MockClient)(nil).PutAndReplicate), arg0, arg1)
}

// Replicate mocks base method
func (m *MockClient) Replicate(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replicate indicates an expected call of Replicate
func (mr *MockClientMockRecorder) Replicate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicate", reflect.TypeOf((*MockClient)(nil).Replicate), arg0)
}