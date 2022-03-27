// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/StageRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockStageRepository is a mock of StageRepository interface.
type MockStageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStageRepositoryMockRecorder
}

// MockStageRepositoryMockRecorder is the mock recorder for MockStageRepository.
type MockStageRepositoryMockRecorder struct {
	mock *MockStageRepository
}

// NewMockStageRepository creates a new mock instance.
func NewMockStageRepository(ctrl *gomock.Controller) *MockStageRepository {
	mock := &MockStageRepository{ctrl: ctrl}
	mock.recorder = &MockStageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStageRepository) EXPECT() *MockStageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStageRepository) Create(stage domain.Stage) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", stage)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStageRepositoryMockRecorder) Create(stage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStageRepository)(nil).Create), stage)
}

// FindAll mocks base method.
func (m *MockStageRepository) FindAll() (*[]domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(*[]domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockStageRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStageRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockStageRepository) FindById(id primitive.ObjectID) (*domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockStageRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockStageRepository)(nil).FindById), id)
}

// FindByName mocks base method.
func (m *MockStageRepository) FindByName(name string) (*domain.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(*domain.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockStageRepositoryMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockStageRepository)(nil).FindByName), name)
}

// Update mocks base method.
func (m *MockStageRepository) Update(stage domain.Stage) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", stage)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStageRepositoryMockRecorder) Update(stage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStageRepository)(nil).Update), stage)
}
