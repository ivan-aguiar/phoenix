package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ivan-aguiar/phoenix/db"
	"github.com/ivan-aguiar/phoenix/models"
)

//Email es el valor del email usado en todos los endpoints
var Email string

//UserID es el ID devuelto del modelo que se usará en todos los endpoints
var UserID string

//ProcessToken procesa el token para extraer sus valores
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("fulcrum")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	//Corrobora que el Token tenga dos palabras
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		_, found, _ := db.CheckUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
