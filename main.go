package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "postgres://wizardofmath:testing@172.17.0.2:5432/food?sslmode=disable&TimeZone=UTC")
	if err != nil {
		panic(err)
	}

	id := 1
	data := Recipe{}
	err = db.Get(&data, "SELECT * FROM recipes WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
