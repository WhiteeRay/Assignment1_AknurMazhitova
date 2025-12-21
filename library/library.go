package library

import "fmt"

type Library struct {
	books map[string]Book
	name  string
}

func NewLibrary(name string) *Library {
	return &Library{
		books: make(map[string]Book),
		name:  name,
	}
}

func (l *Library) AddBook(id, title, author string) error {
	if _, exists := l.books[id]; exists {
		return fmt.Errorf("book with ID %s already exists", id)
	}

	book := NewBook(id, title, author)
	l.books[id] = book
	return nil
}

func (l *Library) BorrowBook(id string) error {
	book, exists := l.books[id]
	if !exists {
		return fmt.Errorf("book with ID %s not found", id)
	}

	if !book.Borrow() {
		return fmt.Errorf("book '%s' is already borrowed", book.Title)
	}

	l.books[id] = book
	return nil
}

func (l *Library) ReturnBook(id string) error {
	book, exists := l.books[id]
	if !exists {
		return fmt.Errorf("book with ID %s not found", id)
	}

	if !book.Return() {
		return fmt.Errorf("book '%s' was not borrowed", book.Title)
	}

	l.books[id] = book
	return nil
}

func (l *Library) GetAvailableBooks() []Book {
	available := make([]Book, 0)

	for _, book := range l.books {
		if book.IsAvailable() {
			available = append(available, book)
		}
	}

	return available
}

func (l *Library) GetAllBooks() []Book {
	books := make([]Book, 0, len(l.books))

	for _, book := range l.books {
		books = append(books, book)
	}

	return books
}

func (l *Library) GetBookByID(id string) (Book, error) {
	book, exists := l.books[id]
	if !exists {
		return Book{}, fmt.Errorf("book with ID %s not found", id)
	}
	return book, nil
}

func (l *Library) GetLibraryName() string {
	return l.name
}

func (l *Library) GetTotalBooks() int {
	return len(l.books)
}

func (l *Library) GetBorrowedCount() int {
	count := 0
	for _, book := range l.books {
		if book.IsBorrowed {
			count++
		}
	}
	return count
}
