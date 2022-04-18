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

func TestGroupRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("FindAll_Group_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		group2 := domain.Group{
			ID:          primitive.ObjectID{},
			Name:        "group2",
			Owner:       user,
			Permissions: []domain.Permission{},
			Active:      true,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: group.ID},
			{Key: "name", Value: group.Name},
			{Key: "owner", Value: group.Owner},
			{Key: "permissions", Value: group.Permissions},
			{Key: "active", Value: group.Active},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: group2.ID},
			{Key: "name", Value: group2.Name},
			{Key: "owner", Value: group2.Owner},
			{Key: "permissions", Value: group2.Permissions},
			{Key: "active", Value: group2.Active},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		Groups, err := l.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, Groups)
		assert.Equal(t, len(*Groups), 2)
		assert.Equal(t, (*Groups)[0].Name, group.Name)
		assert.Equal(t, (*Groups)[1].Name, group2.Name)
	})
}

func TestGroupRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("FindAll_Group_NotConnect", func(mt *mtest.T) {
		_, err := l.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestGroupRepositoryImpl_FindByName_Success(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("FindByName_Group_Success", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: group.ID},
			{Key: "name", Value: group.Name},
			{Key: "owner", Value: group.Owner},
			{Key: "permissions", Value: group.Permissions},
			{Key: "active", Value: group.Active},
		}))
		response, err := l.FindByName(group.Name)
		assert.Nil(t, err)
		assert.Equal(t, group.Name, response.Name)
	})
}

func TestGroupRepositoryImpl_FindByName_NotConnection(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("FindByName_Group_NotConnection", func(mt *mtest.T) {
		_, err := l.FindByName("Name")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestGroupRepositoryImpl_FindByName_NotFound(t *testing.T) {
	mt, l := createGroupMocks(t)
	mt.Run("FindByName_Group_NotFound", func(mt *mtest.T) {
		l.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindGroupByName,
		}))
		_, err := l.FindByName(group.Name)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindGroupByName))
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
