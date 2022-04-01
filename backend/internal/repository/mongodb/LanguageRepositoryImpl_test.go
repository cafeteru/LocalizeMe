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
	mt.Run("Create_Stage_Success", func(mt *mtest.T) {
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
	mt.Run("Create_Stage_NotConnection", func(mt *mtest.T) {
		_, err := l.Create(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestLanguageRepositoryImpl_Create_Error(t *testing.T) {
	mt, l := createLanguageMocks(t)
	mt.Run("Create_Stage_ErrorCreate", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertStage,
		}))
		_, err := l.Create(language)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertLanguage))
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
		stages, err := l.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, stages)
		assert.Equal(t, len(*stages), 2)
		assert.Equal(t, (*stages)[0].IsoCode, language.IsoCode)
		assert.Equal(t, (*stages)[1].IsoCode, language2.IsoCode)
	})
}

func TestLanguageRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, u := createLanguageMocks(t)
	mt.Run("FindAll_Language_NotConnect", func(mt *mtest.T) {
		_, err := u.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
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

func createLanguageMocks(t *testing.T) (*mtest.T, *LanguageRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	repository := CreateLanguageRepository()
	language = domain.Language{
		ID:          primitive.ObjectID{},
		IsoCode:     "IsoCode",
		Description: "Description",
		Active:      true,
	}
	return mt, repository
}
