package middleware

import (
	"errors"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/zenazn/goji/web"
)

// db.go - Simple Goji middleware to pass
// a boltdb instance into the context

var (
	db *bolt.DB
)

func SetDatabase(database *bolt.DB) {
	db = database
}

func WithDB(c *web.C, next http.Handler) http.Handler {
	if db == nil {
		panic(errors.New("Must set DB Middleware DB value"))
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Env["Database"] = db

		next.ServeHTTP(w, r)
	})
}
