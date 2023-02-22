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
   router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
   router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

   fmt.Println("Escutando na porta 3000")
   log.Fatal(http.ListenAndServe(":3000", router))
}