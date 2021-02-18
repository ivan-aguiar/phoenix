package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ivan-aguiar/phoenix/models"
)

func GenerateJWT(user models.User) (string, error) {

	//Clave privada del JWT
	myKey := []byte("fulcrum")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.Lastname,
		"birthday":  user.Birthday,
		"bio":       user.Bio,
		"ubication": user.Ubication,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
