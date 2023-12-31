package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"learningGorillamux/data"
	"learningGorillamux/datatypes"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdstr := vars["bookId"]
	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}
	RequiredBook, bookExists := data.Books[bookId]

	if !bookExists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("couldn't find the book"))
		return
	}

	marsheledBook, err := json.Marshal(RequiredBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marsheledBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	listOfBooks := data.Books
	var dosk datatypes.Books

	for Bid, Bname := range listOfBooks {
		book := datatypes.Book{
			Id:   Bid,
			Name: Bname,
		}

		dosk.AddBooksToList(book)
	}

	marshaledData, err := json.Marshal(dosk)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while marsheling data"))
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledData)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdstr := vars["bookId"]
	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	_, exists := data.Books[bookId]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book doesn't exists"))
		return
	}

	delete(data.Books, bookId)
	w.Write([]byte("Deleted book sucessfully!"))
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook datatypes.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Println(io.ReadAll(r.Body))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while decoding"))
		return
	}

	_, exists := data.Books[newBook.Id]

	if exists {
		w.Write([]byte("book already exists"))
		return
	}

	data.Books[len(data.Books)+1] = newBook.Name
	w.Write([]byte("Book added Sucessfully"))

}
