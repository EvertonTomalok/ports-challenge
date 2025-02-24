// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package ports is a generated GoMock package.
package ports

import (
	reflect "reflect"

	domain "github.com/EvertonTomalok/ports-challenge/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockService) Upsert(ports domain.PortData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ports)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceMockRecorder) Upsert(ports interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockService)(nil).Upsert), ports)
}
