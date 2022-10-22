package impl

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	encryptMock "uniovi-localizeme/internal/core/utils/encrypt/mocks"
	"uniovi-localizeme/internal/repository/mock"
)

var user domain.User
var userDto dto.UserDto

func TestUserServiceImpl_CreateGroupService(t *testing.T) {
	service := CreateUserService()
	assert.NotNil(t, service)
	assert.NotNil(t, service.repository)
}

func TestUserServiceImpl_Create_Successful(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	oneResult := mongo.InsertOneResult{
		InsertedID: user.ID,
	}
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(&oneResult, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.Create(userDto)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, result.ID)
}

func TestUserServiceImpl_Create_Error_EmailRegister(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Create(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorRepository(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	expectedError := errors.New(constants.InsertUser)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	repository.EXPECT().Create(gomock.Any()).Return(nil, expectedError)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Create(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidEmail(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	userDto := dto.UserDto{
		Email:    "",
		Password: "password",
	}
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Create(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorUserRequest_InvalidPassword(t *testing.T) {
	initUserValues()
	userDto.Password = ""
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Create(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Create_ErrorEncrypt(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	expectedError := errors.New(constants.EncryptPasswordUser)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", expectedError)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Create(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Delete_Successful(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	mongoResult := mongo.DeleteResult{
		DeletedCount: 1,
	}
	repository.EXPECT().Delete(gomock.Any()).Return(&mongoResult, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.Delete(user.ID)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestUserServiceImpl_Delete_NotFoundById(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	err := errors.New(constants.FindUserById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.Delete(user.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_Delete_ErrorRepository(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	err := errors.New(constants.DeleteUser)
	repository.EXPECT().Delete(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.Delete(user.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_Disable_Successful(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.Disable(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, user.ID)
}

func TestUserServiceImpl_Disable_NotFoundById(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	err := errors.New(constants.FindUserById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.Disable(user.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_Disable_ErrorRepository(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	err := errors.New(constants.UpdateUser)
	repository.EXPECT().Update(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.Disable(user.ID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_FindById_NotFound(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	err := errors.New(constants.FindUserById)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	objectID, _ := primitive.ObjectIDFromHex("1.1")
	_, expectedError := service.FindById(objectID)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_FindById_UserNotActive(t *testing.T) {
	initUserValues()
	user.Active = false
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.FindById(user.ID)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindById_Success(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	service := UserServiceImpl{repository, encrypt}
	userById, err := service.FindById(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userById.Email)
}

func TestUserServiceImpl_FindByEmail_Empty(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	err := errors.New(constants.InvalidUserRequest)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.FindByEmail("")
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_FindByEmail_NotFound(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	err := errors.New(constants.FindUserByEmail)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, err)
	service := UserServiceImpl{repository, encrypt}
	_, expectedError := service.FindByEmail(user.Email)
	assert.NotNil(t, expectedError)
	assert.Equal(t, expectedError, err)
}

func TestUserServiceImpl_FindByEmail_UserNotActive(t *testing.T) {
	initUserValues()
	user.Active = false
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.FindByEmail(user.Email)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_FindByEmail_Success(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	service := UserServiceImpl{repository, encrypt}
	userById, err := service.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userById.Email)
}

func TestUserServiceImpl_Login_NotFound(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New(""))
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Login(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Login_ErrorPassword(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	encrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Login(userDto)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Login_Successful(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
	encrypt.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Login(userDto)
	assert.Nil(t, err)
}

func TestUserServiceImpl_FindAll_Success(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	user2 := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "email2@email.com",
		Password: "email2@email.com",
		Admin:    true,
		Active:   false,
	}
	users := []domain.User{user, user2}
	repository.EXPECT().FindAll().Return(&users, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(*result), len(users))
}

func TestUserServiceImpl_FindAll_Error(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindAll().Return(nil, errors.New(""))
	service := UserServiceImpl{repository, encrypt}
	_, err := service.FindAll()
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Successful(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.Update(user)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, user.ID)
}

func TestUserServiceImpl_Update_Error_NotIdRegister(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(nil, nil)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Update(user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_EncryptPassword(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", errors.New(""))
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Update(user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_Error_Repository(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)

	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	encrypt.EXPECT().EncryptPassword(gomock.Any()).Return("", nil)
	repository.EXPECT().Update(gomock.Any()).Return(nil, errors.New(""))
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Update(user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_EmailAlreadyRegister(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(&domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "",
		Password: "",
		Admin:    false,
		Active:   false,
	}, nil)
	service := UserServiceImpl{repository, encrypt}
	_, err := service.Update(user)
	assert.NotNil(t, err)
}

func TestUserServiceImpl_Update_NoChangePassword(t *testing.T) {
	initUserValues()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repository := mock.NewMockUserRepository(mockCtrl)
	encrypt := encryptMock.NewMockEncrypt(mockCtrl)
	repository.EXPECT().FindById(gomock.Any()).Return(&user, nil)
	repository.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	mongoResult := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	repository.EXPECT().Update(gomock.Any()).Return(&mongoResult, nil)
	service := UserServiceImpl{repository, encrypt}
	result, err := service.Update(domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: "",
		Admin:    user.Admin,
		Active:   user.Active,
	})
	assert.Nil(t, err)
	assert.Equal(t, result.Password, user.Password)
}

func initUserValues() {
	id := "1"
	userId, _ := primitive.ObjectIDFromHex(id)
	user = domain.User{
		ID:       userId,
		Email:    "user@email.com",
		Password: "password",
		Admin:    true,
		Active:   true,
	}
	userDto = dto.UserDto{
		Email:    user.Email,
		Password: user.Password,
	}
}
