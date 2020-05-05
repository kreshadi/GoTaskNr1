package controllers

import (
	"ExampleProject/models"
	"ExampleProject/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Controller struct {
	
}

func (c Controller)Signup(db *sql.DB) http.HandlerFunc	{


	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)

		//spew.Dump(user)

		if user.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)

		stmt := "insert into users (email, password) values($1, $2) RETURNING id;"

		err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
		fmt.Println(err)
		if err != nil {
			error.Message = "Server error."
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		user.Password = ""

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, user)

	}
}

func (c Controller) Login(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Called login"))

		var user models.User

		json.NewDecoder(r.Body).Decode(&user)

		token, err := utils.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)

	}
}



