package utils

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/mocks"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckTokenIsActive_CheckUserIsActive_IsActive(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	user.IsActive = true
	claims := jwt.MapClaims{
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
		"isActive": user.IsActive,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	isActive, _ := CheckUserIsActive(w, request, mockUserService)
	if !isActive {
		t.Error("Expected", true, "Got", false)
	}
}

func TestCheckTokenIsActive_CheckUserIsActive_IsNotActive(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	user.IsActive = false

	claims := jwt.MapClaims{
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
		"isActive": user.IsActive,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	isActive, _ := CheckUserIsActive(w, request, mockUserService)
	if isActive {
		t.Error("Expected", true, "Got", false)
	}
}

func TestCheckTokenIsActive_CheckUserIsActive_InvalidToken(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"isAdmin": user.IsAdmin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	isActive, _ := CheckUserIsActive(w, request, mockUserService)
	if isActive {
		t.Error("Expected", false, "Got", true)
	}
}

func TestCheckTokenIsActive_CheckUserIsActive_NotRegisterUser(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"email":   user.Email,
		"isAdmin": user.IsAdmin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(nil, nil)
	isActive, _ := CheckUserIsActive(w, request, mockUserService)
	if isActive {
		t.Error("Expected", false, "Got", true)
	}
}

func TestCheckTokenIsActive_CheckUserIsActive_ErrorUser(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"email":   user.Email,
		"isAdmin": user.IsAdmin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(nil, errors.New(constants.UserNoRegister))
	isActive, _ := CheckUserIsActive(w, request, mockUserService)
	if isActive {
		t.Error("Expected", false, "Got", true)
	}
}

func initMocks(t *testing.T) *mocks.MockUserService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mocks.NewMockUserService(mockCtrl)
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

func createRequestWithToken(claims jwt.MapClaims) (*http.Request, *httptest.ResponseRecorder) {
	tools.LoadEnv()
	alg := "HS256"
	secret := "secretLocalizeMe"
	jwtauth.SetExpiry(claims, time.Now().Add(30_000))
	tokenAuth := jwtauth.New(alg, []byte(secret), nil)
	token, _, _ := tokenAuth.Encode(claims)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	ctx := request.Context()
	ctx = jwtauth.NewContext(ctx, token, nil)
	request = request.WithContext(ctx)
	w := httptest.NewRecorder()
	return request, w
}
