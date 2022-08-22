package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // Libreria interna de go
	"github.com/gorilla/mux" // go get -u github.com/gorilla/mux
	"codigo_facilito/server/src/connect"
)

type User struct {
	Username string `json:"username"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
}

func main() {

	connect.InitalizeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter() 							// Permite crear una intancia de NewRouter
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")  // Declaramos el usuario

	log.Println("Servido en puerto 8000")           // Indicamos que la apertura del servidor
	log.Fatal(http.ListenAndServe(":8000", r))      // Levantando el servido
}

func GetUser(w http.ResponseWriter, r* http.Request){  // Asociamos func a url
	vars := mux.Vars(r)
	user_id := vars["id"]
	fmt.Println(user_id)
	// Se asginan valores sin el atributo del struct
	user := User{"Admin", "José", "Machuca",}
	json.NewEncoder(w).Encode(user)
}