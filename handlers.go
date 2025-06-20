package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func CreateBook(c *gin.Context) {
	var b Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(context.Background(),
		"INSERT INTO books (title, author) VALUES ($1, $2)", b.Title, b.Author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book created"})
}

func GetBooks(c *gin.Context) {
	rows, err := DB.Query(context.Background(), "SELECT id, title, author FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch books"})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			continue
		}
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	row := DB.QueryRow(context.Background(), "SELECT id, title, author FROM books WHERE id=$1", id)

	var b Book
	err := row.Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := DB.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
