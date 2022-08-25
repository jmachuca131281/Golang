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
	Id         int    `json:"id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type Response struct {
	Status  string `json:"status"`
	Data    User   `json:"data"`
	Message string `json:"message"`
}

func GetUserRegisterRequest(w http.ResponseWriter, r* http.Request){  // Asociamos func a url
	vars := mux.Vars(r)                                // Vars returns the route variables for the current request, if any. Return Map
	user_id := vars["id"]                              // Separamos el id

	status := "success"
	message := "Return user"

	user := GetUserQuery(user_id)

	if user.Id <= 0 {
		status = "ERROR INVALID_ID"
		message = "User not found"
	}

	response := Response{status, user, message}
	json.NewEncoder(w).Encode(response)   				 // NewEncoder returns a new encoder that writes to w.
}

func InitalizeDatabase() {
	connection = ConnectORM(CreateString())
	log.Println("Conexión DB exitosa")
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func CloseConnection() {
	connection.Close()
	log.Println("Conexión DB cerrada")
}


func GetUserQuery(id string) User {
	user := User{}
	connection.Where("id = ?", id).First(&user)                // Return a new relation
	return user
}

func NewUserRequest(w http.ResponseWriter, r* http.Request) {
	user := GetUserRequest(r)
	response := Response{"success", CreateUser(user), "User is created"}
	json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r* http.Request) User {
	user := User{}
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

func UpdateUserRequest(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]	
	user := GetUserRequest(r)
	response := Response{"success", UpdateUser(user_id, user), "User update correctly"}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(id string, user User) User {
	currentUser := User{}
	connection.Where("id = ?", id).First(&currentUser)                  // Query to DB

	currentUser.Username = user.Username
	currentUser.First_name = user.First_name 
	currentUser.Last_name = user.Last_name

	connection.Save(&currentUser)
	
	return currentUser
}

func DeleteUserRequest(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	var user User
	DeleteUser(user_id)
	response := Response{"succes", user, "Register eliminated"}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(id string) {
	user := User{}
	connection.Where("id = ?", id).First(&user)       // Query to db
	connection.Delete(&user)
}

func main() {
	InitalizeDatabase()
	defer CloseConnection()

	// Register URL paths and handlers
	r := mux.NewRouter() 							                // Permite crear una intancia de NewRouter

	r.HandleFunc("/user/{id}", 		  GetUserRegisterRequest).Methods("GET")      // Declaramos el usuario
	r.HandleFunc("/user/new", 		  NewUserRequest).Methods("POST")              // Declaramos el usuario
	r.HandleFunc("/user/update/{id}", UpdateUserRequest).Methods("PATCH")  // Declaramos el usuario
	r.HandleFunc("/user/delete/{id}", DeleteUserRequest).Methods("DELETE")  // Declaramos el usuario

	log.Println("Servidor en puerto 8000")                           // Indicamos que la apertura del servidor
	log.Fatal(http.ListenAndServe(":8000", r))                      // Levantando el servido
	
}