package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	service "api/service"
)

func main() {
	port := 8080

	router := mux.NewRouter()

	router.HandleFunc("/cakes", service.ListOfCake).Methods("GET")
	router.HandleFunc("/cakes/{id}", service.DetailOfCake).Methods("GET")
	router.HandleFunc("/cakes", service.AddNewCake).Methods("POST")
	router.HandleFunc("/cakes/{id}", service.UpdateCake).Methods("PATCH")
	router.HandleFunc("/cakes/{id}", service.DeleteCake).Methods("DELETE")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Print(err)
		panic(err)
	}

	log.Print("Connected to Port:", port)
}
