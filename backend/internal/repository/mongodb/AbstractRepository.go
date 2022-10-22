package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/tools"
)

type Repository interface {
	GetCollection() (*mongo.Collection, error)
	CloseConnection()
}

type AbstractRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (a *AbstractRepository) GetCollection(name string) (*mongo.Collection, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if a.collection == nil {
		tools.LoadEnv()
		err := a.createConnection()
		if err != nil {
			return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
		}
		var databaseName = os.Getenv("DATABASE_NAME")
		database := a.client.Database(databaseName)
		a.collection = database.Collection(name)
		log.Printf("%s: end", tools.GetCurrentFuncName())
	}
	return a.collection, nil
}

func (a *AbstractRepository) createConnection() error {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	client, err := a.connectDatabase()
	if err != nil {
		return err
	}
	a.client = client
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (a *AbstractRepository) connectDatabase() (*mongo.Client, error) {
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
