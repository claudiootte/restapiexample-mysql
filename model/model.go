package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book []Book
	DB.Find(&book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(book)
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		err02 := DB.Create(&book).Error
		errors.Is(err, gorm.ErrRecordNotFound)

		if err02 != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(http.StatusBadRequest)

		} else {
			json.NewEncoder(w).Encode(book)

		}
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	} else {
		json.NewDecoder(r.Body).Decode(&book)
		DB.Save(&book)
		json.NewEncoder(w).Encode(book)

	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(http.StatusNotFound)

	} else {
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			DB.Delete(&book)
			json.NewEncoder(w).Encode("the book is deleted successfully")

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)

		}
	}

}
