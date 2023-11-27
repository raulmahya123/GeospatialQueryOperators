package gq

import (
	"context"
	"log"

	"github.com/raulmahya123/GeospatialQueryOperators/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//	$geometry: {
//		type: "<GeoJSON object type>",
//		coordinates: [ <coordinates> ]
//	 }
func Geometry(mongoconn *mongo.Database, coordinates [][][]float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("geometry")
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
	var lokasi models.Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		log.Printf("GeoWithin: %v\n", err)
	}
	return lokasi.Properties.Name
}
