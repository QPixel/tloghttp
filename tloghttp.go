package tloghttp

import (
	"github.com/ubergeek77/tinylog"
	"net/http"
)

func NewHandler(lgr *tinylog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(CreateContext(r.Context()))
			next.ServeHTTP(w, r)
		})
	}
}
