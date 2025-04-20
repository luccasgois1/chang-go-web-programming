package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

// This function initialized the Db variable when the application starts
func init() {
	var err error
	Db, err = sql.Open("postgres", "host=chitchat-db user=luccas password=senha dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}
