package db

import "golang.org/x/crypto/bcrypt"

//EncryptPassword es la función que encripta la contraseña del usuario
func EncryptPassword(password string) (string, error){
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}