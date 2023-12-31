package data

var Users = map[string]string{ //[username] password
	"yash": "pass",
	"tush": "password",
}

var Books = map[int]string{ // [BookId] BookName
	1: "Game of Thrones",
	2: "Animal farm",
	3: "Crime And Punishment",
}

var Orders = make(map[string][]int)

var NubmerOfRequests int
