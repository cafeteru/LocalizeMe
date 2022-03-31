package impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserControllerImpl_Login_Successful(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	userRequest := dto.UserDto{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", body)
	w := httptest.NewRecorder()
	tokenDto := dto.TokenDto{Authorization: ""}
	mockUserService.EXPECT().Login(gomock.Any()).Return(&tokenDto, nil)
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUserControllerImpl_Login_EmptyBody(t *testing.T) {
	mockUserService := initMocks(t)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", nil)
	w := httptest.NewRecorder()
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
}

func TestUserControllerImpl_Login_NoRegister(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	userRequest := dto.UserDto{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", body)
	w := httptest.NewRecorder()
	mockUserService.EXPECT().Login(gomock.Any()).Return(nil, errors.New(constants.UserNoRegister))
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestUserControllerImpl_Create_Successful(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	userRequest := dto.UserDto{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", body)
	w := httptest.NewRecorder()
	request := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "",
		Password: "",
		Admin:    false,
		Active:   true,
	}
	mockUserService.EXPECT().Create(gomock.Any()).Return(request, nil)
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestUserControllerImpl_Create_Error_Body(t *testing.T) {
	mockUserService := initMocks(t)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", nil)
	w := httptest.NewRecorder()
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
}

func TestUserControllerImpl_Create_Error_Service(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	userRequest := dto.UserDto{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", body)
	w := httptest.NewRecorder()
	mockUserService.EXPECT().Create(gomock.Any()).Return(domain.User{}, errors.New(""))
	controllerImpl := CreateUserController(mockUserService)
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func initMocks(t *testing.T) *mock.MockUserService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mock.NewMockUserService(mockCtrl)
	return mockUserService
}

func createUser() domain.User {
	user := domain.User{
		ID:     primitive.NewObjectID(),
		Email:  "username",
		Admin:  false,
		Active: false,
	}
	return user
}
