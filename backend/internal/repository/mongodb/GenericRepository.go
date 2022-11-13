package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/tools"
)

type ConfigRepository struct {
	name                 string
	createErrorMessage   string
	findByIdErrorMessage string
	deleteErrorMessage   string
}

type GenericRepository[T any] struct {
	client     *mongo.Client
	Collection *mongo.Collection
	Config     ConfigRepository
}

func (g *GenericRepository[T]) Create(t T) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, g.Config.createErrorMessage, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *UserRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, u.Config.deleteErrorMessage, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (g *GenericRepository[T]) FindAll() (*[]T, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var elements []T
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var element T
		if err = cursor.Decode(&element); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		elements = append(elements, element)
	}
	if err = cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err = cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &elements, nil
}

func (g *GenericRepository[T]) FindById(id primitive.ObjectID) (*T, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var t T
	if err = result.Decode(&t); err != nil {
		return nil, tools.ErrorLogDetails(err, g.Config.findByIdErrorMessage, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &t, nil
}

func (g *GenericRepository[T]) getCollection() (*mongo.Collection, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if g.Collection == nil {
		tools.LoadEnv()
		err := g.createConnection()
		if err != nil {
			return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
		}
		var databaseName = os.Getenv("DATABASE_NAME")
		database := g.client.Database(databaseName)
		g.Collection = database.Collection(g.Config.name)
		log.Printf("%s: end", tools.GetCurrentFuncName())
	}
	return g.Collection, nil
}

func (g *GenericRepository[T]) createConnection() error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	client, err := g.connectDatabase()
	if err != nil {
		return err
	}
	g.client = client
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (g *GenericRepository[T]) connectDatabase() (*mongo.Client, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	uri := os.Getenv("ATLAS_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return client, nil
}
