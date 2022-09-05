package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
)

func createURL() string {
	// URL dinámica
	u, err := url.Parse("/user")
	if err != nil {
		panic(err)
	}
	u.Host = "localhost:3001"
	u.Scheme = "http"

	query := u.Query() // Retorna un map
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()

	// URL estatica
	// u, err := url.Parse("http://localhost:3000/user?nombre:valor")
	// if err != nil {
	// 	panic(err)
	// }

	return u.String()

}

func main() {
	url := createURL()
	// fmt.Println("La URL final es: ", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Encabezado", "Valor") // Crear encabezados eal request
	

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println("El header es ", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("El body es ", string(body))
	fmt.Println("El status es ", response.Status)

	log.Println("Servidor en puerto 8000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// Trabajando con header

// http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(r.Header) // curl -i http://localhost:3000/user -H "ecabezado:valor"

// 	accessToken := r.Header.Get("access_token")
// 	if len(accessToken) != 0 {
// 		fmt.Println("accessToken")
// 	}
// 	r.Header.Set("nombre", "valor")
// 	fmt.Println(r.Header)
// })

// // Trabajar nuestros query

// fmt.Println(r.URL)

// values := r.URL.Query()
// values.Add("name", "eduardo")
// values.Add("Course", "Go Web")
// values.Add("Job", "Codigo facilito")
// values.Del("otro") // clear && curl -i http://localhost:3000/user?otro=otrossss
// // Anexando parametros a la url
// r.URL.RawQuery = values.Encode()

// fmt.Println(r.URL)
// fmt.Fprintf(w, "Hola Mundo")

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {      // Configurar la URI "/", W => Estructura respuesta. r => estructura de envio
// 	w.Header().Add("Nombre", "Valor del header")                         // Encabezados // curl -i http://localhost:3000 Se corre en el terminal para verificar los headers
// 	http.Redirect(w, r, "/dos", http.StatusOK)                           // Redireccionar.  // Hay variables ya definidas en https://golang.org/src/net/http/status.go
// })
// http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
// 	http.NotFound(w, r)
// })

// http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "Este es un error.", http.StatusConflict)
// })

// Curl Modificar el metodo de la función se usa la vandera
// -X mas el metodo "-X POST, -X PUT, -X Delete" => al final del request del curl
// curl -i http://localhost:3000/user?name=joseito
// Query con doble parametro
// curl -i "http://localhost:3000/user?name=jose&apellido=machuca"

// fmt.Println("El método es: ", r.Method)
// switch r.Method {
// case "GET":
// 	fmt.Fprintf(w, "Método GET")
// case "POST":
// 	fmt.Fprintf(w, "Método POST")
// case "PUT":
// 	fmt.Fprintf(w, "Método PUT")
// case "DELETE":
// 	fmt.Fprintf(w, "Método DELETE")
// default:
// 	http.Error(w, "Método no valido", http.StatusBadRequest)
// }

// Pasando parametros con ?
// localhost:3000/user?neme=jose

// Obtener argumentos del url

// // fmt.Println(r.URL) // Trae toda la url
// // fmt.Println(r.URL.RawQuery) // Obtenemos los parametros
// fmt.Println(r.URL.Query()) // Devuelve un map

// name := r.URL.Query().Get("name")
// if len(name) != 0{
// 	fmt.Println(name)
// }
// apellido := r.URL.Query().Get("apellido")
// if len(apellido) != 0{
// 	fmt.Println(apellido)
// }
