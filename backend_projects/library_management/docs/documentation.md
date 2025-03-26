# Library Management System

## Overview
This is a simple console-based library management system in Go.

## Features
- Add books to the library
- Remove books from the library
- Borrow books
- Return books
- List available books
- List borrowed books by a member

## Folder Structure
library_management/ ├── main.go ├── controllers/ │ └── library_controller.go ├── models/ │ └── book.go │ └── member.go ├── services/ │ └── library_service.go │ └── library_service_test.go ├── docs/ │ └── documentation.md └── go.mod

## How to Run
```sh
go run main.go

