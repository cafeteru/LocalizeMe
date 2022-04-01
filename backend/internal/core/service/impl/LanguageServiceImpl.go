package impl

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type LanguageServiceImpl struct {
	repository repository.LanguageRepository
}

func CreateLanguageService(r repository.LanguageRepository) *LanguageServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &LanguageServiceImpl{r}
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

func (l LanguageServiceImpl) FindAll() (*[]domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	users, err := l.repository.FindAll()
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return users, nil
}
