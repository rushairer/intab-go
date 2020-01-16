package middlewares

import (
	pp "intab-core/passport"
	"net/http"
)

//Passport Passport中间件
func Passport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		passport := pp.InitPassport()
		if passport.Get(r, w) {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})
}
