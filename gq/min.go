package gq

import (
	"context"
	"fmt"

	"github.com/petapedia/geoquery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MinDistance(mongoconn *mongo.Database, long float64, lat float64, mindistance float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("petapedia")
	filter := bson.M{
		"batas": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$minDistance": mindistance,
			},
		},
	}
	var lokasi models.Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		fmt.Printf("GetLokasi: %v\n", err)
	}
	return lokasi.Nama
}
