package webserver

import (
	"context"
	"net/http"

	"github.com/GolangWorkshop/library/util"
	"github.com/golang-jwt/jwt"
)

func HTTPLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("%s %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func UserContextMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if c == nil || err != nil {
			log.Info().Msgf("could not find token cookie, continue as anonymous")
			h.ServeHTTP(w, r)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &util.JwtClaims{}

		// Parse the JWT string and store the result in `claims`.
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return util.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Info().Msgf("could not authorize user due to invalid signature, continue as anonymous")
				h.ServeHTTP(w, r)
				return
			}
			log.Info().Msgf("could not authorize user, continue as anonymous")
			h.ServeHTTP(w, r)
			return
		}
		if !tkn.Valid {
			log.Info().Msgf("could not authorize user due to invalid token, continue as anonymous")
			h.ServeHTTP(w, r)
			return
		}

		log.Info().Msgf("request fired by user %s", claims.Username)
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
