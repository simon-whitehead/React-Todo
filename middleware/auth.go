package middleware

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/simon-whitehead/react-todo/services"
	"github.com/zenazn/goji/web"

	"github.com/simon-whitehead/react-todo/domain"
)

const (
	CookieKey = "THE_SUPER_SECRET_COOKIE_KEY"
)

var (
	CookieStore *sessions.CookieStore
)

func init() {

	gob.Register(domain.User{})

	CookieStore = sessions.NewCookieStore([]byte(CookieKey))
	CookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: false,    // HTTPS
	}
}

func NewSession(c *web.C, r *http.Request) (*sessions.Session, error) {
	s, e := CookieStore.Get(r, "React-Todo")
	if e != nil {
		return nil, e
	}

	return s, nil
}

func Auth(c *web.C, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s, e := NewSession(c, r); e == nil && s.Values["user"] != nil {
			s.Values["user"] = func() *domain.User {
				if svc, ok := c.Env["UserService"].(services.UserServicer); ok {
					return svc.GetUserByEmail(s.Values["user"].(*domain.User).Email)
				}

				return nil
			}
			s.Save(r, w)
			c.Env["Session"] = s
			h.ServeHTTP(w, r)
		} else {
			http.RedirectHandler("/login", http.StatusSeeOther).ServeHTTP(w, r)
		}
	})
}
