package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	// "html/template"
)

type Usuario struct {
	UserName string
	Edad int
}

func main(){
	fmt.Println("Hola mundo")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// template, err := template.New("Hola").Parse("Hola mundo") // Estring sera el nombre del template
		template, err := template.ParseFiles("index.html") // Visualizar el html en web
		if err != nil {
			panic(err)	
		}
		
		usuario := Usuario{"Josman", 40}
		// Para que el mesage sea visto por el cliente.
		template.Execute(w, usuario)

	})
	fmt.Println("Servidor escuchando puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}


/**
Archivos que nos permiten trabajar páginas web dinámicas.
html/template => Ayuda a preveer inyección de código para prevenir ataques.



*/ 
