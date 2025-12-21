package library

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

func NewBook(id, title, author string) Book {
	return Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
}

func (b *Book) Borrow() bool {
	if b.IsBorrowed {
		return false
	}
	b.IsBorrowed = true
	return true
}

func (b *Book) Return() bool {
	if !b.IsBorrowed {
		return false
	}
	b.IsBorrowed = false
	return true
}

func (b Book) IsAvailable() bool {
	return !b.IsBorrowed
}
