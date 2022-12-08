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

type StageServiceImpl struct {
	repository repository.StageRepository
}

func CreateStageService() *StageServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &StageServiceImpl{mongodb.CreateStageRepository()}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (s StageServiceImpl) Create(stageDto dto.StageDto) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := s.checkUniqueName(stageDto.Name)
	if err != nil && err.Error() != constants.FindStageByName {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	stage := domain.Stage{
		Name:   stageDto.Name,
		Active: true,
	}
	resultId, err := s.repository.Create(stage)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	stage.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}

func (s StageServiceImpl) Delete(id primitive.ObjectID) (bool, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stage, err := s.repository.FindById(id)
	if stage == nil || err != nil {
		return false, tools.ErrorLog(constants.FindStageById, tools.GetCurrentFuncName())
	}
	_, err = s.repository.Delete(id)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return false, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return true, nil
}

func (s StageServiceImpl) Disable(id primitive.ObjectID) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stage, err := s.repository.FindById(id)
	if stage == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindStageById, tools.GetCurrentFuncName())
	}
	stage.Active = !stage.Active
	_, err = s.repository.Update(*stage)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return stage, nil
}

func (s StageServiceImpl) FindAll() (*[]domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stages, err := s.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return stages, nil
}

func (s StageServiceImpl) FindByName(name string) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stage, err := s.repository.FindByName(name)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return stage, nil
}

func (s StageServiceImpl) Update(stage domain.Stage) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	original, err := s.repository.FindById(stage.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindStageById, tools.GetCurrentFuncName())
	}
	if original.Name != stage.Name {
		err = s.checkUniqueName(stage.Name)
		if err != nil {
			return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
		}
	}
	_, err = s.repository.Update(stage)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}

func (s StageServiceImpl) checkUniqueName(name string) error {
	if name == "" {
		return tools.ErrorLog(constants.NameStageInvalid, tools.GetCurrentFuncName())
	}
	findByName, _ := s.repository.FindByName(name)
	if findByName != nil {
		return tools.ErrorLog(constants.StageAlreadyRegister, tools.GetCurrentFuncName())
	}
	return nil
}
