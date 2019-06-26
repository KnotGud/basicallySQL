package main

import (
	"fmt"
	"log"
	"matching/app"
)

func main() {
	doThings()
}

func doThings() {
	db, err := app.NewDatabaseFromJSON("./example.json")
	if err != nil {
		log.Panic(err)
	}
	m := app.NewManager(db)

	for _, v := range m.GetResults(int64(98607), "test") {
		fmt.Println(v)
	}
}
