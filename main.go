package main

import (
	"strconv"

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

  app.Listen((":8080"))
}

func getBooks(c *fiber.Ctx) error  {
  return c.JSON(books)  
}

func getBook(c *fiber.Ctx) error {
  bookId, err := strconv.Atoi(c.Params("id"))

  if err != nil { 
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
  }

  for _, book := range books {
    if book.ID == bookId {
      return c.JSON(book)
    }
  }
  return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
  book := new(Book)
  c.BodyParser(book)
  return c.JSON(book)
}