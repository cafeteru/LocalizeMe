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
}

func CreateBaseStringService() *BaseStringServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &BaseStringServiceImpl{
		mongodb.CreateBaseStringRepository(),
		mongodb.CreateUserRepository(),
		mongodb.CreateLanguageRepository(),
		mongodb.CreateStageRepository(),
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (b BaseStringServiceImpl) Create(request domain.BaseString, user *domain.User) (domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	errPermission := b.checkPermission(request, *user)
	if errPermission != nil {
		return domain.BaseString{}, tools.ErrorLogWithError(errPermission, tools.GetCurrentFuncName())
	}
	err := b.checkUniqueIdentifier(request.Identifier)
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

func (b BaseStringServiceImpl) FindAll() (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groups, err := b.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return groups, nil
}

func (b BaseStringServiceImpl) FindByPermissions(email string) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	byPermissions, err := b.repository.FindByPermissions(email)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return byPermissions, nil
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
	byIdentifier, err := b.repository.FindByIdentifier(identifier)
	if byIdentifier != nil {
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
		if permission.User.ID == user.ID && permission.CanWriteGroup {
			return nil
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
}
