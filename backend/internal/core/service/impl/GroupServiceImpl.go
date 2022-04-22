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
	findByName, errName, validName := g.checkUniqueName(request.Name)
	if !validName {
		return findByName, errName
	}
	permissions, d, err2 := g.createPermissions(request)
	if err2 != nil {
		return d, err2
	}
	group := domain.Group{
		Name:        request.Name,
		Owner:       request.Owner,
		Active:      true,
		Public:      request.Public,
		Permissions: permissions,
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

func (g GroupServiceImpl) createPermissions(request dto.GroupDto) ([]domain.Permission, domain.Group, error) {
	var permissions []domain.Permission
	for _, permissionDto := range request.Permissions {
		email := permissionDto.Email
		user, err := g.userRepository.FindByEmail(email)
		if err != nil {
			return nil, domain.Group{}, err
		}
		permission := domain.Permission{
			User:          *user,
			CanWriteGroup: permissionDto.CanWriteGroup,
		}
		permissions = append(permissions, permission)
	}
	return permissions, domain.Group{}, nil
}

func (g GroupServiceImpl) FindAll() (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := g.repository.FindAll()
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (g GroupServiceImpl) FindByPermissions(email string) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := g.repository.FindByPermissions(email)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (g GroupServiceImpl) checkUniqueName(name string) (domain.Group, error, bool) {
	if name == "" {
		return domain.Group{}, tools.ErrorLog(constants.NameGroupInvalid, tools.GetCurrentFuncName()), false
	}
	findByName, _ := g.repository.FindByName(name)
	if findByName != nil {
		return domain.Group{}, tools.ErrorLog(constants.GroupAlreadyRegister, tools.GetCurrentFuncName()), false
	}
	return domain.Group{}, nil, true
}
