// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/service/StageService.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	dto "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

// MockStageService is a mock of StageService interface.
type MockStageService struct {
	ctrl     *gomock.Controller
	recorder *MockStageServiceMockRecorder
}

// MockStageServiceMockRecorder is the mock recorder for MockStageService.
type MockStageServiceMockRecorder struct {
	mock *MockStageService
}

// NewMockStageService creates a new mock instance.
func NewMockStageService(ctrl *gomock.Controller) *MockStageService {
	mock := &MockStageService{ctrl: ctrl}
	mock.recorder = &MockStageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStageService) EXPECT() *MockStageServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStageService) Create(request dto.StageRequest) (domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", request)
	ret0, _ := ret[0].(domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStageServiceMockRecorder) Create(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStageService)(nil).Create), request)
}

// FindAll mocks base method.
func (m *MockStageService) FindAll() (*[]domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(*[]domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockStageServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStageService)(nil).FindAll))
}

// Update mocks base method.
func (m *MockStageService) Update(request domain.Stage) (*domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", request)
	ret0, _ := ret[0].(*domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStageServiceMockRecorder) Update(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStageService)(nil).Update), request)
}
