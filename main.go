package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)
 type Book struct {
  ID int `json:"id"` 
  Title string `json:"title"`
  Author string `json:"author"`
}

var books = []Book{}

func cheakMiddleware(c *fiber.Ctx) error {
  user := c.Locals("user").(*jwt.Token)
  claims := user.Claims.(jwt.MapClaims)

  if claims["role"] != "admin" {
    return fiber.ErrUnauthorized
  }

  return c.Next()
}

 func main() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Load .env error")
  }

  engine := html.New("./views", ".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })
  
  books = append(books, Book{ID: 1, Title: "Nine", Author: "NineLnwZa007"})
  books = append(books, Book{ID: 2, Title: "N1", Author: "NineLnwZa007"})

  app.Post("login", login)

  app.Use(jwtware.New(jwtware.Config{
    SigningKey: []byte(os.Getenv("JWT_SECRET")),
  }))

  app.Use(cheakMiddleware)
  
  app.Get("/books", getBooks)
  app.Get("/books/:id", getBook)
  app.Post("/books", createBook)
  app.Put("/books/:id", updateBook)
  app.Delete("/books/:id", deleteBook)

  app.Post("/upload", uploadFile)
  app.Get("/test-html", testHTML)

  app.Get("/config", getENV)

  app.Listen((":8080"))
}

func uploadFile(c *fiber.Ctx) error {
  file, err := c.FormFile("image")

  if err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
  }

  err = c.SaveFile(file, "./uploads/" + file.Filename)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
  }

  return c.SendString("File uploaded successfully: " + file.Filename)
}

func testHTML(c *fiber.Ctx) error {
  return c.Render("index", fiber.Map{
    "title": "Test HTML",
  })
}

func getENV(c *fiber.Ctx) error {
  secret := os.Getenv("SECRET")

  if secret == "" {
    secret = "default_secret"
  }

  return c.JSON(fiber.Map{
    "SECRET": os.Getenv("SECRET"),
  })
}

type User struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

var memberUser = User {
  Email : "user@example.com",
  Password: "password123",
}

func login(c *fiber.Ctx) error {
  user := new(User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if user.Email != memberUser.Email || user.Password != memberUser.Password {
      return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
    }

     // Create token
     token := jwt.New(jwt.SigningMethodHS256)

     // Set claims
     claims := token.Claims.(jwt.MapClaims)
     claims["email"] = user.Email
     claims["role"] = "admin" // example role
     claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
 
     // Generate encoded token
     t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
     if err != nil {
       return c.SendStatus(fiber.StatusInternalServerError)
     }

    return c.JSON(fiber.Map{
      "message": "Login successful",
      "token": t,
    })
  }
