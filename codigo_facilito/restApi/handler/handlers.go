package handler

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se obtiene todlos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se obtiene un usruario!")
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se actualiza el usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se elimina el usuario")
}

