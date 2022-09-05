package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	 fmt.Fprintf(w, "HandleRoot")
}

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Handle home")
}
