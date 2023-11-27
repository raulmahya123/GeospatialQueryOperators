package gq

import (
	"context"
	"log"

	"github.com/raulmahya123/GeospatialQueryOperators/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Near(mongoconn *mongo.Database, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("near")
	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "LineString",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": 1000,
			},
		},
	}
	var lokasi models.Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("Near: %v\n", err)
	}
	return lokasi.Properties.Name
}
