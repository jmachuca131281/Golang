package main

import (
	"encoding/json"
	"log"
	"net/http" // Libreria interna de go
	"github.com/gorilla/mux" // go get -u github.com/gorilla/mux
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var connection *gorm.DB
const engine_sql string = "mysql"

// Utilizar variables de entorno.
const username string = "root"
const password string = "R00t"
const database string = "taller1"

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
}

type Response struct {
	Status string `json:"status"`
	Data User `json:"data"`
	Message string `json:"message"`
}

func Getuser(w http.ResponseWriter, r* http.Request){  // Asociamos func a url
	vars := mux.Vars(r)
	user_id := vars["id"]
	status := "success"
	var message string
	user := GetUser(user_id)

	if user.Id <= 0 {
		status = "error"
		message = "User not found"
	}

	response := Response{status, user, message}
	json.NewEncoder(w).Encode(response)
}

func InitalizeDatabase() {
	connection = ConnectORM(CreateString())
	log.Println("Conexión DB exitosa")
}

func CloseConnection() {
	connection.Close()
	log.Println("Conexión db cerrada")
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func GetUser(id string) User {
	user := User{}
	connection.Where("id=?", id).First(&user)
	return user
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}

func NewUser(w http.ResponseWriter, r* http.Request) {
	user := GetUserRequest(r)
	response := Response{"success", CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r* http.Request) User {
	var user = User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func CreateUser(user User)  User{
	connection.Create(&user)
	return user
}

func Updateuser(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]	
	user := GetUserRequest(r)
	response := Response{"success", UpdateUser(user_id, user), ""}
	UpdateUser(user_id, user)
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(id string, user User) User {
	currentUser := User{}
	connection.Where("id=?", id).First(&currentUser)
	currentUser.Username = user.Username
	currentUser.First_name = user.First_name 
	currentUser.Last_name = user.Last_name
	connection.Save(&currentUser)
	return currentUser
}

func Deleteuser(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	var user User
	DeleteUser(user_id)
	response := Response{"succes", user, ""}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(id string) {
	user := User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}

func main() {
	InitalizeDatabase()
	defer CloseConnection()
	r := mux.NewRouter() 							                // Permite crear una intancia de NewRouter

	r.HandleFunc("/user/{id}", Getuser).Methods("GET")              // Declaramos el usuario
	r.HandleFunc("/user/new", NewUser).Methods("POST")              // Declaramos el usuario
	r.HandleFunc("/user/update/{id}", Updateuser).Methods("PATCH")  // Declaramos el usuario
	r.HandleFunc("/user/delete/{id}", Deleteuser).Methods("DELETE")  // Declaramos el usuario

	log.Println("Servido en puerto 8000")                           // Indicamos que la apertura del servidor
	log.Fatal(http.ListenAndServe(":8000", r))                      // Levantando el servido
}

