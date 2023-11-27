package geoquery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/petapedia/geoquery/config"
	"github.com/petapedia/geoquery/helper"
	"github.com/raulmahya123/GeospatialQueryOperators/gq"
)

func PostGeoIntersects(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == config.EndpointSecret {
		resp.Response = gq.GeoIntersects(config.Mongocon, msg.Longitude, msg.Latitude)
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, helper.Json(resp))
}

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
