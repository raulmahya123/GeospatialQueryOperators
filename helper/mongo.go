package helper

import (
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

func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []models.GeoJson {
	lokasi := atdb.GetAllDoc[[]models.GeoJson](mongoconn, collection)
	return lokasi
}

func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata models.User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[models.User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}
