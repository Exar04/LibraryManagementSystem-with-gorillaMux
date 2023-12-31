package controllers

import (
	"encoding/json"
	"learningGorillamux/data"
	"learningGorillamux/datatypes"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func OrderBook(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	vars := mux.Vars(r)
	bookIdstr, errbool := vars["bookId"]
	if !errbool {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book id not provided in url"))
		return
	}

	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book id not provided in url"))
		return
	}

	_, exists := data.Books[bookId]

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Books is not in the store"))
		return
	}

	data.Orders[username] = append(data.Orders[username], bookId)
	w.Write([]byte("Book ordered Successfully"))
}

func ListAllOrderedBooks(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	var listOfBooks datatypes.Books

	orders := data.Orders[username]

	for _, bookId := range orders {
		book := datatypes.Book{
			Id:   bookId,
			Name: data.Books[bookId],
		}
		listOfBooks.AddBooksToList(book)
	}
	marsheledData, err := json.Marshal(listOfBooks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatoin/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marsheledData)
}
