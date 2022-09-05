package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

// Handlers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	user := User{Id: 1, Username: "eduardo_gpg", Password: "password123"}
	vars := mux.Vars(r)
	userId, _ = strconv.Atoi(vars["id"])
	user, err := GetUserModel(user.Id)
	output, _ := json.Marshal(&user)
	fmt.Fprintf(w, string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza el usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina el usuario")
}

// Models
type User struct {
	Id       int    `json:"id"`       // Por las reglas de Go hay que colocar los sobrenombres para que
	Username string `json:"username"` // Recordar que en mayusculas es publico
	Password string `json:"password"`
}

// Mapas para simular una base de datos.

var users = make(map[id]User)

func SetDefaultUser() {
	user := User{Id: 1, Username: "eduardo_gpg", Password: "password123"}
	users[user.Id] = user
}

func GetUserModel(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return user{}, errors.New("El usuario no existe")
}

func main() {
	mux := mux.NewRouter()
	SetDefaultUser()
	mux.HandleFunc("/api/v1/users/", GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", DeleteUser).Methods("DELETE")

	log.Println("\nServidor puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}

//$ clear && curl http://localhost:8000/api/v1/users/1 -X GET -i
