// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/GroupRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "uniovi-localizeme/internal/core/domain"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockGroupRepository is a mock of GroupRepository interface.
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository.
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance.
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGroupRepository) Create(group domain.Group) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", group)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockGroupRepositoryMockRecorder) Create(group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGroupRepository)(nil).Create), group)
}

// Delete mocks base method.
func (m *MockGroupRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockGroupRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupRepository)(nil).Delete), id)
}

// FindAll mocks base method.
func (m *MockGroupRepository) FindAll() (*[]domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(*[]domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockGroupRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockGroupRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockGroupRepository) FindById(id primitive.ObjectID) (*domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockGroupRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockGroupRepository)(nil).FindById), id)
}

// FindByName mocks base method.
func (m *MockGroupRepository) FindByName(name string) (*domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(*domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockGroupRepositoryMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockGroupRepository)(nil).FindByName), name)
}

// FindByPermissions mocks base method.
func (m *MockGroupRepository) FindByPermissions(id primitive.ObjectID) (*[]domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPermissions", id)
	ret0, _ := ret[0].(*[]domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPermissions indicates an expected call of FindByPermissions.
func (mr *MockGroupRepositoryMockRecorder) FindByPermissions(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPermissions", reflect.TypeOf((*MockGroupRepository)(nil).FindByPermissions), id)
}

// FindCanWrite mocks base method.
func (m *MockGroupRepository) FindCanWrite(id primitive.ObjectID) (*[]domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCanWrite", id)
	ret0, _ := ret[0].(*[]domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCanWrite indicates an expected call of FindCanWrite.
func (mr *MockGroupRepositoryMockRecorder) FindCanWrite(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCanWrite", reflect.TypeOf((*MockGroupRepository)(nil).FindCanWrite), id)
}

// Update mocks base method.
func (m *MockGroupRepository) Update(group domain.Group) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", group)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockGroupRepositoryMockRecorder) Update(group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupRepository)(nil).Update), group)
}
