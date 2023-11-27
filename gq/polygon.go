package gq

import (
	"context"
	"log"

	"github.com/raulmahya123/GeospatialQueryOperators/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create2DSphereIndexGeo creates a 2dsphere index on the "geometry" field of the given collection.
func Create2DSphereIndexGeo(lokasicollection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"geometry": "2dsphere"},
	}

	_, err := lokasicollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}

// Polygon performs a $geoWithin query with a Polygon on the "polygon" collection.
func Polygon(mongoconn *mongo.Database, coordinates [][][]float64) []models.Lokasi {
	lokasicollection := mongoconn.Collection("polygon")

	// Ensure that the 2dsphere index is created (should be done once, not in every function call)
	// Uncomment the next line if you haven't created the index yet
	Create2DSphereIndexGeo(lokasicollection)

	filter := bson.M{
		"geometry": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates,
				},
			},
		},
	}

	// Use Find to retrieve multiple documents
	cursor, err := lokasicollection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Polygon: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	var lokasis []models.Lokasi
	for cursor.Next(context.TODO()) {
		var lokasi models.Lokasi
		if err := cursor.Decode(&lokasi); err != nil {
			log.Printf("Error decoding document: %v\n", err)
			continue
		}
		lokasis = append(lokasis, lokasi)
	}

	return lokasis
}
