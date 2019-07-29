// Code generated by MockGen. DO NOT EDIT.
// Source: credentials/sources.go

// Package credentials is a generated GoMock package.
package credentials

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSource is a mock of Source interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Credentials mocks base method
func (m *MockSource) Credentials() ([]Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credentials")
	ret0, _ := ret[0].([]Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Credentials indicates an expected call of Credentials
func (mr *MockSourceMockRecorder) Credentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credentials", reflect.TypeOf((*MockSource)(nil).Credentials))
}

// Type mocks base method
func (m *MockSource) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockSourceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockSource)(nil).Type))
}

// ValidateConfiguration mocks base method
func (m *MockSource) ValidateConfiguration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateConfiguration")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateConfiguration indicates an expected call of ValidateConfiguration
func (mr *MockSourceMockRecorder) ValidateConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateConfiguration", reflect.TypeOf((*MockSource)(nil).ValidateConfiguration))
}

// MockSourceCollection is a mock of SourceCollection interface
type MockSourceCollection struct {
	ctrl     *gomock.Controller
	recorder *MockSourceCollectionMockRecorder
}

// MockSourceCollectionMockRecorder is the mock recorder for MockSourceCollection
type MockSourceCollectionMockRecorder struct {
	mock *MockSourceCollection
}

// NewMockSourceCollection creates a new mock instance
func NewMockSourceCollection(ctrl *gomock.Controller) *MockSourceCollection {
	mock := &MockSourceCollection{ctrl: ctrl}
	mock.recorder = &MockSourceCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceCollection) EXPECT() *MockSourceCollectionMockRecorder {
	return m.recorder
}

// AllSources mocks base method
func (m *MockSourceCollection) AllSources() []Source {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSources")
	ret0, _ := ret[0].([]Source)
	return ret0
}

// AllSources indicates an expected call of AllSources
func (mr *MockSourceCollectionMockRecorder) AllSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSources", reflect.TypeOf((*MockSourceCollection)(nil).AllSources))
}

// Credentials mocks base method
func (m *MockSourceCollection) Credentials() ([]Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credentials")
	ret0, _ := ret[0].([]Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Credentials indicates an expected call of Credentials
func (mr *MockSourceCollectionMockRecorder) Credentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credentials", reflect.TypeOf((*MockSourceCollection)(nil).Credentials))
}

// ValidateConfiguration mocks base method
func (m *MockSourceCollection) ValidateConfiguration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateConfiguration")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateConfiguration indicates an expected call of ValidateConfiguration
func (mr *MockSourceCollectionMockRecorder) ValidateConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateConfiguration", reflect.TypeOf((*MockSourceCollection)(nil).ValidateConfiguration))
}
