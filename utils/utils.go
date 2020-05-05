package utils

import (
	"ExampleProject/models"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

//reponses JSON. It accepsts responsewriter and the data obj, and we use jsonencoder and pass data to encode
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
//return String that is a Token that we are going to generate, and an error if theer is an error
func GenerateToken(user models.User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	spew.Dump(token)


	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}