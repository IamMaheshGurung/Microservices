package data

type Books struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var bookList = []*Books{
	{
		ID:          1,
		Name:        "Computer Science",
		Description: "A programmers Perspective",
	},
}

func GetBook() []*Books {
	return bookList
}

func AddBook(book *Books) {
	bookList = append(bookList, book)
}

func SetBooks(books []*Books) {
	bookList = books
}
