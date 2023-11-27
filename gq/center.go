package gq

import (
	"context"
	"fmt"

	"github.com/petapedia/geoquery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// geoWithin
//
//	{
//		<location field>: {
//		   $geoWithin: { $center: [ [ <x>, <y> ] , <radius> ] }
//		}
//	 }
func Center(mongoconn *mongo.Database, long float64, lat float64, radius float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("petapedia")
	filter := bson.M{
		"batas": bson.M{
			"$geoWithin": bson.M{
				"$center": [][]float64{{long, lat}, {radius}},
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
