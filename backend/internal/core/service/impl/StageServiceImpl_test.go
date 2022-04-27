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

var stage domain.Stage
var stageDto dto.StageDto

func TestStageServiceImpl_CreateGroupService(t *testing.T) {
	service := CreateStageService()
	assert.NotNil(t, service)
	assert.NotNil(t, service.repository)
}

func TestStageServiceImpl_Create_Successful(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	oneResult := mongo.InsertOneResult{
		InsertedID: stage.ID,
	}
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&oneResult, nil)
	service := StageServiceImpl{repository}
	result, err := service.Create(stageDto)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestStageServiceImpl_Create_Error_NameRegister(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindByName(gomock.Any()).Return(&stage, nil)
	service := StageServiceImpl{repository}
	_, err := service.Create(stageDto)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Create_ErrorRepository(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	expectedError := errors.New(constants.InsertStage)
	repository.EXPECT().FindByName(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := StageServiceImpl{repository}
	_, err := service.Create(stageDto)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Create_ErrorStageRequest_InvalidName(t *testing.T) {
	initStageValues()
	stageDto.Name = ""
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	service := StageServiceImpl{repository}
	_, err := service.Create(stageDto)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Delete_Successful(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	mongoResult := mongo.DeleteResult{
		DeletedCount: 1,
	}
	repository.EXPECT().Delete(gomock.Any()).Return(&mongoResult, nil)
	service := StageServiceImpl{repository}
	result, err := service.Delete(stage.ID)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestStageServiceImpl_Delete_NotFoundById(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	err := errors.New(constants.FindStageById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := StageServiceImpl{repository}
	_, expectedError := service.Delete(stage.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestStageServiceImpl_Delete_ErrorRepository(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	err := errors.New(constants.DeleteStage)
	repository.EXPECT().Delete(gomock.Any()).Return(nil, err)
	service := StageServiceImpl{repository}
	_, expectedError := service.Delete(stage.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestStageServiceImpl_Disable_Successful(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := StageServiceImpl{repository}
	result, err := service.Disable(stage.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestStageServiceImpl_Disable_NotFoundById(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	err := errors.New(constants.FindStageById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := StageServiceImpl{repository}
	_, expectedError := service.Disable(stage.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestStageServiceImpl_Disable_ErrorRepository(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	err := errors.New(constants.UpdateStage)
	repository.EXPECT().Update(gomock.Any()).Return(nil, err)
	service := StageServiceImpl{repository}
	_, expectedError := service.Disable(stage.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestStageServiceImpl_FindAll_Success(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	stage2 := domain.Stage{
		ID:     primitive.NewObjectID(),
		Name:   "name2",
		Active: true,
	}
	stages := []domain.Stage{stage, stage2}
	repository.EXPECT().FindAll().Return(&stages, nil)
	service := StageServiceImpl{repository}
	result, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(stages))
}

func TestStageServiceImpl_FindAll_Error(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindAll().Return(nil, errors.New(""))
	service := StageServiceImpl{repository}
	_, err := service.FindAll()
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Update_Successful(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := StageServiceImpl{repository}
	result, err := service.Update(stage)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestStageServiceImpl_Update_Error_NotIdRegister(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
	service := StageServiceImpl{repository}
	_, err := service.Update(stage)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Update_Error_Repository(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	repository.EXPECT().Update(gomock.Any()).Return(nil, errors.New(""))
	service := StageServiceImpl{repository}
	_, err := service.Update(stage)
	assert.NotNil(t, err)
}

func TestStageServiceImpl_Update_NameAlreadyRegister(t *testing.T) {
	initStageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockStageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&stage, nil)
	repository.EXPECT().FindByName(gomock.Any()).Return(&domain.Stage{
		ID:     primitive.NewObjectID(),
		Name:   stage.Name,
		Active: false,
	}, nil)
	service := StageServiceImpl{repository}
	newStage := domain.Stage{
		ID:     stage.ID,
		Name:   "newName",
		Active: stage.Active,
	}
	_, err := service.Update(newStage)
	assert.NotNil(t, err)
}

func initStageValues() {
	id := "1"
	stageId, _ := primitive.ObjectIDFromHex(id)
	stage = domain.Stage{
		ID:     stageId,
		Active: true,
		Name:   "Name",
	}
	stageDto = dto.StageDto{
		Name: stage.Name,
	}
}
