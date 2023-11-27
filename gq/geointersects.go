package gq

import (
	"context"
	"fmt"

	"github.com/petapedia/geoquery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// {
// 	"type": "Feature",
// 	"geometry": {
// 	  "type": "Point",
// 	  "coordinates": [
// 		107.66152981234228,
// 		-6.916179557176008
// 	  ]
// 	},
// 	"properties": {
// 	  "name": "Bahari Selera Nusantara"
// 	}
//   }
// }

func GeoIntersects(mongoconn *mongo.Database, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("petapedia")
	filter := bson.M{
		"batas": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
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
