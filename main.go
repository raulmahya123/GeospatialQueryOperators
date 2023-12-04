package geoquery

import (
	"encoding/json"
	"net/http"

	"github.com/raulmahya123/GeospatialQueryOperators/gq"
	"github.com/raulmahya123/GeospatialQueryOperators/helper"
	"github.com/raulmahya123/GeospatialQueryOperators/models"
)

// func PostGeoIntersects(w http.ResponseWriter, r *http.Request) {
// 	var msg model.IteungMessage
// 	var resp atmessage.Response
// 	json.NewDecoder(r.Body).Decode(&msg)
// 	if r.Header.Get("Secret") == config.EndpointSecret {
// 		resp.Response = gq.GeoIntersects(config.Mongocon, msg.Longitude, msg.Latitude)
// 	} else {
// 		resp.Response = "Secret Salah"
// 	}
// 	fmt.Fprintf(w, helper.Json(resp))
// }

// func PostGeoWithin(w http.ResponseWriter, r *http.Request) {
// 	var msg model.IteungMessage
// 	var resp atmessage.Response
// 	json.NewDecoder(r.Body).Decode(&msg)
// 	if r.Header.Get("Secret") == config.EndpointSecret {
// 		resp.Response = gq.GeoWithin(config.Mongocon, msg.coordinates)
// 	} else {
// 		resp.Response = "Secret Salah"
// 	}
// 	fmt.Fprintf(w, helper.Json(resp))
// }

func PostGeoIntersects(mongoenv, dbname string, r *http.Request) string {
	var longlat models.LongLat
	var response models.Pesan
	response.Status = false
	mconn := helper.SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = gq.GeoIntersects(mconn, longlat.Longitude, longlat.Latitude)
	}
	return helper.ReturnStruct(response)
}

func PostGeoWithin(mongoenv, dbname string, r *http.Request) string {
	var coordinate models.GeometryPolygon
	var response models.Pesan
	response.Status = false
	mconn := helper.SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = gq.GeoWithin(mconn, coordinate.Coordinates)
	}
	return helper.ReturnStruct(response)
}

func PostNear(mongoenv, dbname string, r *http.Request) string {
	var longlat models.LongLat
	var response models.Pesan
	response.Status = false
	mconn := helper.SetConnection2dsphereTest(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = gq.Near(mconn, longlat.Longitude, longlat.Latitude)
	}
	return helper.ReturnStruct(response)
}
