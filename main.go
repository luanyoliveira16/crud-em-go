package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

   router := mux.NewRouter()
   router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

   fmt.Println("Escutando na porta 3000")
   log.Fatal(http.ListenAndServe(":3000", router))
}