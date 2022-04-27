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

var language domain.Language

func TestLanguageRepositoryImpl_Create_Success(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Create_Language_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := l.Create(language)
		assert.Nil(t, err)
	})
}

func TestLanguageRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Create_Language_NotConnection", func(mt *mtest.T) {
		_, err := l.Create(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_Create_Error(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Create_Language_ErrorCreate", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertLanguage,
		}))
		_, err := l.Create(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertLanguage))
	})
}

func TestLanguageRepositoryImpl_Delete_Success(t *testing.T) {
	mt, s := createLanguageMocks(t)
	mt.Run("Delete_Language_Success", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "DeletedCount", Value: 1},
		}))
		_, err := s.Delete(language.ID)
		assert.Nil(t, err)
	})
}

func TestLanguageRepositoryImpl_Delete_NotConnection(t *testing.T) {
	mt, s := createLanguageMocks(t)
	mt.Run("Delete_Language_NotConnection", func(mt *mtest.T) {
		_, err := s.Delete(language.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_Delete_NotFound(t *testing.T) {
	mt, s := createLanguageMocks(t)
	mt.Run("Delete_Language_NotFound", func(mt *mtest.T) {
		s.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.DeleteLanguage,
		}))
		_, err := s.Delete(language.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.DeleteLanguage))
	})
}

func TestLanguageRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindAll_Language_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		language2 := domain.Language{
			ID:          primitive.NewObjectID(),
			IsoCode:     "isoCode",
			Description: "description",
			Active:      true,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: language.ID},
			{Key: "isoCode", Value: language.IsoCode},
			{Key: "description", Value: language.Description},
			{Key: "active", Value: language.Active},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: language2.ID},
			{Key: "isoCode", Value: language2.IsoCode},
			{Key: "description", Value: language2.Description},
			{Key: "active", Value: language2.Active},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		languages, err := l.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, languages)
		assert.Equal(t, len(*languages), 2)
		assert.Equal(t, (*languages)[0].IsoCode, language.IsoCode)
		assert.Equal(t, (*languages)[1].IsoCode, language2.IsoCode)
	})
}

func TestLanguageRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindAll_Language_NotConnect", func(mt *mtest.T) {
		_, err := l.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_FindById_Success(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindById_Language_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: language.ID},
			{Key: "description", Value: language.Description},
			{Key: "isoCode", Value: language.IsoCode},
			{Key: "active", Value: language.Active},
		}))
		response, err := l.FindById(language.ID)
		assert.Nil(t, err)
		assert.Equal(t, language.IsoCode, response.IsoCode)
	})
}

func TestLanguageRepositoryImpl_FindById_NotConnection(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindById_Language_NotConnection", func(mt *mtest.T) {
		_, err := l.FindById(primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_FindById_NotFound(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindById_Language_NotFound", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindLanguageById,
		}))
		_, err := l.FindById(language.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindLanguageById))
	})
}

func TestLanguageRepositoryImpl_FindByIsoCode_Success(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindByIsoCode_Language_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: language.ID},
			{Key: "isoCode", Value: language.IsoCode},
			{Key: "description", Value: language.Description},
			{Key: "active", Value: language.Active},
		}))
		response, err := l.FindByIsoCode(language.IsoCode)
		assert.Nil(t, err)
		assert.Equal(t, language.IsoCode, response.IsoCode)
	})
}

func TestLanguageRepositoryImpl_FindByIsoCode_NotConnection(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindByIsoCode_Language_NotConnection", func(mt *mtest.T) {
		_, err := l.FindByIsoCode("isoCode")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_FindByIsoCode_NotFound(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("FindByIsoCode_Language_NotFound", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindLanguageByIsoCode,
		}))
		_, err := l.FindByIsoCode(language.IsoCode)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindLanguageByIsoCode))
	})
}

func TestLanguageRepositoryImpl_Update_Success(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Update_Language_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "MatchedCount", Value: 0},
			{Key: "ModifiedCount", Value: 1},
			{Key: "UpsertedCount", Value: 0},
			{Key: "UpsertedID", Value: language.ID},
		}))
		_, err := l.Update(language)
		assert.Nil(t, err)
	})
}

func TestLanguageRepositoryImpl_Update_NotConnection(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Update_Language_NotConnection", func(mt *mtest.T) {
		_, err := l.Update(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_Update_Error(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Update_Language_Error", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.UpdateLanguage,
		}))
		_, err := l.Update(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.UpdateLanguage))
	})
}

func createLanguageMocks(t *testing.T) (*mtest.T, *LanguageRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	repository := CreateLanguageRepository()
	language = domain.Language{
		ID:          primitive.ObjectID{},
		IsoCode:     "IsoCode",
		Description: "Identifier",
		Active:      true,
	}
	return mt, repository
}
