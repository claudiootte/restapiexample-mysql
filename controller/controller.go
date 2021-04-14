package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudiootte/restapiexample/model"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	MyRouter := mux.NewRouter()
	MyRouter.HandleFunc("/books", model.GetAllBooks).Methods("GET")
	MyRouter.HandleFunc("/books/{id}", model.GetBook).Methods("GET")
	MyRouter.HandleFunc("/books", model.CreateBook).Methods("POST")
	MyRouter.HandleFunc("/books/{id}", model.UpdateBook).Methods("PUT")
	MyRouter.HandleFunc("/books/{id}", model.DeleteBook).Methods("DELETE")

	fmt.Println("Connecting to server...")
	log.Fatal(http.ListenAndServe(":8080", MyRouter))
}
