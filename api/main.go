package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity uint64 `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home page",
	})
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	b, err := getBook(id)

	if err == nil {
		c.IndentedJSON(http.StatusOK, b)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})

}

func getBook(id string) (b *Book, e error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	b, err := getBook(id)

	if err != nil || b.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not available"})
		return
	}

	b.Quantity -= 1

	c.IndentedJSON(http.StatusOK, b)

}

func returnBook(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "malformed id"})
		return
	}

	b, err := getBook(id)

	if err != nil {
		fmt.Println(b)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not avalilabe book"})
		return
	}

	b.Quantity += 1

	c.IndentedJSON(http.StatusOK, b)

}

func main() {
	router := gin.Default()

	router.GET("/", homepage)

	router.GET("/book/:id", getBookById)

	router.GET("/books", getBooks)

	router.POST("/newbook", createBook)

	router.PATCH("/getbook/", checkoutBook)

	router.PATCH("/returnbook/:id", returnBook)

	router.Run("localhost:8080")
}
