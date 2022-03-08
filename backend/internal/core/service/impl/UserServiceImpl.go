package impl

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	slog "github.com/go-eden/slf4go"
	"github.com/golang-jwt/jwt"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	encrypt    encrypt.Encrypt
}

func CreateUserService(r repository.UserRepository, e encrypt.Encrypt) *UserServiceImpl {
	return &UserServiceImpl{r, e}
}

func (u UserServiceImpl) Create(request dto.UserRequest) (domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user, err := u.checkRequest(request)
	if err != nil {
		return user, err
	}
	password, err := u.encrypt.EncryptPassword(request.Password)
	user.Password = password
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return domain.User{}, tools.ErrorLogDetails(err, constants.EncryptPasswordUser, tools.GetCurrentFuncName())
	}
	resultId, err := u.repository.Create(user)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return domain.User{}, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return domain.User{
		ID:       resultId.InsertedID.(primitive.ObjectID),
		Email:    user.Email,
		Password: "",
		IsAdmin:  user.IsAdmin,
		IsActive: true,
	}, nil
}

func (u UserServiceImpl) checkRequest(request dto.UserRequest) (domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	if request.Email == "" || request.Password == "" {
		return domain.User{}, tools.ErrorLog(constants.InvalidUserRequest, tools.GetCurrentFuncName())
	}
	user, _ := u.repository.FindByEmail(request.Email)
	if user != nil {
		return domain.User{}, tools.ErrorLog(constants.EmailAlreadyRegister, tools.GetCurrentFuncName())
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return domain.User{
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  false,
		IsActive: true,
	}, nil
}

func (u UserServiceImpl) Delete(id string) (bool, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if user == nil || err != nil {
		return false, tools.ErrorLog(constants.FindUserById, tools.GetCurrentFuncName())
	}
	_, err = u.repository.Delete(id)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (u UserServiceImpl) Disable(id string) (*domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if user == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindUserByEmail, tools.GetCurrentFuncName())
	}
	user.IsActive = !user.IsActive
	_, err = u.repository.Update(id, *user)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		IsActive: user.IsActive,
	}, nil
}

func (u UserServiceImpl) FindAll() (*[]domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	users, err := u.repository.FindAll()
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return users, nil
}

func (u UserServiceImpl) FindByEmail(email string) (*domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	if email == "" {
		return nil, tools.ErrorLog(constants.InvalidUserRequest, tools.GetCurrentFuncName())
	}
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	if !user.IsActive {
		errActive := errors.New(constants.UserNoActive)
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, errActive
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return user, nil
}

func (u UserServiceImpl) Login(request dto.UserRequest) (*dto.TokenDto, error) {
	user, err := u.FindByEmail(request.Email)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	if !u.encrypt.CheckPassword(user.Password, request.Password) {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, errors.New(constants.DataLogin)
	}
	claims := jwt.MapClaims{
		"Email":    user.Email,
		"IsAdmin":  user.IsAdmin,
		"IsActive": user.IsActive,
	}
	tools.LoadEnv()
	hours, _ := time.ParseDuration(os.Getenv("HOURS"))
	alg := os.Getenv("ALG")
	secret := os.Getenv("SECRET")
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour+hours))
	tokenAuth := jwtauth.New(alg, []byte(secret), nil)
	_, tokenString, _ := tokenAuth.Encode(claims)
	return &dto.TokenDto{Authorization: tokenString}, nil
}

func (u UserServiceImpl) Update(id string, request domain.User) (*domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if user == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindUserById, tools.GetCurrentFuncName())
	}
	byUserEmail, _ := u.repository.FindByEmail(request.Email)
	objectId, _ := primitive.ObjectIDFromHex(id)
	if byUserEmail != nil && byUserEmail.ID != objectId {
		return nil, tools.ErrorLog(constants.EmailAlreadyRegister, tools.GetCurrentFuncName())
	}
	if request.Password != "" {
		password, err := u.encrypt.EncryptPassword(request.Password)
		request.Password = password
		if err != nil {
			slog.Errorf("%s: error", tools.GetCurrentFuncName())
			return nil, tools.ErrorLogDetails(err, constants.EncryptPasswordUser, tools.GetCurrentFuncName())
		}
	}
	_, err = u.repository.Update(id, request)
	if err != nil {
		slog.Errorf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &domain.User{
		ID:       user.ID,
		Email:    request.Email,
		IsAdmin:  request.IsAdmin,
		IsActive: request.IsActive,
	}, nil
}
