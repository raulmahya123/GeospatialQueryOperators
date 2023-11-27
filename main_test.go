package geoquery

import (
	"fmt"
	"testing"

	"github.com/raulmahya123/GeospatialQueryOperators/gq"
	"github.com/raulmahya123/GeospatialQueryOperators/helper"
)

// func TestUpdateGetData(t *testing.T) {
// 	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapedia")
// 	datagedung := gq.GeoIntersects(mconn, 107.66417814690686, -6.901637720654366)
// 	fmt.Println(datagedung)
// }

func TestGeoWithin(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	coordinates := [][][]float64{
		{
			{103.62052506248301,
				-1.6105001000148462},
			{103.62061804929925,
				-1.6106710617710007},
			{103.62071435707355,
				-1.6106229269090022},
			{103.62061472834131,
				-1.6104420062116702},
			{103.62052506248301,
				-1.6105001000148462},
		},
	}

	datagedung := gq.GeoWithin(mconn, coordinates)
	fmt.Println(datagedung)
}

// func TestBox(t *testing.T) {
// 	mconn := helper.SetConnection("mongodb://localhost:27017", "petapedia")
// 	datagedung := gq.Box(mconn, 107.64914828236414, -6.912194782091035, 107.64914828236414, -6.912194782091035)
// 	fmt.Println(datagedung)
// }
