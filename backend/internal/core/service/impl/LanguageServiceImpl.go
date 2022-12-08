package impl

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/internal/repository"
	"uniovi-localizeme/internal/repository/mongodb"
	"uniovi-localizeme/tools"
)

type LanguageServiceImpl struct {
	repository repository.LanguageRepository
}

func CreateLanguageService() *LanguageServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &LanguageServiceImpl{mongodb.CreateLanguageRepository()}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (l LanguageServiceImpl) Create(request dto.LanguageDto) (domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	findByName, errName, validName := l.checkUniqueIsoCode(request.IsoCode)
	if !validName {
		return findByName, errName
	}
	language := domain.Language{
		IsoCode:     request.IsoCode,
		Description: request.Description,
		Active:      true,
	}
	resultId, err := l.repository.Create(language)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.Language{}, err
	}
	language.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return language, nil
}

func (l LanguageServiceImpl) checkUniqueIsoCode(isoCode string) (domain.Language, error, bool) {
	if isoCode == "" {
		return domain.Language{}, tools.ErrorLog(constants.IsoCodeLanguageInvalid, tools.GetCurrentFuncName()), false
	}
	language, _ := l.repository.FindByIsoCode(isoCode)
	if language != nil {
		return domain.Language{}, tools.ErrorLog(constants.LanguageAlreadyRegister, tools.GetCurrentFuncName()), false
	}
	return domain.Language{}, nil, true
}

func (l LanguageServiceImpl) Delete(id primitive.ObjectID) (bool, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	language, err := l.repository.FindById(id)
	if language == nil || err != nil {
		return false, tools.ErrorLog(constants.FindLanguageById, tools.GetCurrentFuncName())
	}
	_, err = l.repository.Delete(id)
	if err != nil {
		return false, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (l LanguageServiceImpl) Disable(id primitive.ObjectID) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	language, err := l.repository.FindById(id)
	if language == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindLanguageById, tools.GetCurrentFuncName())
	}
	language.Active = !language.Active
	_, err = l.repository.Update(*language)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return language, nil
}

func (l LanguageServiceImpl) FindAll() (*[]domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	languages, err := l.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return languages, nil
}

func (l LanguageServiceImpl) Update(language domain.Language) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	original, err := l.repository.FindById(language.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindLanguageById, tools.GetCurrentFuncName())
	}
	if original.IsoCode != language.IsoCode {
		_, errName, validName := l.checkUniqueIsoCode(language.IsoCode)
		if !validName {
			return nil, tools.ErrorLogWithError(errName, tools.GetCurrentFuncName())
		}
	}
	_, err = l.repository.Update(language)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &language, nil
}
