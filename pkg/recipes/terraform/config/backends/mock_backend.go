// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/recipes/terraform/config/backends (interfaces: Backend)

// Package backends is a generated GoMock package.
package backends

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	recipes "github.com/project-radius/radius/pkg/recipes"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// BuildBackend mocks base method.
func (m *MockBackend) BuildBackend(arg0 *recipes.ResourceMetadata) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildBackend", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildBackend indicates an expected call of BuildBackend.
func (mr *MockBackendMockRecorder) BuildBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildBackend", reflect.TypeOf((*MockBackend)(nil).BuildBackend), arg0)
}
