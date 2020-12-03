package database

import (
	"context"
	"fmt"
	"time"
	"os"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sageflow/sageflow-backend/internal/graphql/model"
)

// Connect connects to specified mongoDB URI.
func Connect(mongodbURI string) *DB {
	fmt.Printf("Mongdb URI [Connect] (%v)\n", os.Getenv("MONGODB_URI"))
	mongodbURI = "mongodb://test:test@localhost:27017/test"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatalf("MongoDB Client Creation Failed: Error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("MongoDB Connection Failed: Error: %v", err)
	}

	return &DB{client}
}

// DB holds a mongo client.
type DB struct {
	client *mongo.Client
}

// FindUser gets the first user it finds in the db.
func (db *DB) FindUser() *model.User {
	collection := db.client.Database("test").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user *model.User

	err := collection.FindOne(ctx, bson.D{}).Decode(&user)
	if err != nil {
		log.Fatalf("MongoDB Failed to find User: Error: %v", err)
	}

	return user
}
