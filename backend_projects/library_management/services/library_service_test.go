package services

import (
	"library_management/models"
	"testing"
)

func TestLibraryFunctions(t *testing.T) {
	library := NewLibrary()

	// Test AddBook
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}
	library.AddBook(book)

	if len(library.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(library.Books))
	}

	// Test RemoveBook
	library.RemoveBook(1)
	if len(library.Books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(library.Books))
	}

	// Test BorrowBook
	library.AddBook(book)
	member := models.Member{ID: 101, Name: "Alice"}
	library.Members[101] = member

	err := library.BorrowBook(1, 101)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if library.Books[1].Status != "Borrowed" {
		t.Errorf("Expected book status 'Borrowed', got %s", library.Books[1].Status)
	}

	if len(library.Members[101].BorrowedBooks) != 1 {
		t.Errorf("Expected 1 borrowed book, got %d", len(library.Members[101].BorrowedBooks))
	}

	// Test ReturnBook
	err = library.ReturnBook(1, 101)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if library.Books[1].Status != "Available" {
		t.Errorf("Expected book status 'Available', got %s", library.Books[1].Status)
	}

	if len(library.Members[101].BorrowedBooks) != 0 {
		t.Errorf("Expected 0 borrowed books, got %d", len(library.Members[101].BorrowedBooks))
	}
}
