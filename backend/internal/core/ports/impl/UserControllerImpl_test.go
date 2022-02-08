package impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
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
	userRequest := dto.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest("POST", "http://localhost:8080/login", body)
	ctx := r.Context()
	r = r.WithContext(ctx)
	w := httptest.NewRecorder()
	tokenDto := dto.TokenDto{Authorization: ""}
	mockUserService.EXPECT().Login(gomock.Any()).Return(&tokenDto, nil)
	controllerImpl := UserControllerImpl{mockUserService}

	controllerImpl.Login(w, r)
	if w.Code != http.StatusOK {
		t.Error("Expected", http.StatusOK, "Got", w.Code)
	}
}

func TestUserControllerImpl_Login_EmptyBody(t *testing.T) {
	mockUserService := initMocks(t)
	r := httptest.NewRequest("POST", "http://localhost:8080/login", nil)
	ctx := r.Context()
	r = r.WithContext(ctx)
	w := httptest.NewRecorder()
	controllerImpl := UserControllerImpl{mockUserService}

	controllerImpl.Login(w, r)
	if w.Code != http.StatusUnprocessableEntity {
		t.Error("Expected", http.StatusUnprocessableEntity, "Got", w.Code)
	}
}

func TestUserControllerImpl_Login_NoRegister(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	userRequest := dto.UserRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest("POST", "http://localhost:8080/login", body)
	ctx := r.Context()
	r = r.WithContext(ctx)
	w := httptest.NewRecorder()
	mockUserService.EXPECT().Login(gomock.Any()).Return(nil, errors.New(constants.UserNoRegister))
	controllerImpl := UserControllerImpl{mockUserService}
	controllerImpl.Login(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("Expected", http.StatusBadRequest, "Got", w.Code)
	}
}

func initMocks(t *testing.T) *mock.MockUserService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mock.NewMockUserService(mockCtrl)
	return mockUserService
}

func createUser() domain.User {
	user := domain.User{
		ID:       primitive.ObjectID{},
		Email:    "username",
		Password: "",
		IsAdmin:  false,
		IsActive: false,
	}
	return user
}
