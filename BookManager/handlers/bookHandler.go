package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"xample.com/m/data"
)

type BookHandler struct {
	L *log.Logger
}

func NewBook(l *log.Logger) *BookHandler {
	return &BookHandler{l}
}

func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	gb := data.GetBook()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(gb)
	if err != nil {
		http.Error(w, "Unable to encode the data", http.StatusBadRequest)

	}

}
func (b *BookHandler) GetBookId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, book := range data.GetBook() {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
		http.Error(w, "BOok Not found in", http.StatusNotFound)
	}
}

func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book data.Books
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(data.GetBook()) + 1
	data.AddBook(&book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}

func (b *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedBook data.Books
	for i, book := range data.GetBook() {
		if book.ID == id {
			updatedBook.ID = id
			data.GetBook()[i] = &updatedBook

			json.NewEncoder(w).Encode(updatedBook)
			return

		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
func (b *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	books := data.GetBook()
	for i, book := range books {
		if book.ID == id {
			books = append(books[:1], books[i+1:]...)
			data.SetBooks(books)
			w.WriteHeader(http.StatusNoContent)
			return

		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
