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

var group domain.Group

func TestGroupRepositoryImpl_Create_Success(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("Create_Group_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := l.Create(group)
		assert.Nil(t, err)
	})
}

func TestGroupRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("Create_Group_NotConnection", func(mt *mtest.T) {
		_, err := l.Create(group)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestGroupRepositoryImpl_Create_Error(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("Create_Group_ErrorCreate", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertStage,
		}))
		_, err := l.Create(group)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertGroup))
	})
}

func createGroupMocks(t *testing.T) (*mtest.T, *GroupRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	repository := CreateGroupRepository()
	group = domain.Group{
		ID:          primitive.ObjectID{},
		Name:        "group",
		Owner:       user,
		Permissions: []domain.Permission{},
		Active:      true,
	}
	return mt, repository
}
