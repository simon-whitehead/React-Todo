package middleware

import (
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/zenazn/goji/web"

	"github.com/simon-whitehead/react-todo/services"
)

func UserService(c *web.C, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Env["UserService"] = services.NewUserService(c.Env["Database"].(*bolt.DB))

		next.ServeHTTP(w, r)
	})
}
