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

func (b BaseStringServiceImpl) Create(request domain.BaseString) (domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
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
		Translations:   nil,
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
