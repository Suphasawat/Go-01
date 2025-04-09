package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
}

var books = []Book{}

 func main() {
  app := fiber.New()
  
  books = append(books, Book{ID: 1, Title: "Nine", Author: "NineLnwZa007"})
  books = append(books, Book{ID: 2, Title: "N1", Author: "NineLnwZa007"})

  app.Get("/books", getBooks)
  app.Get("/books/:id", getBook)
  app.Post("/books", createBook)
  app.Put("/books/:id", updateBook)
  app.Delete("/books/:id", deleteBook)

  app.Listen((":8080"))
}

