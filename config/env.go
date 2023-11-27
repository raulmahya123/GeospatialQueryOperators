package config

import "os"

var EndpointSecret = os.Getenv("SECRET")
var Mongostring = os.Getenv("MONGOSTRING")
