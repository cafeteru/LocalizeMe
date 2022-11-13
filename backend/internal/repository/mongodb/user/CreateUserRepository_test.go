package mongodb

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/repository/mongodb"
)

var user = domain.User{
	ID:       primitive.NewObjectID(),
	Email:    "user@email.com",
	Password: "password",
	Admin:    false,
	Active:   true,
}

func TestUserRepositoryImpl_Create_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Create_User_Success", func(mt *mtest.T) {
		u.Collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := u.Create(user)
		assert.Nil(t, err)
	})
}

func TestUserRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Create_User_NotConnection", func(mt *mtest.T) {
		_, err := u.Create(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_Create_Error(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Create_User_ErrorCreate", func(mt *mtest.T) {
		u.Collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertUser,
		}))
		_, err := u.Create(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertUser))
	})
}

func createUserMocks(t *testing.T) (*mtest.T, *mongodb.UserRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	u := mongodb.CreateUserRepository()
	return mt, u
}
