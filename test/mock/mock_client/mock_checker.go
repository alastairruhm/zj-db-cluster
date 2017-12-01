// Code generated by MockGen. DO NOT EDIT.
// Source: client/base.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecker is a mock of Checker interface
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// CheckConnection mocks base method
func (m *MockChecker) CheckConnection() string {
	ret := m.ctrl.Call(m, "CheckConnection")
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckConnection indicates an expected call of CheckConnection
func (mr *MockCheckerMockRecorder) CheckConnection() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConnection", reflect.TypeOf((*MockChecker)(nil).CheckConnection))
}

// CheckReplicaStatus mocks base method
func (m *MockChecker) CheckReplicaStatus() {
	m.ctrl.Call(m, "CheckReplicaStatus")
}

// CheckReplicaStatus indicates an expected call of CheckReplicaStatus
func (mr *MockCheckerMockRecorder) CheckReplicaStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReplicaStatus", reflect.TypeOf((*MockChecker)(nil).CheckReplicaStatus))
}

// CheckReplicaConsistency mocks base method
func (m *MockChecker) CheckReplicaConsistency() {
	m.ctrl.Call(m, "CheckReplicaConsistency")
}

// CheckReplicaConsistency indicates an expected call of CheckReplicaConsistency
func (mr *MockCheckerMockRecorder) CheckReplicaConsistency() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReplicaConsistency", reflect.TypeOf((*MockChecker)(nil).CheckReplicaConsistency))
}