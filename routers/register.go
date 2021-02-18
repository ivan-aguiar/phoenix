package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ivan-aguiar/phoenix/db"
	"github.com/ivan-aguiar/phoenix/models"
)

//Register añade a la DB el usuario
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Your password must be at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserExists(user.Email)
	if found == true {
		http.Error(w, "Email already registered", 400)
	}

	_, status, err := db.InsertRegister(user)
	if err != nil {
		http.Error(w, "An error occurred while registering", 400)
	}

	if status == false {
		http.Error(w, "Error while registering", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
