package gq

import (
	"context"
	"fmt"
	"log"

	"github.com/raulmahya123/GeospatialQueryOperators/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Polygon(mongoconn *mongo.Database, coordinates [][][]float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("polygon")

	// Log coordinates for debugging
	fmt.Println("Coordinates:", coordinates)

	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$geometry": bson.M{
					"type":        "Polygon",
					"coordinates": coordinates,
				},
			},
		},
	}

	fmt.Println("Filter:", filter)

	var lokasi models.Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("Polygon: %v\n", err)
		return ""
	}

	return lokasi.Properties.Name
}
