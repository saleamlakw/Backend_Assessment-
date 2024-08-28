package passwordutil

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string,error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func VerifyPassword(userPassword string, providedPassword string) (bool, error) {
	check := true
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		check = false
	}
	return check, err
}