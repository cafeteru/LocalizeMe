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

var user = domain.User{
	ID:       primitive.NewObjectID(),
	Email:    "user@email.com",
	Password: "password",
	Admin:    false,
	Active:   true,
}

func TestUserRepositoryImpl_Delete_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Delete_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "DeletedCount", Value: 1},
		}))
		_, err := u.Delete(user.ID)
		assert.Nil(t, err)
	})
}

func TestUserRepositoryImpl_Delete_NotConnection(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Delete_User_NotConnection", func(mt *mtest.T) {
		_, err := u.Delete(user.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_Delete_NotFound(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Delete_User_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.DeleteUser,
		}))
		_, err := u.Delete(user.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.DeleteUser))
	})
}

func TestUserRepositoryImpl_FindByEmail_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindByEmail_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
			{Key: "admin", Value: user.Admin},
			{Key: "active", Value: user.Active},
		}))
		response, err := u.FindByEmail(user.Email)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, response.Email)
	})
}

func TestUserRepositoryImpl_FindByEmail_NotConnection(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindByEmail_User_NotConnection", func(mt *mtest.T) {
		_, err := u.FindByEmail("email")
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_FindByEmail_NotFound(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindByEmail_User_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindUserByEmail,
		}))
		_, err := u.FindByEmail(user.Email)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindUserByEmail))
	})
}

func TestUserRepositoryImpl_FindById_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindById_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
			{Key: "admin", Value: user.Admin},
			{Key: "active", Value: user.Active},
		}))
		response, err := u.FindById(user.ID)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, response.Email)
	})
}

func TestUserRepositoryImpl_FindById_NotConnection(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindById_User_NotConnection", func(mt *mtest.T) {
		_, err := u.FindById(user.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_FindById_NotFound(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindById_User_NotFound", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.FindUserById,
		}))
		_, err := u.FindById(user.ID)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.FindUserById))
	})
}

func TestUserRepositoryImpl_Create_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Create_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
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
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.InsertUser,
		}))
		_, err := u.Create(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.InsertUser))
	})
}

func TestUserRepositoryImpl_FindAll_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("FindAll_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		user2 := domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "john2.doe@test.com",
			Password: "",
			Admin:    false,
			Active:   false,
		}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
			{Key: "admin", Value: user.Admin},
			{Key: "active", Value: user.Active},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: user2.ID},
			{Key: "email", Value: user2.Email},
			{Key: "password", Value: user2.Password},
			{Key: "admin", Value: user2.Admin},
			{Key: "active", Value: user2.Active},
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
	mt, u := createUserMocks(t)
	mt.Run("FindAll_User_Success", func(mt *mtest.T) {
		_, err := u.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_Update_Success(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Update_User_Success", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "MatchedCount", Value: 0},
			{Key: "ModifiedCount", Value: 1},
			{Key: "UpsertedCount", Value: 0},
			{Key: "UpsertedID", Value: user.ID},
		}))
		_, err := u.Update(user)
		assert.Nil(t, err)
	})
}

func TestUserRepositoryImpl_Update_NotConnection(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Update_User_NotConnection", func(mt *mtest.T) {
		_, err := u.Update(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.CreateConnection))
	})
}

func TestUserRepositoryImpl_Update_Error(t *testing.T) {
	mt, u := createUserMocks(t)
	mt.Run("Update_User_Error", func(mt *mtest.T) {
		u.collection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: constants.UpdateUser,
		}))
		_, err := u.Update(user)
		assert.NotNil(t, err)
		assert.Equal(t, err, errors.New(constants.UpdateUser))
	})
}

func createUserMocks(t *testing.T) (*mtest.T, *UserRepositoryImpl) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	u := CreateUserRepository()
	return mt, u
}
