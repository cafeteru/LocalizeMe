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

type StageServiceImpl struct {
	repository repository.StageRepository
}

func CreateStageService(r repository.StageRepository) *StageServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &StageServiceImpl{r}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (s StageServiceImpl) Create(request dto.StageRequest) (domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	name := request.Name
	if name == "" {
		return domain.Stage{}, tools.ErrorLog(constants.StageInvalid, tools.GetCurrentFuncName())
	}
	findByName, _ := s.repository.FindByName(name)
	if findByName != nil {
		return domain.Stage{}, tools.ErrorLog(constants.StageAlreadyRegister, tools.GetCurrentFuncName())
	}
	stage := domain.Stage{
		Name:   request.Name,
		Active: true,
	}
	resultId, err := s.repository.Create(stage)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.Stage{}, err
	}
	stage.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return stage, nil
}

func (s StageServiceImpl) FindAll() (*[]domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	users, err := s.repository.FindAll()
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return users, nil
}
