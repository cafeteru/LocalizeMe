package impl

import (
	"encoding/xml"
	"errors"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/xmlDto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mongodb"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type BaseStringServiceImpl struct {
	baseStringRepository repository.BaseStringRepository
	userRepository       repository.UserRepository
	languageRepository   repository.LanguageRepository
	stageRepository      repository.StageRepository
	groupRepository      repository.GroupRepository
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
	resultId, err := b.baseStringRepository.Create(baseString)
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
	baseString, err := b.baseStringRepository.FindById(id)
	if err != nil {
		return false, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	if !user.Admin && baseString.Group.Owner.ID != user.ID {
		return false, tools.ErrorLog(constants.GroupNotHavePermissions, tools.GetCurrentFuncName())
	}
	_, err = b.baseStringRepository.Delete(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (b BaseStringServiceImpl) Disable(id primitive.ObjectID, user *domain.User) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseString, err := b.baseStringRepository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = b.checkPermission(*baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseString.Active = !baseString.Active
	_, err = b.baseStringRepository.Update(*baseString)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseString, nil
}

func (b BaseStringServiceImpl) FindAll() (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseStrings, err := b.baseStringRepository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) FindById(id primitive.ObjectID) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseString, err := b.baseStringRepository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseString, nil
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
	baseStrings, err := b.baseStringRepository.FindByGroup(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) FindByIdentifier(identifier string, user *domain.User) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseString, err := b.baseStringRepository.FindByIdentifier(identifier)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = b.checkPermission(*baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseString, nil
}

func (b BaseStringServiceImpl) FindByIdentifierAndLanguage(identifier string, isoCode string, user *domain.User) (*string, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	language, err := b.languageRepository.FindByIsoCode(isoCode)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseString, err := b.baseStringRepository.FindByIdentifierAndLanguage(identifier, isoCode)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	err = b.checkPermission(*baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	translationContent := baseString.FindTranslationLastVersionByLanguage(*language)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &translationContent, nil
}

func (b BaseStringServiceImpl) FindByLanguage(id primitive.ObjectID, user *domain.User) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	_, err := b.languageRepository.FindById(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	baseStrings, err := b.baseStringRepository.FindByLanguage(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	for _, baseString := range *baseStrings {
		err = b.checkPermission(baseString, *user)
		if err != nil {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseStrings, err := b.baseStringRepository.FindByPermissions(id)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return baseStrings, nil
}

func (b BaseStringServiceImpl) Read(xliff xmlDto.Xliff, user *domain.User, stageId primitive.ObjectID, groupId primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	sourceLanguage, err := b.languageRepository.FindByIsoCode(xliff.SrcLang)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	targetLanguage, err := b.languageRepository.FindByIsoCode(xliff.TrgLang)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	stage, err := b.stageRepository.FindById(stageId)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	group, err := b.groupRepository.FindById(groupId)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	var baseStrings []domain.BaseString
	for _, unit := range xliff.FileXml.Units {
		baseString, err := b.baseStringRepository.FindByIdentifier(unit.Id)
		if err != nil && err.Error() != constants.FindBaseStringByIdentifier {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		sourceTranslation := domain.Translation{
			Content:  unit.Segment.Source,
			Language: sourceLanguage,
			Version:  1,
			Active:   true,
			Author:   user,
			Date:     time.Now(),
			Stage:    stage,
		}
		targetTranslation := domain.Translation{
			Content:  unit.Segment.Target,
			Language: targetLanguage,
			Version:  1,
			Active:   true,
			Author:   user,
			Date:     time.Now(),
			Stage:    stage,
		}
		if baseString == nil {
			newBaseString := domain.BaseString{
				SourceLanguage: sourceLanguage,
				Identifier:     unit.Id,
				Group:          group,
				Author:         user,
				Active:         true,
				Translations: []domain.Translation{
					sourceTranslation,
					targetTranslation,
				},
			}
			newBaseString, err := b.Create(newBaseString, user)
			if err != nil {
				return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
			}
			baseStrings = append(baseStrings, newBaseString)
		} else {
			b.checkExistTranslation(baseString, sourceTranslation)
			b.checkExistTranslation(baseString, targetTranslation)
			_, err = b.Update(*baseString, user)
			if err != nil {
				return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
			}
			baseStrings = append(baseStrings, *baseString)
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseStrings, nil
}

func (b BaseStringServiceImpl) Update(baseString domain.BaseString, user *domain.User) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := b.checkPermission(baseString, *user)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	original, err := b.baseStringRepository.FindById(baseString.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	if original.Identifier != baseString.Identifier {
		if baseString.Identifier == "" {
			return nil, tools.ErrorLog(constants.IdentifierBaseStringInvalid, tools.GetCurrentFuncName())
		}
		_, err := b.baseStringRepository.FindByIdentifier(baseString.Identifier)
		if err != nil && err.Error() != constants.FindBaseStringByIdentifier {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	_, err = b.baseStringRepository.Update(baseString)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b BaseStringServiceImpl) Write(xliffDto dto.XliffDto) (*xmlDto.Xliff, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	sourceLanguage, err := b.checkLanguageById(xliffDto.SourceLanguageId)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	targetLanguage, err := b.checkLanguageById(xliffDto.TargetLanguageId)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	if xliffDto.BaseStringIds == nil {
		return nil, errors.New(constants.BaseStringIdsNoValid)
	}

	var units []xmlDto.Unit
	for _, id := range xliffDto.BaseStringIds {
		sourceId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, errors.New(constants.IdNoValid)
		}
		baseString, err := b.FindById(sourceId)
		if err != nil {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		units = append(units, xmlDto.Unit{
			XMLName: xml.Name{},
			Id:      baseString.Identifier,
			Segment: xmlDto.Segment{
				XMLName: xml.Name{},
				Source:  baseString.FindTranslationLastVersionByLanguage(*sourceLanguage),
				Target:  baseString.FindTranslationLastVersionByLanguage(*targetLanguage),
			},
		})
	}
	result := xmlDto.Xliff{
		XMLName: xml.Name{},
		FileXml: xmlDto.FileXml{
			XMLName: xml.Name{},
			Units:   units,
		},
		Version: "2.0",
		SrcLang: sourceLanguage.IsoCode,
		TrgLang: targetLanguage.IsoCode,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &result, nil
}

func (b BaseStringServiceImpl) checkExistTranslation(baseString *domain.BaseString, translation domain.Translation) {
	count := 1
	for _, originalTranslation := range baseString.Translations {
		if originalTranslation.Stage.ID == translation.Stage.ID && originalTranslation.Language.ID == translation.Language.ID {
			count += 1
		}
	}
	translation.Version = count
	baseString.Translations = append(baseString.Translations, translation)
}

func (b BaseStringServiceImpl) checkLanguageById(id string) (*domain.Language, error) {
	sourceId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(constants.IdNoValid)
	}
	language, err := b.languageRepository.FindById(sourceId)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	return language, nil
}

func (b BaseStringServiceImpl) checkUniqueIdentifier(identifier string) error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if identifier == "" {
		return tools.ErrorLog(constants.IdentifierBaseStringInvalid, tools.GetCurrentFuncName())
	}
	baseString, err := b.baseStringRepository.FindByIdentifier(identifier)
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
		if translation.Author == nil {
			return errors.New(constants.TranslationNullAuthor)
		}
		email := translation.Author.Email
		_, err := b.userRepository.FindByEmail(email)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		isoCode := translation.Language.IsoCode
		if translation.Language == nil {
			return errors.New(constants.TranslationNullLanguage)
		}
		_, err = b.languageRepository.FindByIsoCode(isoCode)
		if err != nil {
			return tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
		if translation.Stage == nil {
			return errors.New(constants.TranslationNullStage)
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
