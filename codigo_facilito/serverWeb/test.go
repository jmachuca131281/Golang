package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo")
	})
	log.Println("Servidor en puerto 3001")
	log.Fatal(http.ListenAndServe("localhost:3001", nil))

	
}