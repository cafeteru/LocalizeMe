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

func TestStageServiceImpl_Create_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	request := dto.StageRequest{
		Name: "name",
	}
	id := primitive.NewObjectID()
	result := mongo.InsertOneResult{
		InsertedID: id,
	}
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&result, nil)
	service := CreateStageService(repository)
	user, err := service.Create(request)
	assert.Nil(t, err)
	assert.Equal(t, id, user.ID)
}

func TestStageServiceImpl_Create_Error_NameRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	request := dto.StageRequest{
		Name: "name",
	}
	stage := domain.Stage{
		ID:     primitive.NewObjectID(),
		Name:   "name",
		Active: true,
	}
	repository.EXPECT().FindByName(gomock.Any()).Return(&stage, nil)
	service := CreateStageService(repository)
	_, err := service.Create(request)
	if err == nil {
		t.Error("Expected", errors.New(constants.StageAlreadyRegister), "Got", err)
	}
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Create_ErrorRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	userRequest := dto.StageRequest{
		Name: "name",
	}
	expectedError := errors.New(constants.InsertStage)
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := CreateStageService(repository)
	_, err := service.Create(userRequest)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Create_ErrorStageRequest_InvalidName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	request := dto.StageRequest{
		Name: "",
	}
	stageService := CreateStageService(repository)
	_, err := stageService.Create(request)
	assert.NotNil(t, err)
}
