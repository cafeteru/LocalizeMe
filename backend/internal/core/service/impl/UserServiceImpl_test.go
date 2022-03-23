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
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	id := primitive.NewObjectID()
	result := mongo.InsertOneResult{
		InsertedID: id,
	}
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().FindByEmail(request.Email).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&result, nil)
	service := CreateUserService(repository, encrypt)
	user, err := service.Create(request)
	assert.Nil(t, err)
	assert.Equal(t, id, user.ID)
}

func TestUserServiceImpl_Create_Error_EmailRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  false,
		IsActive: true,
	}
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	service := CreateUserService(repository, encrypt)
	_, err := service.Create(request)
	if err == nil {
		t.Error("Expected", errors.New(constants.EmailAlreadyRegister), "Got", err)
	}
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.InsertUser)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := CreateUserService(repository, encrypt)
	_, err := service.Create(request)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userRequest := dto.UserRequest{
		Email:    "",
		Password: "password",
	}
	service := CreateUserService(repository, encrypt)
	_, err := service.Create(userRequest)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	service := CreateUserService(repository, encrypt)
	request := dto.UserRequest{
		Email:    "email",
		Password: "",
	}
	_, err := service.Create(request)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorEncrypt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
		Email:    "email",
		Password: "password",
	}
	expectedError := errors.New(constants.EncryptPasswordUser)
	repository.EXPECT().FindByEmail(request.Email).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", expectedError)
	service := CreateUserService(repository, encrypt)
	_, err := service.Create(request)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindById_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, errors.New(""))
	service := CreateUserService(repository, encrypt)
	objectID, _ := primitive.ObjectIDFromHex("1.1")
	_, err := service.FindById(objectID)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_UserNotActive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	objectID, _ := primitive.ObjectIDFromHex("1.1")
	user := domain.User{
		ID:       objectID,
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: false,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	service := CreateUserService(repository, encrypt)
	_, err := service.FindById(objectID)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	objectID, _ := primitive.ObjectIDFromHex("1.1")
	user := domain.User{
		ID:       objectID,
		Email:    email,
		Password: email,
		IsAdmin:  false,
		IsActive: true,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	service := CreateUserService(repository, encrypt)
	userById, err := service.FindById(objectID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userById.Email)
}

func TestUserServiceImpl_Login_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	email := "email@email.com"
	repository.EXPECT().FindByEmail(email).Return(nil, errors.New(""))
	service := CreateUserService(repository, encrypt)
	request := dto.UserRequest{
		Email:    email,
		Password: email,
	}
	_, err := service.Login(request)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Login_ErrorPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
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
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	encrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false)
	service := CreateUserService(repository, encrypt)
	_, err := service.Login(request)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Login_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	request := dto.UserRequest{
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
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	encrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true)
	service := CreateUserService(repository, encrypt)
	_, err := service.Login(request)
	assert.Nil(t, err)
}

func TestUserServiceImpl_FindAll_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
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
	repository.EXPECT().FindAll().Return(&users, nil)
	service := CreateUserService(repository, encrypt)
	result, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(users))
}

func TestUserServiceImpl_FindAll_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindAll().Return(nil, errors.New(""))
	service := CreateUserService(repository, encrypt)
	_, err := service.FindAll()
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Successful(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&mongoResult, nil)
	service := CreateUserService(repository, encrypt)
	result, err := service.Update(userId, user)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, user.ID)
}

func TestUserServiceImpl_Update_Error_NotIdRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
	service := CreateUserService(repository, encrypt)
	_, err := service.Update(userId, user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_EncryptPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", errors.New(""))
	service := CreateUserService(repository, encrypt)
	_, err := service.Update(userId, user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_Repository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	id := "123"
	userId, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{
		ID:       userId,
		Email:    "email2@email.com",
		Password: "email2@email.com",
		IsAdmin:  true,
		IsActive: true,
	}
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))
	service := CreateUserService(repository, encrypt)
	_, err := service.Update(userId, user)
	assert.NotNil(t, err)
}
