// Code generated by MockGen. DO NOT EDIT.
// Source: mock_test.go

// Package mock is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	mytest "carrymec/go_mock/interface"
	gomock "github.com/golang/mock/gomock"
)

// MockFooInterface is a mock of FooInterface interface.
type MockFooInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFooInterfaceMockRecorder
}

// MockFooInterfaceMockRecorder is the mock recorder for MockFooInterface.
type MockFooInterfaceMockRecorder struct {
	mock *MockFooInterface
}

// NewMockFooInterface creates a new mock instance.
func NewMockFooInterface(ctrl *gomock.Controller) *MockFooInterface {
	mock := &MockFooInterface{ctrl: ctrl}
	mock.recorder = &MockFooInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFooInterface) EXPECT() *MockFooInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFooInterface) Get(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFooInterfaceMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFooInterface)(nil).Get), name)
}

// Set mocks base method.
func (m *MockFooInterface) Set(user mytest.FooUser) (mytest.FooUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", user)
	ret0, _ := ret[0].(mytest.FooUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockFooInterfaceMockRecorder) Set(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockFooInterface)(nil).Set), user)
}
