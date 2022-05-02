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

var id string
var user domain.User
var claims jwt.MapClaims

func TestCheckTokenIsActive_CheckUserIsActive_IsActive(t *testing.T) {
	mockUserService := initMocks(t)
	user.Active = true
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindById(user.ID).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.NotNil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_IsNotActive(t *testing.T) {
	mockUserService := initMocks(t)
	user.Active = false
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindById(user.ID).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_InvalidToken(t *testing.T) {
	mockUserService := initMocks(t)
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindById(user.ID).Return(&user, nil)
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func TestCheckTokenIsActive_CheckUserIsActive_ErrorUser(t *testing.T) {
	mockUserService := initMocks(t)
	request, w := createRequestWithToken(claims)
	mockUserService.EXPECT().FindById(user.ID).Return(nil, errors.New(constants.UserNoRegister))
	result := CheckUserIsActive(w, request, mockUserService)
	assert.Nil(t, result)
}

func initMocks(t *testing.T) *mock.MockUserService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mock.NewMockUserService(mockCtrl)
	id = "62212b92ab63141a684739f3"
	objectID, _ := primitive.ObjectIDFromHex(id)
	user = domain.User{
		ID:       objectID,
		Email:    "username",
		Password: "",
		Admin:    false,
		Active:   false,
	}
	claims = jwt.MapClaims{
		"email":  user.Email,
		"admin":  user.Admin,
		"active": user.Active,
		"id":     id,
	}
	return mockUserService
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
