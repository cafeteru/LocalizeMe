package mongodb

import (
	"context"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Repository interface {
	GetCollection() (*mongo.Collection, error)
	CloseConnection()
}

type AbstractRepository struct {
	Client *mongo.Client
	Repository
}

func (r *AbstractRepository) GetCollection(name string) (*mongo.Collection, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	tools.LoadEnv()
	err := r.createConnection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var databaseName = os.Getenv("DATABASE_NAME")
	database := r.Client.Database(databaseName)
	collection := database.Collection(name)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return collection, nil
}

func (r *AbstractRepository) createConnection() error {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	client, err := r.connectDatabase()
	if err != nil {
		return err
	}
	r.Client = client
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return nil
}

func (r *AbstractRepository) connectDatabase() (*mongo.Client, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	uri := os.Getenv("ATLAS_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return client, nil
}

func (r *AbstractRepository) CloseConnection() {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	defer r.disconnectDatabase()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

func (r *AbstractRepository) disconnectDatabase() {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	err := r.Client.Disconnect(context.TODO())
	if err != nil {
		slog.Errorf("%s: %s", tools.GetCurrentFuncName(), err)
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}
