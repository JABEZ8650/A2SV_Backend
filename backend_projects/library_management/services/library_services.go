package services

import (
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *library {
	return &library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// methods
func (l *library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book not found")
	}
	if book.Status == "Borrowed" {
		return fmt.Errorf("book already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member not found")
	}

	bookIndex := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		return fmt.Errorf("book not found in member’s borrowed list")
	}

	book := member.BorrowedBooks[bookIndex]
	book.Status = "Available"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	l.Members[memberID] = member
	return nil
}

func (l *library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *library) ListBorrowedBooks(memberID int) []models.Book {
	if member, exists := l.Members[memberID]; exists {
		return member.BorrowedBooks
	}
	return nil
}
