package utils

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/mock"
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
	user.Active = true
	claims := jwt.MapClaims{
		"Email":  user.Email,
		"Admin":  user.Admin,
		"Active": user.Active,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.NotNil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_IsNotActive(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	user.Active = false

	claims := jwt.MapClaims{
		"Email":  user.Email,
		"Admin":  user.Admin,
		"Active": user.Active,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_InvalidToken(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"Admin": user.Admin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_NotRegisterUser(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"Email": user.Email,
		"Admin": user.Admin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(nil, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_ErrorUser(t *testing.T) {
	mockUserService := initMocks(t)
	user := createUser()
	claims := jwt.MapClaims{
		"Email": user.Email,
		"Admin": user.Admin,
	}
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindByEmail(user.Email).Return(nil, errors.New(constants.UserNoRegister))
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
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
		Admin:    false,
		Active:   false,
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
