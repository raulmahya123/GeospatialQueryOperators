package geoquery

import (
	"fmt"
	"testing"

	"github.com/raulmahya123/GeospatialQueryOperators/gq"
	"github.com/raulmahya123/GeospatialQueryOperators/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapedia")
	datagedung := gq.GeoIntersects(mconn, 107.66417814690686, -6.901637720654366)
	fmt.Println(datagedung)
}

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

func TestGeomerty(t *testing.T) {
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

	datagedung := gq.Geometry(mconn, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := helper.SetConnection2dsphereTest("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	datagedung := gq.Near(mconn, 103.61028753823456, -1.6246312350403684)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := helper.SetConnection2dsphereTestPoint("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
	datagedung := gq.NearSpehere(mconn, 40.78, -73.9667)
	fmt.Println(datagedung)
}

// func TestPolygon(t *testing.T) {
// 	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")
// 	coordinates := [][][]float64{
// 		{
// 			{103.62052506248301,
// 				-1.6105001000148462},
// 			{103.62061804929925,
// 				-1.6106710617710007},
// 			{103.62071435707355,
// 				-1.6106229269090022},
// 			{103.62061472834131,
// 				-1.6104420062116702},
// 			{103.62052506248301,
// 				-1.6105001000148462},
// 		},
// 	}

// 	datagedung := gq.Polygon(mconn, coordinates)
// 	fmt.Println(datagedung)
// }
