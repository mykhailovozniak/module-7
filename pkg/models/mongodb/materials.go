package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"module-7/pkg/models"
)

type MaterialsModel struct {
	Collection *mongo.Collection
}

func (m *MaterialsModel) FindAll() (mat []*models.Material, err error) {
	findOptions := options.Find()
	filter := bson.M{}
	cur, err := m.Collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	var materials []*models.Material

	for cur.Next(context.TODO()) {
		var material models.Material
		err := cur.Decode(&material)
		if err != nil {
			log.Fatal(err)
		}

		materials = append(materials, &material)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return materials, nil
}
