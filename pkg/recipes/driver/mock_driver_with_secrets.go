// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/recipes/driver (interfaces: DriverWithSecrets)
//
// Generated by this command:
//
//	mockgen -destination=./mock_driver_with_secrets.go -package=driver -self_package github.com/radius-project/radius/pkg/recipes/driver github.com/radius-project/radius/pkg/recipes/driver DriverWithSecrets
//

// Package driver is a generated GoMock package.
package driver

import (
	context "context"
	reflect "reflect"

	recipes "github.com/radius-project/radius/pkg/recipes"
	gomock "go.uber.org/mock/gomock"
)

// MockDriverWithSecrets is a mock of DriverWithSecrets interface.
type MockDriverWithSecrets struct {
	ctrl     *gomock.Controller
	recorder *MockDriverWithSecretsMockRecorder
}

// MockDriverWithSecretsMockRecorder is the mock recorder for MockDriverWithSecrets.
type MockDriverWithSecretsMockRecorder struct {
	mock *MockDriverWithSecrets
}

// NewMockDriverWithSecrets creates a new mock instance.
func NewMockDriverWithSecrets(ctrl *gomock.Controller) *MockDriverWithSecrets {
	mock := &MockDriverWithSecrets{ctrl: ctrl}
	mock.recorder = &MockDriverWithSecretsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriverWithSecrets) EXPECT() *MockDriverWithSecretsMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDriverWithSecrets) Delete(arg0 context.Context, arg1 DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDriverWithSecretsMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDriverWithSecrets)(nil).Delete), arg0, arg1)
}

// Execute mocks base method.
func (m *MockDriverWithSecrets) Execute(arg0 context.Context, arg1 ExecuteOptions) (*recipes.RecipeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(*recipes.RecipeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockDriverWithSecretsMockRecorder) Execute(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDriverWithSecrets)(nil).Execute), arg0, arg1)
}

// FindSecretIDs mocks base method.
func (m *MockDriverWithSecrets) FindSecretIDs(arg0 context.Context, arg1 recipes.Configuration, arg2 recipes.EnvironmentDefinition) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSecretIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSecretIDs indicates an expected call of FindSecretIDs.
func (mr *MockDriverWithSecretsMockRecorder) FindSecretIDs(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSecretIDs", reflect.TypeOf((*MockDriverWithSecrets)(nil).FindSecretIDs), arg0, arg1, arg2)
}

// GetRecipeMetadata mocks base method.
func (m *MockDriverWithSecrets) GetRecipeMetadata(arg0 context.Context, arg1 BaseOptions) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipeMetadata", arg0, arg1)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipeMetadata indicates an expected call of GetRecipeMetadata.
func (mr *MockDriverWithSecretsMockRecorder) GetRecipeMetadata(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipeMetadata", reflect.TypeOf((*MockDriverWithSecrets)(nil).GetRecipeMetadata), arg0, arg1)
}
