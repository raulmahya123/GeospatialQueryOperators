package config

import "github.com/petapedia/geoquery/helper"

var Mongocon = helper.SetConnection(Mongostring, "geojson")
