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

type StageServiceImpl struct {
	repository repository.StageRepository
}

func CreateStageService() *StageServiceImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	service := &StageServiceImpl{mongodb.CreateStageRepository()}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return service
}

func (s StageServiceImpl) Create(stageDto dto.StageDto) (domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	findByName, errName, validName := s.checkUniqueName(stageDto.Name)
	if !validName {
		return findByName, tools.ErrorLogWithError(errName, tools.GetCurrentFuncName())
	}
	stage := domain.Stage{
		Name:   stageDto.Name,
		Active: true,
	}
	resultId, err := s.repository.Create(stage)
	if err != nil {
		log.Printf("%s: error", tools.GetCurrentFuncName())
		return domain.Stage{}, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	stage.ID = resultId.InsertedID.(primitive.ObjectID)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return stage, nil
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
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return users, nil
}

func (s StageServiceImpl) Update(stage domain.Stage) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	original, err := s.repository.FindById(stage.ID)
	if original == nil || err != nil {
		return nil, tools.ErrorLog(constants.FindStageById, tools.GetCurrentFuncName())
	}
	if original.Name != stage.Name {
		_, errName, validName := s.checkUniqueName(stage.Name)
		if !validName {
			return nil, tools.ErrorLogWithError(errName, tools.GetCurrentFuncName())
		}
	}
	_, err = s.repository.Update(stage)
	if err != nil {
		return nil, tools.ErrorLogWithError(err, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}

func (s StageServiceImpl) checkUniqueName(name string) (domain.Stage, error, bool) {
	if name == "" {
		return domain.Stage{}, tools.ErrorLog(constants.NameStageInvalid, tools.GetCurrentFuncName()), false
	}
	findByName, _ := s.repository.FindByName(name)
	if findByName != nil {
		return domain.Stage{}, tools.ErrorLog(constants.StageAlreadyRegister, tools.GetCurrentFuncName()), false
	}
	return domain.Stage{}, nil, true
}
