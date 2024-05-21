// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/ucp/dataprovider (interfaces: DataStorageProvider)
//
// Generated by this command:
//
//	mockgen -destination=./mock_datastorage_provider.go -package=dataprovider -self_package github.com/radius-project/radius/pkg/ucp/dataprovider github.com/radius-project/radius/pkg/ucp/dataprovider DataStorageProvider
//

// Package dataprovider is a generated GoMock package.
package dataprovider

import (
	context "context"
	reflect "reflect"

	store "github.com/radius-project/radius/pkg/ucp/store"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStorageProvider is a mock of DataStorageProvider interface.
type MockDataStorageProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDataStorageProviderMockRecorder
}

// MockDataStorageProviderMockRecorder is the mock recorder for MockDataStorageProvider.
type MockDataStorageProviderMockRecorder struct {
	mock *MockDataStorageProvider
}

// NewMockDataStorageProvider creates a new mock instance.
func NewMockDataStorageProvider(ctrl *gomock.Controller) *MockDataStorageProvider {
	mock := &MockDataStorageProvider{ctrl: ctrl}
	mock.recorder = &MockDataStorageProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStorageProvider) EXPECT() *MockDataStorageProviderMockRecorder {
	return m.recorder
}

// GetStorageClient mocks base method.
func (m *MockDataStorageProvider) GetStorageClient(arg0 context.Context, arg1 string) (store.StorageClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageClient", arg0, arg1)
	ret0, _ := ret[0].(store.StorageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageClient indicates an expected call of GetStorageClient.
func (mr *MockDataStorageProviderMockRecorder) GetStorageClient(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageClient", reflect.TypeOf((*MockDataStorageProvider)(nil).GetStorageClient), arg0, arg1)
}
