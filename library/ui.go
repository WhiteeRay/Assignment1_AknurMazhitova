package Library

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LibraryUI struct {
	library *Library
	scanner *bufio.Scanner
}

func NewLibraryUI(library *Library) *LibraryUI {
	return &LibraryUI{
		library: library,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (ui *LibraryUI) ShowMenu() {
	for {
		ui.displayMenu()

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		ui.scanner.Scan()

		if !ui.handleMenuChoice(choice) {
			break
		}
	}
}

func (ui *LibraryUI) displayMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("=== %s ===\n", ui.library.GetLibraryName())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total Books: %d | Borrowed: %d | Available: %d\n",
		ui.library.GetTotalBooks(),
		ui.library.GetBorrowedCount(),
		ui.library.GetTotalBooks()-ui.library.GetBorrowedCount())
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("1. Add Book")
	fmt.Println("2. Borrow Book")
	fmt.Println("3. Return Book")
	fmt.Println("4. List Available Books")
	fmt.Println("5. List All Books")
	fmt.Println("6. Search Book by ID")
	fmt.Println("7. Exit")
}

func (ui *LibraryUI) handleMenuChoice(choice int) bool {
	switch choice {
	case 1:
		ui.handleAddBook()
	case 2:
		ui.handleBorrowBook()
	case 3:
		ui.handleReturnBook()
	case 4:
		ui.handleListAvailableBooks()
	case 5:
		ui.handleListAllBooks()
	case 6:
		ui.handleSearchBook()
	case 7:
		fmt.Println("Exiting Library System. Goodbye!")
		return false
	default:
		fmt.Println("Invalid option. Please try again.")
	}
	return true
}

func (ui *LibraryUI) handleAddBook() {
	fmt.Println("\n--- Add New Book ---")

	fmt.Print("Enter Book ID: ")
	id := ui.readLine()

	fmt.Print("Enter Book Title: ")
	title := ui.readLine()

	fmt.Print("Enter Book Author: ")
	author := ui.readLine()

	if err := ui.library.AddBook(id, title, author); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Book '%s' added successfully!\n", title)
	}
}

func (ui *LibraryUI) handleBorrowBook() {
	fmt.Println("\n--- Borrow Book ---")

	fmt.Print("Enter Book ID: ")
	id := ui.readLine()

	if err := ui.library.BorrowBook(id); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		book, _ := ui.library.GetBookByID(id)
		fmt.Printf("Book '%s' borrowed successfully!\n", book.Title)
	}
}

func (ui *LibraryUI) handleReturnBook() {
	fmt.Println("\n--- Return Book ---")

	fmt.Print("Enter Book ID: ")
	id := ui.readLine()

	if err := ui.library.ReturnBook(id); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		book, _ := ui.library.GetBookByID(id)
		fmt.Printf("Book '%s' returned successfully!\n", book.Title)
	}
}

func (ui *LibraryUI) handleListAvailableBooks() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("=== Available Books ===")
	fmt.Println(strings.Repeat("=", 50))

	books := ui.library.GetAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No available books at the moment.")
		return
	}

	for i, book := range books {
		fmt.Printf("%d. [ID: %s] %s by %s\n",
			i+1, book.ID, book.Title, book.Author)
	}
}

func (ui *LibraryUI) handleListAllBooks() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("=== All Books ===")
	fmt.Println(strings.Repeat("=", 50))

	books := ui.library.GetAllBooks()

	if len(books) == 0 {
		fmt.Println("No books in the library.")
		return
	}

	for i, book := range books {
		status := "Available"
		if book.IsBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("%d. [ID: %s] %s by %s - [%s]\n",
			i+1, book.ID, book.Title, book.Author, status)
	}
}

func (ui *LibraryUI) handleSearchBook() {
	fmt.Println("\n--- Search Book ---")

	fmt.Print("Enter Book ID: ")
	id := ui.readLine()

	book, err := ui.library.GetBookByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	status := "Available"
	if book.IsBorrowed {
		status = "Borrowed"
	}

	fmt.Println("\n--- Book Details ---")
	fmt.Printf("ID: %s\n", book.ID)
	fmt.Printf("Title: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Status: %s\n", status)
}

func (ui *LibraryUI) readLine() string {
	ui.scanner.Scan()
	return strings.TrimSpace(ui.scanner.Text())
}
