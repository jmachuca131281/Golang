package connect

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorn/dialects/mysql"	
	"log"
)

var connection *gorm.DB

const engine_sql string = "mysql"

// Utilizar variables de entorno.
const username string = "jm"
const password string = "R00t"
const database string = "taller1"

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

func CreateString() string {
	return username + ":" + password + "@/" + database
}