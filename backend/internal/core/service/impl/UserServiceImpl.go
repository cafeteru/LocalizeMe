package impl

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user, err := u.checkRequest(request)
	if err != nil {
		return user, err
	}
	password, err := u.encrypt.EncryptPassword(request.Password)
	user.Password = password
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.User{}, tools.ErrorLogDetails(err, constants.EncryptPasswordUser, tools.GetCurrentFuncName())
	}
	resultId, err := u.repository.Create(user)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.User{}, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return domain.User{
		ID:       resultId.InsertedID.(primitive.ObjectID),
		Email:    user.Email,
		Password: "",
		Admin:    user.Admin,
		Active:   true,
	}, nil
}

func (u UserServiceImpl) checkRequest(request dto.UserRequest) (domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if request.Email == "" || request.Password == "" {
		return domain.User{}, tools.ErrorLog(constants.InvalidUserRequest, tools.GetCurrentFuncName())
	}
	user, _ := u.repository.FindByEmail(request.Email)
	if user != nil {
		return domain.User{}, tools.ErrorLog(constants.EmailAlreadyRegister, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return domain.User{
		Email:    request.Email,
		Password: request.Password,
		Admin:    false,
		Active:   true,
	}, nil
}

func (u UserServiceImpl) Delete(id primitive.ObjectID) (bool, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if user == nil || err != nil {
		return false, tools.ErrorLog(constants.FindUserById, tools.GetCurrentFuncName())
	}
	_, err = u.repository.Delete(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (u UserServiceImpl) Disable(id primitive.ObjectID) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if user == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindUserById, tools.GetCurrentFuncName())
	}
	user.Active = !user.Active
	_, err = u.repository.Update(*user)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return user, nil
}

func (u UserServiceImpl) FindAll() (*[]domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	users, err := u.repository.FindAll()
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return users, nil
}

func (u UserServiceImpl) FindByEmail(email string) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if email == "" {
		return nil, tools.ErrorLog(constants.InvalidUserRequest, tools.GetCurrentFuncName())
	}
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	if !user.Active {
		errActive := errors.New(constants.UserNoActive)
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, errActive
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return user, nil
}

func (u UserServiceImpl) FindById(id primitive.ObjectID) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user, err := u.repository.FindById(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	if !user.Active {
		errActive := errors.New(constants.UserNoActive)
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, errActive
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return user, nil
}

func (u UserServiceImpl) Login(request dto.UserRequest) (*dto.TokenDto, error) {
	user, err := u.repository.FindByEmail(request.Email)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	if !u.encrypt.CheckPassword(user.Password, request.Password) {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, errors.New(constants.DataLogin)
	}
	claims := jwt.MapClaims{
		"ID":     user.ID,
		"Email":  user.Email,
		"Admin":  user.Admin,
		"Active": user.Active,
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

func (u UserServiceImpl) Update(request domain.User) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	original, err := u.repository.FindById(request.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindUserById, tools.GetCurrentFuncName())
	}
	userEmail, _ := u.repository.FindByEmail(request.Email)
	if userEmail != nil && userEmail.ID != request.ID {
		return nil, tools.ErrorLog(constants.EmailAlreadyRegister, tools.GetCurrentFuncName())
	}
	if request.Password != "" {
		password, err := u.encrypt.EncryptPassword(request.Password)
		if err != nil {
			log.Printf("%s: error", tools.GetCurrentFuncName())
			return nil, tools.ErrorLogDetails(err, constants.EncryptPasswordUser, tools.GetCurrentFuncName())
		}
		request.Password = password
	} else {
		request.Password = original.Password
	}
	_, err = u.repository.Update(request)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &request, nil
}
