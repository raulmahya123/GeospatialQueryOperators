package geoquery

import (
	"fmt"
	"testing"

	"github.com/petapedia/geoquery/helper"
	"github.com/raulmahya123/GeospatialQueryOperators/gq"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapedia")
	datagedung := gq.GeoIntersects(mconn, 107.66417814690686, -6.901637720654366)
	fmt.Println(datagedung)
}

/**
* Paste one or more documents here
 */
// {
// 	"type": "Feature",
// 	"properties": {
// 	  "name": "sd negeri 28"
// 	},
// 	"geometry": {
// 	  "coordinates": [
// 		[
// 		  [
// 			103.626550656458,
// 			-1.6150541661371705
// 		  ],
// 		  [
// 			103.62633759669836,
// 			-1.615292034441552
// 		  ],
// 		  [
// 			103.62602492459337,
// 			-1.6157235163933308
// 		  ],
// 		  [
// 			103.62610240086707,
// 			-1.6157788345851287
// 		  ],
// 		  [
// 			103.62662536571582,
// 			-1.6150790593407436
// 		  ],
// 		  [
// 			103.626550656458,
// 			-1.6150541661371705
// 		  ]
// 		]
// 	  ],
// 	  "type": "Polygon"
// 	}
//   }
