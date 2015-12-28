package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/zenazn/goji"

	"github.com/simon-whitehead/react-todo/middleware"
)

func main() {
	db := getBoltDb()
	middleware.SetDatabase(db)

	defer db.Close()

	// flag.Set("bind", ":other_port_here") - the default is 8000
	goji.Serve()
}

func getBoltDb() *bolt.DB {
	db, err := bolt.Open("todo.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
