package impl

import (
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return domain.User{}, tools.ErrorLogDetails(err, constants.ErrorEncryptPasswordUser, tools.GetCurrentFuncName())
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
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
		IsActive: true,
	}, err
}

func (u UserServiceImpl) checkRequest(request dto.UserRequest) (domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	if request.Email == "" || request.Password == "" {
		return domain.User{}, tools.ErrorLog(constants.ErrorInvalidUserRequest, tools.GetCurrentFuncName())
	}
	user, err := u.repository.FindByEmail(request.Email)
	if user != nil {
		return domain.User{}, tools.ErrorLogDetails(err, constants.ErrorEmailAlreadyRegister, tools.GetCurrentFuncName())
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return domain.User{}, nil
}
