package impl

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mongodb"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type GroupServiceImpl struct {
	repository     repository.GroupRepository
	userRepository repository.UserRepository
}

func CreateGroupService() *GroupServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &GroupServiceImpl{mongodb.CreateGroupRepository(), mongodb.CreateUserRepository()}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (g GroupServiceImpl) Create(request dto.GroupDto) (domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if request.Name == "" {
		return domain.Group{}, tools.ErrorLog(constants.NameGroupInvalid, tools.GetCurrentFuncName())
	}
	findByName, err := g.repository.FindByName(request.Name)
	if findByName != nil || err != nil && err.Error() != constants.FindGroupByName {
		return domain.Group{}, tools.ErrorLog(constants.GroupAlreadyRegister, tools.GetCurrentFuncName())
	}
	errPermissions := g.createPermissions(request.Permissions)
	if errPermissions != nil {
		return domain.Group{}, errPermissions
	}
	group := domain.Group{
		Name:        request.Name,
		Owner:       &request.Owner,
		Active:      true,
		Public:      request.Public,
		Permissions: request.Permissions,
	}
	group.Owner.ClearPassword()
	resultId, err := g.repository.Create(group)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.Group{}, err
	}
	group.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return group, nil
}

func (g GroupServiceImpl) Delete(id primitive.ObjectID, user *domain.User) (bool, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	group, err := g.repository.FindById(id)
	if group == nil || err != nil {
		return false, tools.ErrorLog(constants.FindGroupById, tools.GetCurrentFuncName())
	}
	if !user.Admin && group.Owner.ID != user.ID {
		return false, tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
	}
	_, err = g.repository.Delete(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (g GroupServiceImpl) Disable(id primitive.ObjectID, user *domain.User) (*domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	group, err := g.repository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	errPermission := g.checkPermission(*group, *user)
	if errPermission != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	group.Active = !group.Active
	_, err = g.repository.Update(*group)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return group, nil
}

func (g GroupServiceImpl) FindAll() (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := g.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (g GroupServiceImpl) FindByPermissions(email string) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := g.repository.FindByPermissions(email)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (g GroupServiceImpl) FindCanWrite(email string) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := g.repository.FindCanWrite(email)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (g GroupServiceImpl) Update(group domain.Group, user *domain.User) (*domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	errPermission := g.checkPermission(group, *user)
	if errPermission != nil {
		return nil, tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
	}
	original, err := g.repository.FindById(group.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindGroupById, tools.GetCurrentFuncName())
	}
	if original.Name != group.Name {
		if group.Name == "" {
			return nil, tools.ErrorLog(constants.NameGroupInvalid, tools.GetCurrentFuncName())
		}
		findByName, err := g.repository.FindByName(group.Name)
		if findByName != nil {
			return nil, tools.ErrorLog(constants.GroupAlreadyRegister, tools.GetCurrentFuncName())
		}
		if err != nil && err.Error() != constants.FindGroupByName {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	errPermission = g.createPermissions(group.Permissions)
	if errPermission != nil {
		return nil, errPermission
	}
	_, err = g.repository.Update(group)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &group, nil
}

func (g GroupServiceImpl) createPermissions(request []domain.Permission) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	for _, permission := range request {
		email := permission.User.Email
		_, err := g.userRepository.FindByEmail(email)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (g GroupServiceImpl) checkPermission(group domain.Group, user domain.User) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user.Admin || group.Public || group.Owner.ID == user.ID {
		return nil
	}
	for _, permission := range group.Permissions {
		if permission.User.ID == user.ID && permission.CanWrite {
			return nil
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
}
