package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ivan-aguiar/phoenix/db"
	"github.com/ivan-aguiar/phoenix/jwt"
	"github.com/ivan-aguiar/phoenix/models"
)

//Login realiza el login del User
func Login(w http.ResponseWriter, r *http.Request) {
	//Seteamos al header que el contenido que devuelve (w) es tipo JSON
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid User or Password "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email is required ", 400)
	}

	document, exists := db.AttemptLogin(user.Email, user.Password)
	if exists == false {
		http.Error(w, "Invalid User or Password ", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error has ocurred while generating token "+err.Error(), 400)
		return
	}
	//Si el Token se encuentra generado
	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//Grabación de Cookie
	//Generación de un campo fecha para ver la expiración de la Cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
