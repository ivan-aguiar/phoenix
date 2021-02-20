package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ivan-aguiar/phoenix/db"
)

//ViewProfile permite extraer los valores del perfil
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error has ocurred while searching registry"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)
}
