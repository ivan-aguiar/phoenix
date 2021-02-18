package db

import (
	"github.com/ivan-aguiar/phoenix/models"
	"golang.org/x/crypto/bcrypt"
)

/*AttemptLogin chequea el login con la DB*/
func AttemptLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)                               //Password del login
	passwordDB := []byte(user.Password)                             //Password encriptada en la DB
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes) //Bcrypt compara ambas passwords
	if err != nil {
		return user, false
	}

	return user, true
}
