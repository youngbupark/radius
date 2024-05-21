// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/cli/prompt (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -destination=./mock_prompter.go -package=prompt -self_package github.com/radius-project/radius/pkg/cli/prompt github.com/radius-project/radius/pkg/cli/prompt Interface
//

// Package prompt is a generated GoMock package.
package prompt

import (
	reflect "reflect"

	tea "github.com/charmbracelet/bubbletea"
	text "github.com/radius-project/radius/pkg/cli/prompt/text"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// GetListInput mocks base method.
func (m *MockInterface) GetListInput(arg0 []string, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListInput", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListInput indicates an expected call of GetListInput.
func (mr *MockInterfaceMockRecorder) GetListInput(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListInput", reflect.TypeOf((*MockInterface)(nil).GetListInput), arg0, arg1)
}

// GetTextInput mocks base method.
func (m *MockInterface) GetTextInput(arg0 string, arg1 text.TextModelOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextInput", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTextInput indicates an expected call of GetTextInput.
func (mr *MockInterfaceMockRecorder) GetTextInput(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextInput", reflect.TypeOf((*MockInterface)(nil).GetTextInput), arg0, arg1)
}

// RunProgram mocks base method.
func (m *MockInterface) RunProgram(arg0 *tea.Program) (tea.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunProgram", arg0)
	ret0, _ := ret[0].(tea.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunProgram indicates an expected call of RunProgram.
func (mr *MockInterfaceMockRecorder) RunProgram(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunProgram", reflect.TypeOf((*MockInterface)(nil).RunProgram), arg0)
}
