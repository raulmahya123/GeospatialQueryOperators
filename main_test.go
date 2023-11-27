package geoquery

import (
	"fmt"
	"testing"

	"github.com/petapedia/geoquery/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("MONGOSTRING", "geojson")
	datagedung := helper.GetAllBangunanLineString(mconn, "petapedia")
	fmt.Println(datagedung)
}
