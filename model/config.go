package model

type Config struct {
	File string
}

var App = &Config{
	File: "storage/output.csv",
}
