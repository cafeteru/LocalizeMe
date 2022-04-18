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

var language domain.Language
var languageDto dto.LanguageDto

func TestLanguageServiceImpl_CreateGroupService(t *testing.T) {
	service := CreateLanguageService()
	assert.NotNil(t, service)
	assert.NotNil(t, service.repository)
}

func TestLanguageServiceImpl_Create_Successful(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	oneResult := mongo.InsertOneResult{
		InsertedID: stage.ID,
	}
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&oneResult, nil)
	service := LanguageServiceImpl{repository}
	result, err := service.Create(languageDto)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestLanguageServiceImpl_Create_Error_NameRegister(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(&language, nil)
	service := LanguageServiceImpl{repository}
	_, err := service.Create(languageDto)
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Create_ErrorRepository(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	expectedError := errors.New(constants.InsertLanguage)
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := LanguageServiceImpl{repository}
	_, err := service.Create(languageDto)
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Create_ErrorRequest_InvalidName(t *testing.T) {
	initLanguageValues()
	languageDto.IsoCode = ""
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	service := LanguageServiceImpl{repository}
	_, err := service.Create(languageDto)
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Delete_Successful(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	mongoResult := mongo.DeleteResult{
		DeletedCount: 1,
	}
	repository.EXPECT().Delete(gomock.Any()).Return(&mongoResult, nil)
	service := LanguageServiceImpl{repository}
	result, err := service.Delete(language.ID)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestLanguageServiceImpl_Delete_NotFoundById(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	err := errors.New(constants.FindLanguageById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := LanguageServiceImpl{repository}
	_, expectedError := service.Delete(language.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestLanguageServiceImpl_Delete_ErrorRepository(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	err := errors.New(constants.DeleteLanguage)
	repository.EXPECT().Delete(gomock.Any()).Return(nil, err)
	service := LanguageServiceImpl{repository}
	_, expectedError := service.Delete(language.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestLanguageServiceImpl_Disable_Successful(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := LanguageServiceImpl{repository}
	result, err := service.Disable(stage.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, stage.ID)
}

func TestLanguageServiceImpl_Disable_NotFoundById(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	err := errors.New(constants.FindLanguageById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := LanguageServiceImpl{repository}
	_, expectedError := service.Disable(stage.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestLanguageServiceImpl_Disable_ErrorRepository(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	err := errors.New(constants.UpdateLanguage)
	repository.EXPECT().Update(gomock.Any()).Return(nil, err)
	service := LanguageServiceImpl{repository}
	_, expectedError := service.Disable(language.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestLanguageServiceImpl_FindAll_Success(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	language2 := domain.Language{
		ID:          primitive.NewObjectID(),
		IsoCode:     "isoCode",
		Description: "description",
		Active:      true,
	}
	languages := []domain.Language{language, language2}
	repository.EXPECT().FindAll().Return(&languages, nil)
	service := LanguageServiceImpl{repository}
	result, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(languages))
}

func TestLanguageServiceImpl_FindAll_Error(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindAll().Return(nil, errors.New(""))
	service := LanguageServiceImpl{repository}
	_, err := service.FindAll()
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Update_Successful(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(nil, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := LanguageServiceImpl{repository}
	result, err := service.Update(language)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, language.ID)
}

func TestLanguageServiceImpl_Update_Error_NotIdRegister(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
	service := LanguageServiceImpl{repository}
	_, err := service.Update(language)
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Update_Error_Repository(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Update(gomock.Any()).Return(nil, errors.New(""))
	service := LanguageServiceImpl{repository}
	_, err := service.Update(language)
	assert.NotNil(t, err)
}

func TestLanguageServiceImpl_Update_NameAlreadyRegister(t *testing.T) {
	initLanguageValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockLanguageRepository(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&language, nil)
	repository.EXPECT().FindByIsoCode(gomock.Any()).Return(&domain.Language{
		ID:          primitive.NewObjectID(),
		IsoCode:     language.IsoCode,
		Description: "",
		Active:      false,
	}, nil)
	service := LanguageServiceImpl{repository}
	_, err := service.Update(language)
	assert.NotNil(t, err)
}

func initLanguageValues() {
	id := "1"
	objectID, _ := primitive.ObjectIDFromHex(id)
	language = domain.Language{
		ID:          objectID,
		IsoCode:     "IsoCode",
		Description: "Description",
		Active:      true,
	}
	languageDto = dto.LanguageDto{
		IsoCode:     language.IsoCode,
		Description: language.Description,
	}
}
