package geoquery

import (
	"context"
	"fmt"
	"testing"

	"github.com/raulmahya123/GeospatialQueryOperators/gq"
	"github.com/raulmahya123/GeospatialQueryOperators/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func TestPolygon(t *testing.T) {
	// Set up MongoDB connection for testing
	mconn := helper.SetConnection("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority", "petapediaaa")

	// Example coordinates for a polygon
	coordinates := [][][]float64{
		{
			{103.62052506248301, -1.6105001000148462},
			{103.62061804929925, -1.6106710617710007},
			{103.62071435707355, -1.6106229269090022},
			{103.62061472834131, -1.6104420062116702},
			{103.62052506248301, -1.6105001000148462},
		},
	}

	// Call the function being tested
	result := gq.Polygon(mconn, coordinates)

	// Add your assertions based on expected behavior
	expectedResult := ""
	if result != expectedResult {
		t.Errorf("Expected '%s', got '%s'", expectedResult, result)
	}
}

func TestGeoBoxQuery(t *testing.T) {
	// Connect to MongoDB (replace with your MongoDB connection details)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://raulgantengbanget:0nGCVlPPoCsXNhqG@cluster0.9ofhjs3.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Check if MongoDB is reachable
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		t.Fatalf("Error pinging MongoDB: %v", err)
	}

	// Define lowerLeft and upperRight coordinates for testing
	lowerLeft := []float64{103.6205, -1.6106}
	upperRight := []float64{103.6207, -1.6104}

	// Call the GeoBoxQuery function being tested
	results, err := gq.GeoBoxQuery(client.Database("box"), lowerLeft, upperRight)

	// Check for errors
	if err != nil {
		t.Fatalf("Error in GeoBoxQuery: %v", err)
	}

	// Print or use the results as needed for testing assertions
	fmt.Println("Results:", results)

	// Add your assertions here based on the expected results
	// For example, you can check if the length of the results is as expected
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
