package impl

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	encryptMock "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt/mocks"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserServiceImpl_Create_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	id := primitive.NewObjectID()
	result := mongo.InsertOneResult{
		InsertedID: id,
	}
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mockUserRepository.EXPECT().FindByEmail(userRequest.Email).Return(nil, nil)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(&result, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	user, err := userService.Create(userRequest)
	assert.Nil(t, err)
	assert.Equal(t, id, user.ID)
}

func TestUserServiceImpl_Create_Error_EmailRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	user := domain.User{
		ID:       primitive.NewObjectID(),
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
	assert.NotNil(t, err)
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
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidEmail(t *testing.T) {
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
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	userRequest := dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	_, err := userService.Create(userRequest)
	assert.NotNil(t, err)
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
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_EmptyEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindByEmail("")
	assert.NotNil(t, err)
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
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_UserNotActive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: false,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindByEmail(email)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	user2, err := userService.FindByEmail(email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, user2.Email)
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
	assert.NotNil(t, err)
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
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	mockEncrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Login(userRequest)
	assert.NotNil(t, err)
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
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	mockEncrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Login(userRequest)
	assert.Nil(t, err)
}

func TestUserServiceImpl_FindAll_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "email@email.com",
		Password: "email@email.com",
		IsAdmin:  false,
		IsActive: true,
	}
	user2 := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: false,
	}
	users := []domain.User{user, user2}
	mockUserRepository.EXPECT().FindAll().Return(&users, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	result, err := userService.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(users))
}

func TestUserServiceImpl_FindAll_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	mockUserRepository.EXPECT().FindAll().Return(nil, errors.New(""))
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.FindAll()
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	mockUserRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&mongoResult, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	result, err := userService.Update(id, user)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, user.ID)
}

func TestUserServiceImpl_Update_Error_NotIdRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Update(id, user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_EncryptPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", errors.New(""))
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Update(id, user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_Repository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mock.NewMockUserRepository(mockCtrl)
	mockEncrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	mockUserRepository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	mockUserRepository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	mockEncrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mockUserRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))
	userService := CreateUserService(mockUserRepository, mockEncrypt)
	_, err := userService.Update(id, user)
	assert.NotNil(t, err)
}
