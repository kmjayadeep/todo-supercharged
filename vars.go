package main

import (
	"fmt"
	"os"
)

var MONGO_URL = os.Getenv("MONGO_URL")
var MONGO_DB = os.Getenv("MONGO_DB")
var MONGO_USER = os.Getenv("MONGO_USER")
var MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")

var CONNECTION_STRING = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?authSource=admin&tls=true&tlsCAFile=ca-certificate.crt", MONGO_USER, MONGO_PASSWORD, MONGO_URL, MONGO_DB)
