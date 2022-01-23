package impl

import (
	"errors"
	"github.com/golang/mock/gomock"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	encryptMock "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt/mocks"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserServiceImpl_Create_NotRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	id := primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	result := mongo.InsertOneResult{
		InsertedID: id,
	}
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mockUserRepository.EXPECT().FindByEmail(userRequest.Email).Return(nil, nil)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(&result, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err != nil {
		t.Error("Expected", errors.New(constants.ErrorEmailAlreadyRegister), "Got", err)
	}
}

func TestUserServiceImpl_Create_Register(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	user := domain.User{
		ID:       primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Email:    userRequest.Email,
		Password: userRequest.Password,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(constants.ErrorEmailAlreadyRegister), "Got", err)
	}
}

func TestUserServiceImpl_Create_ErrorRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.ErrorInsertUser)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", expectedError, "Got", err)
	}
}

func TestUserServiceImpl_Create_ErrorUserRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)

	userRequest := dto.UserRequest{
		Email:    "",
		Password: "password",
	}
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(constants.ErrorInvalidUserRequest), "Got", err)
	}

	userRequest = dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	_, err = userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(constants.ErrorInvalidUserRequest), "Got", err)
	}
}

func TestUserServiceImpl_Create_ErrorEncrypt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)

	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.ErrorEncryptPasswordUser)
	mockUserRepository.EXPECT().FindByEmail(userRequest.Email).Return(nil, nil)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", expectedError)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", expectedError, "Got", err)
	}
}
