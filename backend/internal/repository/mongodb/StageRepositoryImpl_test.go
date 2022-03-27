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
	mt, u := createStageMocks(t)
	mt.Run("Create_Stage_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := u.Create(stage)
		assert.Nil(t, err)
	})
}

func TestStageRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("Create_Stage_NotConnection", func(mt *mtest.T) {
		_, err := u.Create(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Create_Error(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("Create_Stage_ErrorCreate", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertStage,
		}))
		_, err := u.Create(stage)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertStage))
	})
}

func TestStageRepositoryImpl_FindById_Success(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindById_Stage_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "Name", Value: stage.Name},
			{Key: "Active", Value: stage.Active},
		}))
		response, err := u.FindById(stage.ID)
		assert.Nil(t, err)
		assert.Equal(t, stage.Name, response.Name)
	})
}

func TestStageRepositoryImpl_FindById_NotConnection(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindById_Stage_NotConnection", func(mt *mtest.T) {
		_, err := u.FindById(primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_FindById_NotFound(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindById_Stage_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindStageById,
		}))
		_, err := u.FindById(stage.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindStageById))
	})
}

func TestStageRepositoryImpl_FindByName_Success(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindByName_Stage_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "Name", Value: stage.Name},
			{Key: "Active", Value: stage.Active},
		}))
		response, err := u.FindByName(stage.Name)
		assert.Nil(t, err)
		assert.Equal(t, stage.Name, response.Name)
	})
}

func TestStageRepositoryImpl_FindByName_NotConnection(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindByName_Stage_NotConnection", func(mt *mtest.T) {
		_, err := u.FindByName("email")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_FindByName_NotFound(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindByName_Stage_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindStageByName,
		}))
		_, err := u.FindByName(stage.Name)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindStageByName))
	})
}

func TestStageRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindAll_Stage_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		stage2 := domain.Stage{
			ID:     primitive.NewObjectID(),
			Name:   "name2",
			Active: true,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: stage.ID},
			{Key: "Name", Value: stage.Name},
			{Key: "Active", Value: stage.Active},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: stage2.ID},
			{Key: "Name", Value: stage2.Name},
			{Key: "Active", Value: stage2.Active},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		stages, err := u.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, stages)
		assert.Equal(t, len(*stages), 2)
		assert.Equal(t, (*stages)[0].Name, stage.Name)
		assert.Equal(t, (*stages)[1].Name, stage2.Name)
	})
}

func TestStageRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("FindAll_Stage_Success", func(mt *mtest.T) {
		_, err := u.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestStageRepositoryImpl_Update_Success(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("Update_Stage_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "MatchedCount", Value: 0},
			{Key: "ModifiedCount", Value: 1},
			{Key: "UpsertedCount", Value: 0},
			{Key: "UpsertedID", Value: stage.ID},
		}))
		_, err := u.Update(stage)
		assert.Nil(t, err)
	})
}

func TestRepositoryImpl_Update_NotConnection(t *testing.T) {
	mt, u := createStageMocks(t)
	mt.Run("Update_Stage_NotConnection", func(mt *mtest.T) {
		_, err := u.Update(stage)
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
	u := CreateStageRepository()
	return mt, u
}
