package helper

import (
	"context"
	"log"

	"github.com/aiteung/atdb"
	"github.com/petapedia/geoquery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongostring, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: mongostring,
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func SetConnection2dsphereTest(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: mongoenv,
		DBName:   dbname,
	}
	db := atdb.MongoConnect(DBmongoinfo)

	// Create a geospatial index if it doesn't exist
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "geometry", Value: "2dsphere"},
		},
	}

	_, err := db.Collection("near").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Printf("Error creating geospatial index: %v\n", err)
	}
	return db
}

func SetConnection2dsphereTestPoint(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: mongoenv,
		DBName:   dbname,
	}
	db := atdb.MongoConnect(DBmongoinfo)

	// Create a geospatial index if it doesn't exist
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "geometry", Value: "2dsphere"},
		},
	}

	_, err := db.Collection("nearspehere").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Printf("Error creating geospatial index: %v\n", err)
	}

	return db
}

func SetConnection2dsphereTestGeo(mongoenv, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: mongoenv,
		DBName:   dbname,
	}
	db := atdb.MongoConnect(DBmongoinfo)

	// Create a geospatial index if it doesn't exist
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "geometry", Value: "2dsphere"},
		},
	}

	_, err := db.Collection("polygon").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Printf("Error creating geospatial index: %v\n", err)
	}
	return db
}
func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []models.GeoJson {
	lokasi := atdb.GetAllDoc[[]models.GeoJson](mongoconn, collection)
	return lokasi
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata models.User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[models.User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}
