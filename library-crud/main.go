package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	Server struct {
		p string
	}
	Book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	BookInfo struct {
		Book     Book
		Borrowed bool
	}
	Library struct {
		Books map[string]*BookInfo
	}
	Response struct {
		Result string `json:"Result"`
		Error  string `json:"Error"`
	}
)

var library = Library{
	Books: make(map[string]*BookInfo),
}

func NewServer(port string) *Server {
	return &Server{p: port}
}

func (s *Server) Start() {
	http.HandleFunc("/book", book)
	err := http.ListenAndServe(":"+s.p, nil)
	if err != nil {
		panic("server didn't start")
	}
}

func book(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Create a new Book object
		var newBook Book

		// Check the Content-Type header to decide how to parse the request
		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}
			newBook.Title = r.FormValue("title")
			newBook.Author = r.FormValue("author")
		} else if strings.HasPrefix(contentType, "application/json") {
			// Parse JSON body
			err := json.NewDecoder(r.Body).Decode(&newBook)
			if err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
			return
		}

		// Check if title or author is empty
		if newBook.Title == "" || newBook.Author == "" {
			response := Response{
				Result: "",
				Error:  "title or author cannot be empty",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		// Convert title and author to lowercase to perform a case-insensitive check
		lowerTitle := strings.ToLower(newBook.Title)
		lowerAuthor := strings.ToLower(newBook.Author)

		// Check if the book already exists in the library (case-insensitive)
		if bookInfo, exists := library.Books[lowerTitle]; exists && strings.ToLower(bookInfo.Book.Author) == lowerAuthor {
			// Return a message saying the book is already in the library
			response := Response{
				Result: "this book is already in the library",
				Error:  "",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		// Add the book to the library
		library.Books[lowerTitle] = &BookInfo{
			Book:     Book{Title: newBook.Title, Author: newBook.Author},
			Borrowed: false,
		}

		// Return success message
		response := Response{
			Result: fmt.Sprintf("added book %s by %s", lowerTitle, lowerAuthor),
			Error:  "",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	case http.MethodGet:
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")

		// Check if the title or author is empty
		if title == "" || author == "" {
			response := Response{
				Result: "",
				Error:  "title or author cannot be empty",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Search for the book in the library (using lowercased title)
		bookInfo, exists := library.Books[strings.ToLower(title)]

		// Check if the book exists and the author matches
		if !exists || strings.ToLower(bookInfo.Book.Author) != strings.ToLower(author) {
			response := Response{
				Result: "",
				Error:  "this book does not exist",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Check if the book is borrowed
		if bookInfo.Borrowed {
			response := Response{
				Result: "",
				Error:  "this book is borrowed",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Return the book information, with title and author in lowercase
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Book{
			Title:  strings.ToLower(bookInfo.Book.Title),
			Author: strings.ToLower(bookInfo.Book.Author),
		})

	case http.MethodPut:
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")

		// Check if the title or author is empty
		if title == "" || author == "" {
			response := Response{
				Result: "",
				Error:  "title or author cannot be empty",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Search for the book in the library (case-insensitive title search)
		bookInfo, exists := library.Books[strings.ToLower(title)]

		// Check if the book exists and the author matches
		if !exists || strings.ToLower(bookInfo.Book.Author) != strings.ToLower(author) {
			response := Response{
				Result: "",
				Error:  "this book does not exist",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Parse the request body to get the "borrow" field
		var body struct {
			Borrow *bool `json:"borrow"` // Use a pointer to detect missing "borrow" field
		}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil || body.Borrow == nil {
			response := Response{
				Result: "",
				Error:  "borrow value cannot be empty",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Handle borrowing
		if *body.Borrow {
			if bookInfo.Borrowed {
				response := Response{
					Result: "",
					Error:  "this book is already borrowed",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(response)
				return
			}
			// Mark the book as borrowed
			bookInfo.Borrowed = true
			response := Response{
				Result: "you have borrowed this book successfully",
				Error:  "",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Handle returning
		if !bookInfo.Borrowed {
			response := Response{
				Result: "",
				Error:  "this book is already in the library",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)
			return
		}
		// Mark the book as returned
		bookInfo.Borrowed = false
		response := Response{
			Result: "thank you for returning this book",
			Error:  "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	case http.MethodDelete:
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")

		// Check if the title or author is empty
		if title == "" || author == "" {
			response := Response{
				Result: "",
				Error:  "title or author cannot be empty",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Search for the book in the library (case-insensitive title search)
		bookInfo, exists := library.Books[strings.ToLower(title)]

		// Check if the book exists and the author matches
		if !exists || strings.ToLower(bookInfo.Book.Author) != strings.ToLower(author) {
			response := Response{
				Result: "",
				Error:  "this book does not exist",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Delete the book from the library
		delete(library.Books, strings.ToLower(title))

		// Return success response
		response := Response{
			Result: "successfully deleted",
			Error:  "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}

func main() {
	port := "4001"
	fmt.Printf("Starting server on port %s...\n", port)
	server := NewServer(port)
	server.Start()

}
