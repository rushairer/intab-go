package middlewares

import (
	"encoding/json"
	oauth2 "intab-core/oauth2"
	pp "intab-core/passport"
	repo "intab-core/repositories"
	"net/http"
	"strconv"

	"github.com/rushairer/ago"
)

//Auth Auth中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenInfo, err := oauth2.Server().ValidationBearerToken(r)

		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(ago.Result401); err != nil {
				panic(err)
			}
		} else {
			//userLogin
			userID, _ := strconv.Atoi(tokenInfo.GetUserID())
			passport := pp.Passport{}
			passport.AccountRepository = repo.NewAccountRepositoryWithID(userID)
			passport.Store(r, w)
			next.ServeHTTP(w, r)
		}
	})
}
