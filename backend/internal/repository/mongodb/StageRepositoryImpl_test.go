package mongodb

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

var stage = domain.Stage{
	ID:     primitive.NewObjectID(),
	Name:   "name",
	Active: true,
}

func TestStageRepositoryImpl_Create_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Create_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := s.Create(stage)
		assert.Nil(t, err)
	})
}

func TestStageRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Create_Stage_NotConnection", func(mt *mtest.T) {
		_, err := s.Create(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Create_Error(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Create_Stage_ErrorCreate", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertStage,
		}))
		_, err := s.Create(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertStage))
	})
}

func TestStageRepositoryImpl_Delete_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Delete_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "DeletedCount", Value: 1},
		}))
		_, err := s.Delete(stage.ID)
		assert.Nil(t, err)
	})
}

func TestStageRepositoryImpl_Delete_NotConnection(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Delete_Stage_NotConnection", func(mt *mtest.T) {
		_, err := s.Delete(stage.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Delete_NotFound(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Delete_Stage_NotFound", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.DeleteStage,
		}))
		_, err := s.Delete(stage.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.DeleteStage))
	})
}

func TestStageRepositoryImpl_FindById_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindById_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "name", Value: stage.Name},
			{Key: "active", Value: stage.Active},
		}))
		response, err := s.FindById(stage.ID)
		assert.Nil(t, err)
		assert.Equal(t, stage.Name, response.Name)
	})
}

func TestStageRepositoryImpl_FindById_NotConnection(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindById_Stage_NotConnection", func(mt *mtest.T) {
		_, err := s.FindById(primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_FindById_NotFound(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindById_Stage_NotFound", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindStageById,
		}))
		_, err := s.FindById(stage.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindStageById))
	})
}

func TestStageRepositoryImpl_FindByName_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindByName_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "name", Value: stage.Name},
			{Key: "active", Value: stage.Active},
		}))
		response, err := s.FindByName(stage.Name)
		assert.Nil(t, err)
		assert.Equal(t, stage.Name, response.Name)
	})
}

func TestStageRepositoryImpl_FindByName_NotConnection(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindByName_Stage_NotConnection", func(mt *mtest.T) {
		_, err := s.FindByName("email")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_FindByName_NotFound(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindByName_Stage_NotFound", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindStageByName,
		}))
		_, err := s.FindByName(stage.Name)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindStageByName))
	})
}

func TestStageRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindAll_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		stage2 := domain.Stage{
			ID:     primitive.NewObjectID(),
			Name:   "name2",
			Active: true,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "name", Value: stage.Name},
			{Key: "active", Value: stage.Active},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: stage2.ID},
			{Key: "name", Value: stage2.Name},
			{Key: "active", Value: stage2.Active},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		stages, err := s.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, stages)
		assert.Equal(t, len(*stages), 2)
		assert.Equal(t, (*stages)[0].Name, stage.Name)
		assert.Equal(t, (*stages)[1].Name, stage2.Name)
	})
}

func TestStageRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("FindAll_Stage_NotConnect", func(mt *mtest.T) {
		_, err := s.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Update_Success(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Update_Stage_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "MatchedCount", Value: 0},
			{Key: "ModifiedCount", Value: 1},
			{Key: "UpsertedCount", Value: 0},
			{Key: "UpsertedID", Value: stage.ID},
		}))
		_, err := s.Update(stage)
		assert.Nil(t, err)
	})
}

func TestStageRepositoryImpl_Update_NotConnection(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Update_Stage_NotConnection", func(mt *mtest.T) {
		_, err := s.Update(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Update_Error(t *testing.T) {
	mt, s := createStageMocks(t)
	mt.Run("Update_Stage_Error", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.UpdateStage,
		}))
		_, err := s.Update(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.UpdateStage))
	})
}

func createStageMocks(t *testing.T) (*mtest.T, *StageRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	repository := CreateStageRepository()
	return mt, repository
}
