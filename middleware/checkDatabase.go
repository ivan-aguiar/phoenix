package middleware

import (
	"net/http"

	"github.com/ivan-aguiar/phoenix/db"
)

//CHeckDatabase es el middleware que permite conocer el estado de la DB
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Database connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
