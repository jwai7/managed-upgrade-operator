// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/drain (interfaces: DrainStrategy)

// Package drain is a generated GoMock package.
package drain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockDrainStrategy is a mock of DrainStrategy interface.
type MockDrainStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDrainStrategyMockRecorder
}

// MockDrainStrategyMockRecorder is the mock recorder for MockDrainStrategy.
type MockDrainStrategyMockRecorder struct {
	mock *MockDrainStrategy
}

// NewMockDrainStrategy creates a new mock instance.
func NewMockDrainStrategy(ctrl *gomock.Controller) *MockDrainStrategy {
	mock := &MockDrainStrategy{ctrl: ctrl}
	mock.recorder = &MockDrainStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDrainStrategy) EXPECT() *MockDrainStrategyMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDrainStrategy) Execute(arg0 *v1.Node) (*DrainStrategyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(*DrainStrategyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockDrainStrategyMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDrainStrategy)(nil).Execute), arg0)
}

// IsValid mocks base method.
func (m *MockDrainStrategy) IsValid(arg0 *v1.Node) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockDrainStrategyMockRecorder) IsValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockDrainStrategy)(nil).IsValid), arg0)
}
