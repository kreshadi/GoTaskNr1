package main

import (
	"ExampleProject/controllers"
	"ExampleProject/driver"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")

	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listen on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {

}

//we axpect a handlerFunction as a returned type
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {

	return nil
}
