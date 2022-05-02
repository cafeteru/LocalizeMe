package impl

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mongodb"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BaseStringServiceImpl struct {
	repository         repository.BaseStringRepository
	userRepository     repository.UserRepository
	languageRepository repository.LanguageRepository
	stageRepository    repository.StageRepository
	groupRepository    repository.GroupRepository
}

func CreateBaseStringService() *BaseStringServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &BaseStringServiceImpl{
		mongodb.CreateBaseStringRepository(),
		mongodb.CreateUserRepository(),
		mongodb.CreateLanguageRepository(),
		mongodb.CreateStageRepository(),
		mongodb.CreateGroupRepository(),
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (b BaseStringServiceImpl) Create(request domain.BaseString, user *domain.User) (domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := b.checkPermission(request, *user)
	if err != nil {
		return domain.BaseString{}, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = b.checkUniqueIdentifier(request.Identifier)
	if err != nil && err.Error() != constants.FindBaseStringByIdentifier {
		return domain.BaseString{}, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseString := domain.BaseString{
		SourceLanguage: request.SourceLanguage,
		Identifier:     request.Identifier,
		Group:          request.Group,
		Author:         request.Author,
		Active:         true,
		Translations:   request.Translations,
	}
	baseString.Author.ClearPassword()
	err = b.createTranslations(request.Translations)
	if err != nil {
		return domain.BaseString{}, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	resultId, err := b.repository.Create(baseString)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.BaseString{}, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseString.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseString, nil
}

func (b BaseStringServiceImpl) Delete(id primitive.ObjectID, user *domain.User) (bool, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseString, err := b.repository.FindById(id)
	if err != nil {
		return false, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	if !user.Admin && baseString.Group.Owner.ID != user.ID {
		return false, tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
	}
	_, err = b.repository.Delete(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (b BaseStringServiceImpl) Disable(id primitive.ObjectID, user *domain.User) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseString, err := b.repository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = b.checkPermission(*baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseString.Active = !baseString.Active
	_, err = b.repository.Update(*baseString)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseString, nil
}

func (b BaseStringServiceImpl) FindAll() (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseStrings, err := b.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) FindByGroup(id primitive.ObjectID, user *domain.User) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	group, err := b.groupRepository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = CheckGroupPermission(*group, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseStrings, err := b.repository.FindByGroup(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}
func (b BaseStringServiceImpl) FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseStrings, err := b.repository.FindByPermissions(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) Update(baseString domain.BaseString, user *domain.User) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := b.checkPermission(baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	original, err := b.repository.FindById(baseString.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	if original.Identifier != baseString.Identifier {
		if baseString.Identifier == "" {
			return nil, tools.ErrorLog(constants.IdentifierBaseStringInvalid, tools.GetCurrentFuncName())
		}
		findByName, err := b.repository.FindByIdentifier(baseString.Identifier)
		if findByName != nil {
			return nil, tools.ErrorLog(constants.IdentifierBaseStringAlreadyRegister, tools.GetCurrentFuncName())
		}
		if err != nil && err.Error() != constants.FindGroupByName {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	_, err = b.repository.Update(baseString)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b BaseStringServiceImpl) checkUniqueIdentifier(identifier string) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if identifier == "" {
		return tools.ErrorLog(constants.IdentifierBaseStringInvalid, tools.GetCurrentFuncName())
	}
	baseString, err := b.repository.FindByIdentifier(identifier)
	if baseString != nil {
		return tools.ErrorLog(constants.IdentifierBaseStringAlreadyRegister, tools.GetCurrentFuncName())
	}
	if err != nil {
		return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (b BaseStringServiceImpl) createTranslations(request []domain.Translation) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	for _, translation := range request {
		email := translation.Author.Email
		_, err := b.userRepository.FindByEmail(email)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		isoCode := translation.Language.IsoCode
		_, err = b.languageRepository.FindByIsoCode(isoCode)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		name := translation.Stage.Name
		_, err = b.stageRepository.FindByName(name)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (b BaseStringServiceImpl) checkPermission(baseString domain.BaseString, user domain.User) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user.Admin || baseString.Group.Public || baseString.Group.Owner.ID == user.ID {
		return nil
	}
	for _, permission := range baseString.Group.Permissions {
		if permission.User.ID == user.ID && permission.CanWrite {
			return nil
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
}
