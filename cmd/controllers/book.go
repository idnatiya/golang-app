package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/models"
)

type StoreBookType struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.First(&book, c.Param("bookID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func StoreBook(c *gin.Context) {
	var requestData StoreBookType
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  requestData.Title,
		Author: requestData.Author,
	}

	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

type UpdateBookType struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.First(&book, c.Param("bookID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateBookRequest UpdateBookType
	if err := c.ShouldBindJSON(&updateBookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = updateBookRequest.Title
	book.Author = updateBookRequest.Author

	models.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   book,
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.First(&book, c.Param("bookID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusNoContent, nil)
}
