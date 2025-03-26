package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	//"strings"
)

func StartLibraryCLI() {
	library := services.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			library.AddBook(models.Book{ID: bookID, Title: title, Author: author, Status: "Available"})
			fmt.Println("Book added successfully.")

		case "2":
			fmt.Print("Enter Book ID to remove: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			library.RemoveBook(bookID)
			fmt.Println("Book removed successfully.")

		case "3":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Book ID to borrow: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case "4":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter Book ID to return: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case "5":
			books := library.ListAvailableBooks()
			fmt.Println("\nAvailable Books:")
			for _, book := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
			}

		case "6":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			books := library.ListBorrowedBooks(memberID)
			if len(books) == 0 {
				fmt.Println("No borrowed books found for this member.")
			} else {
				fmt.Println("\nBorrowed Books:")
				for _, book := range books {
					fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
				}
			}

		case "7":
			fmt.Println("Exiting Library Management System...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
