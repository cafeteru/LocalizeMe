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
	oneResult := mongo.InsertOneResult{
		InsertedID: stage.ID,
	}
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&oneResult, nil)
	service := GroupServiceImpl{repository}
	result, err := service.Create(groupDto)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestGroupServiceImpl_Create_Error_NameRegister(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	repository.EXPECT().FindByName(gomock.Any()).Return(&group, nil)
	service := GroupServiceImpl{repository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func TestGroupServiceImpl_Create_ErrorRepository(t *testing.T) {
	initGroupValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	expectedError := errors.New(constants.InsertGroup)
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := GroupServiceImpl{repository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func TestGroupServiceImpl_Create_ErrorRequest_InvalidName(t *testing.T) {
	initGroupValues()
	groupDto.Name = ""
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockGroupRepository(mockCtrl)
	service := GroupServiceImpl{repository}
	_, err := service.Create(groupDto)
	assert.NotNil(t, err)
}

func initGroupValues() {
	id := "1"
	objectID, _ := primitive.ObjectIDFromHex(id)
	group = domain.Group{
		ID:          objectID,
		Name:        "group",
		Owner:       user,
		Permissions: []domain.Permission{},
		Active:      true,
	}
	groupDto = dto.GroupDto{
		Name:        group.Name,
		Owner:       group.Owner,
		Permissions: []dto.PermissionDto{},
	}
}
