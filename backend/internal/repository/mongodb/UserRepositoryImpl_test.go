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

func TestUserRepositoryImpl_FindByEmail_Success(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("FindByEmail_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		user := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "Email", Value: user.Email},
			{Key: "Password", Value: user.Password},
			{Key: "IsAdmin", Value: user.IsAdmin},
			{Key: "IsActive", Value: user.IsActive},
		}))
		response, err := u.FindByEmail(user.Email)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, response.Email)
	})
}

func TestUserRepositoryImpl_FindByEmail_NotConnection(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("FindByEmail_NotConnection", func(mt *mtest.T) {
		_, err := u.FindByEmail("email")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_FindByEmail_NotFound(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("FindByEmail_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		user := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindUserByEmail,
		}))
		_, err := u.FindByEmail(user.Email)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindUserByEmail))
	})
}

func TestUserRepositoryImpl_Create_Success(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("Create_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		user := domain.User{
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: " InsertedID", Value: primitive.NewObjectID()},
		}))
		_, err := u.Create(user)
		assert.Nil(t, err)
	})
}

func TestUserRepositoryImpl_Create_NotConnection(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("Create_NotConnection", func(mt *mtest.T) {
		user := domain.User{
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		_, err := u.Create(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_Create_ErrorCreate(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("Create_ErrorCreate", func(mt *mtest.T) {
		u.collection = mt.Coll
		user := domain.User{
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertUser,
		}))
		_, err := u.Create(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertUser))
	})
}

func TestUserRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("FindAll_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		user := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "john.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		user2 := domain.User{
			Email:    "john2.doe@test.com",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "Email", Value: user.Email},
			{Key: "Password", Value: user.Password},
			{Key: "IsAdmin", Value: user.IsAdmin},
			{Key: "IsActive", Value: user.IsActive},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: user2.ID},
			{Key: "Email", Value: user2.Email},
			{Key: "Password", Value: user2.Password},
			{Key: "IsAdmin", Value: user2.IsAdmin},
			{Key: "IsActive", Value: user2.IsActive},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		users, err := u.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, len(*users), 2)
		assert.Equal(t, (*users)[0].Email, user.Email)
		assert.Equal(t, (*users)[1].Email, user2.Email)
	})
}

func TestUserRepositoryImpl_FindAll_NotConnect(t *testing.T) {
	mt, u := createMocks(t)
	mt.Run("FindAll_Success", func(mt *mtest.T) {
		_, err := u.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func createMocks(t *testing.T) (*mtest.T, *UserRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	u := CreateUserRepository()
	return mt, u
}
