package mongodb

import (
	"context"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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

func (a *AbstractRepository) CloseConnection() {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	defer a.disconnectDatabase()
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func (a *AbstractRepository) disconnectDatabase() {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if a.client != nil {
		err := a.client.Disconnect(context.TODO())
		if err != nil {
			log.Printf("%s: %s", tools.GetCurrentFuncName(), err)
		}
	}
	a.clearClientCollection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func (a *AbstractRepository) clearClientCollection() {
	a.client = nil
	a.collection = nil
}
