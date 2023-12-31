package datatypes

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Books struct {
	ListOfBooks []Book
}

func (B *Books) AddBooksToList(book Book) {
	B.ListOfBooks = append(B.ListOfBooks, book)
}
