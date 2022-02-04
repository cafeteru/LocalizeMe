package impl

import (
	"errors"
	"github.com/golang/mock/gomock"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	encryptMock "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt/mocks"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserServiceImpl_Create_NotRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
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
		t.Error("Expected", errors.New(constants.EmailAlreadyRegister), "Got", err)
	}
}

func TestUserServiceImpl_Create_Register(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
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
		t.Error("Expected", errors.New(constants.EmailAlreadyRegister), "Got", err)
	}
}

func TestUserServiceImpl_Create_ErrorRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.InsertUser)
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
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)

	userRequest := dto.UserRequest{
		Email:    "",
		Password: "password",
	}
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(constants.InvalidUserRequest), "Got", err)
	}

	userRequest = dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	_, err = userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(constants.InvalidUserRequest), "Got", err)
	}
}

func TestUserServiceImpl_Create_ErrorEncrypt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)

	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.EncryptPasswordUser)
	mockUserRepository.EXPECT().FindByEmail(userRequest.Email).Return(nil, nil)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", expectedError)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Create(userRequest)
	if err == nil {
		t.Error("Expected", expectedError, "Got", err)
	}
}

func TestUserServiceImpl_FindByEmail_EmptyEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindByEmail("")
	if err == nil {
		t.Error("Expected", errors.New(constants.InvalidUserRequest), "Got", err)
	}
}

func TestUserServiceImpl_FindByEmail_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	mockUserRepository.EXPECT().FindByEmail(email).Return(nil, errors.New(""))
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindByEmail(email)
	if err == nil {
		t.Error("Expected", errors.New(""), "Got", err)
	}
}

func TestUserServiceImpl_FindByEmail_UserNotActive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: false,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindByEmail(email)
	if err == nil {
		t.Error("Expected", errors.New(constants.UserNoActive), "Got", err)
	}
}

func TestUserServiceImpl_FindByEmail_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	user2, err := userService.FindByEmail(email)
	if err != nil || user.Email != user2.Email {
		t.Error("Expected", errors.New(constants.UserNoActive), "Got", err)
	}
}

func TestUserServiceImpl_Login_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	mockUserRepository.EXPECT().FindByEmail(email).Return(nil, errors.New(""))
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	userRequest := dto.UserRequest{
		Email:    email,
		Password: email,
	}
	_, err := userService.Login(userRequest)
	if err == nil {
		t.Error("Expected", errors.New(""), "Got", err)
	}
}

func TestUserServiceImpl_Login_ErrorPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	mockEncrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Login(userRequest)
	if err == nil {
		t.Error("Expected", "", "Got", err)
	}
}

func TestUserServiceImpl_Login_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.ObjectID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	mockEncrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Login(userRequest)
	if err != nil {
		t.Error("Expected", "", "Got", err)
	}
}
