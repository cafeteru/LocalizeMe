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

var baseString domain.BaseString

func TestBaseStringRepositoryImpl_Create_Success(t *testing.T) {
	mt, l := createBaseStringMocks(t)
	mt.Run("Create_BaseString_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := l.Create(baseString)
		assert.Nil(t, err)
	})
}

func TestBaseStringRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, l := createBaseStringMocks(t)
	mt.Run("Create_BaseString_NotConnection", func(mt *mtest.T) {
		_, err := l.Create(baseString)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestBaseStringRepositoryImpl_Create_Error(t *testing.T) {
	mt, l := createBaseStringMocks(t)
	mt.Run("Create_BaseString_ErrorCreate", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertStage,
		}))
		_, err := l.Create(baseString)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertBaseString))
	})
}

func createBaseStringMocks(t *testing.T) (*mtest.T, *BaseStringRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	repository := CreateBaseStringRepository()
	baseString = domain.BaseString{
		ID:             primitive.ObjectID{},
		SourceLanguage: &language,
		Identifier:     "description of string",
		Group:          &group,
		Author:         &user,
		Active:         true,
		Translations:   nil,
	}
	return mt, repository
}
