package impl

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

var group domain.Group
var groupDto dto.GroupDto

func TestGroupServiceImpl_CreateGroupService(t *testing.T) {
	service := CreateGroupService()
	assert.NotNil(t, service)
	assert.NotNil(t, service.repository)
}

func TestGroupServiceImpl_Create_Successful(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	oneResult := mongo.InsertOneResult{
		InsertedID: stage.ID,
	}
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&oneResult, nil)
	service := GroupServiceImpl{repository, userRepository}
	result, err := service.Create(groupDto)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestGroupServiceImpl_Create_Error_NameRegister(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	repository.EXPECT().FindByName(gomock.Any()).Return(&group, nil)
	service := GroupServiceImpl{repository, userRepository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func TestGroupServiceImpl_Create_ErrorRepository(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	expectedError := errors.New(constants.InsertGroup)
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := GroupServiceImpl{repository, userRepository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func TestGroupServiceImpl_Create_ErrorRequest_InvalidName(t *testing.T) {
	initGroupValues()
	groupDto.Name = ""
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	service := GroupServiceImpl{repository, userRepository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func TestGroupServiceImpl_FindAll_Success(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	language2 := domain.Group{
		ID:          primitive.NewObjectID(),
		Name:        "group2",
		Owner:       nil,
		Permissions: nil,
		Active:      true,
	}
	languages := []domain.Group{group, language2}
	repository.EXPECT().FindAll().Return(&languages, nil)
	service := GroupServiceImpl{repository, userRepository}
	result, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(languages))
}

func TestGroupServiceImpl_FindAll_Error(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	userRepository := mock.NewMockUserRepository(mockCtrl)
	repository.EXPECT().FindAll().Return(nil, errors.New(""))
	service := GroupServiceImpl{repository, userRepository}
	_, err := service.FindAll()
	assert.NotNil(t, err)
}

func initGroupValues() {
	id := "1"
	objectID, _ := primitive.ObjectIDFromHex(id)
	group = domain.Group{
		ID:          objectID,
		Name:        "group",
		Owner:       &user,
		Permissions: []domain.Permission{},
		Active:      true,
	}
	groupDto = dto.GroupDto{
		Name:        group.Name,
		Owner:       *group.Owner,
		Permissions: group.Permissions,
		Public:      group.Public,
	}
}
