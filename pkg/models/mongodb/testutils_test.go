package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"module-7/pkg/models"
	"testing"
)

func newTestDB(t *testing.T) (*mongo.Database, func()) {
	mongoUri := "mongodb://root:rootpassword@localhost:27017"

	clientOptions := options.
		Client().
		ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	database := client.Database("test-local")
	database.CreateCollection(context.TODO(), "materials-local")
	collection := database.Collection("materials-local")

	material1 := models.Material{Name: "Some name for testing #1"}
	material2 := models.Material{Name: "Some name for testing #2"}

	var materials []interface{}
	materials = append(materials, material1, material2)

	collection.InsertMany(context.TODO(), materials)

	return database, func() {
		filter := bson.M{}
		collection.DeleteMany(context.TODO(), filter)
	}
}
