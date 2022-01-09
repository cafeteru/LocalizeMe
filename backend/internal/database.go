package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"path/filepath"
	"time"
)

// Example how to create a user in MongoDB
func main() {
	loadEnv()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectDatabase(ctx)
	addTestData(client, ctx)
	defer closeConnection(client, ctx)
}

func loadEnv() {
	environmentPath := filepath.Join("./", ".env")
	_ = godotenv.Load(environmentPath)
}

func connectDatabase(ctx context.Context) *mongo.Client {
	var certificate = os.Getenv("CERTIFICATE")
	var uri = os.Getenv("ATLAS_URI") + certificate

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func addTestData(client *mongo.Client, ctx context.Context) {
	database := client.Database("localizeMe")
	userCollection := database.Collection("users")

	user := domain.User{
		Name:     "admin",
		Password: "admin",
		Role:     "admin",
		Active:   true,
	}
	user, _ = domain.User.EncryptPassword(user)
	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func closeConnection(client *mongo.Client, ctx context.Context) {
	func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(client, ctx)
}
