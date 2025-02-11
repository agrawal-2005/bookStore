package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/agrawal-2005/bookStore/pkg/utils"
	"github.com/agrawal-2005/bookStore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(uint(bookID))
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)

	newBook, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newBook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	existingBook, err := models.GetBookById(uint(bookID))
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updatedData models.Book
	utils.ParseBody(r, &updatedData)

	if updatedData.Name != "" {
		existingBook.Name = updatedData.Name
	}
	if updatedData.Author != "" {
		existingBook.Author = updatedData.Author
	}
	if updatedData.Publication != "" {
		existingBook.Publication = updatedData.Publication
	}

	updatedBook, err := existingBook.UpdateBook()
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteBook(uint(bookID)); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Book deleted successfully"}
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
