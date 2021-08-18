package auth

import (
	"context"
	"log"
	"net/http"
	"simvino/models/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			email, err := ParseToken(header)

			if err != nil {
				log.Fatal(err)
			}

			user, err := users.GetUserByEmail(email)

			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
