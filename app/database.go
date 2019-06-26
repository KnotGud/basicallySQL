package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Database is an instance of an object that works off
type Database struct {
	Providers []Provider `json:"providers"`
	Claims    []Claim    `json:"claims"`
}

// NewDatabaseFromJSON builds a database from json files
func NewDatabaseFromJSON(jsonPath string) (*Database, error) {
	var db Database

	f, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
