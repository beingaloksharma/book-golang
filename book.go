package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Books Data
var BooksData []BookDTO

// Book Structure
type BookDTO struct {
	ID          string  `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"author" validate:"required"`
	Publisher   string  `json:"publisher" validate:"required"`
	PublishedAt string  `json:"published_at" validate:"required"`
	ISBN        string  `json:"isbn" validate:"required"`
	Pages       int     `json:"pages" validate:"required"`
	Language    string  `json:"language" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

// Update Book
type PatchBookDTO struct {
	ID       string  `json:"id" validate:"required"`
	Pages    int     `json:"pages" validate:"required"`
	Language string  `json:"language" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

// Save a new record
func CreateBook(c *gin.Context) {
	//Declare DTO for Book
	var book BookDTO
	//BindJSON
	jsonRes := Bindjson(c, &book)
	if jsonRes {
		//Print Incoming Request
		log.Error().Msgf("Bind JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
		//Print Incoming Request
		log.Error().Msgf("Validate JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
	//Check Already Exists Book Record
	if isBookExists(book.ID, book.Title) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s and Title - %s already exists", book.ID, book.Title),
		})
		return
	}
	BooksData = append(BooksData, book)
	log.Info().Msgf("Book Data :: %+v", BooksData)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "Book Record Created",
		CustomMessage:  BooksData,
	})
}

// Get All Books
func GetBooks(c *gin.Context) {
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s ", c.Request.URL, c.Request.Method)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:   fmt.Sprintf("%d", http.StatusOK),
		Total:         len(BooksData),
		CustomMessage: BooksData,
	})
}

// Get Book By Id
func GetBook(c *gin.Context) {
	//Book ID
	id := c.Params.ByName("id")
	//Book Details
	var book BookDTO
	book.ID = id
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Book ID :: %s", c.Request.URL, c.Request.Method, id)
	//IsBookExists
	if !isBookExists(id, "") {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s  don't exists", book.ID),
		})
		return
	}
	//If Book exists, save the response in book
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == book.ID {
			book.ID = BooksData[i].ID
			book.Title = BooksData[i].Title
			book.Author = BooksData[i].Author
			book.PublishedAt = BooksData[i].PublishedAt
			book.Publisher = BooksData[i].Publisher
			book.ISBN = BooksData[i].ISBN
			book.Language = BooksData[i].Language
			book.Pages = BooksData[i].Pages
			book.Price = BooksData[i].Price
		}
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:   fmt.Sprintf("%d", http.StatusOK),
		CustomMessage: book,
	})
}

// Put Book By Id
func PatchBook(c *gin.Context) {
	//Book ID
	id := c.Params.ByName("id")
	//Book Details
	var book PatchBookDTO
	//BindJSON
	jsonRes := Bindjson(c, &book)
	if jsonRes {
		//Print Incoming Request
		log.Error().Msgf("Bind JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
		//Print Incoming Request
		log.Error().Msgf("Validate JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Book :: %+v", c.Request.URL, c.Request.Method, book)
	//check Id
	if book.ID != id {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("URL Param ID - %s and Book ID - %s are different - ", book.ID, id),
		})
		return
	}
	//IsBookExists
	if !isBookExists(book.ID, "") {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s  don't exists", book.ID),
		})
		return
	}
	//If Book exists, save the response in book
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == book.ID {
			BooksData[i].Price = book.Price
			BooksData[i].Pages = book.Pages
			BooksData[i].Language = book.Language
		}
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book Id %s record has been updated successfully", book.ID),
		CustomMessage:  GetBookById(book.ID),
	})
}

func PutBook(c *gin.Context) {
	//Book ID
	id := c.Params.ByName("id")
	//Book Details
	var book BookDTO
	//BindJSON
	jsonRes := Bindjson(c, &book)
	if jsonRes {
		//Print Incoming Request
		log.Error().Msgf("Bind JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
		//Print Incoming Request
		log.Error().Msgf("Validate JSON :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, book)
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Book  :: %v", c.Request.URL, c.Request.Method, book)
	//check Id
	if book.ID != id {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("URL Param ID - %s and Book ID - %s are different - ", book.ID, id),
		})
		return
	}
	//IsBookExists
	if !isBookExists(book.ID, "") {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s  don't exists", book.ID),
		})
		return
	}
	//If Book exists, save the response in book
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == book.ID {
			BooksData[i].ID = book.ID
			BooksData[i].Title = book.Title
			BooksData[i].Author = book.Author
			BooksData[i].PublishedAt = book.PublishedAt
			BooksData[i].Publisher = book.Publisher
			BooksData[i].ISBN = book.ISBN
			BooksData[i].Language = book.Language
			BooksData[i].Pages = book.Pages
			BooksData[i].Price = book.Price
		}
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book Id %s record has been updated successfully", book.ID),
		CustomMessage:  GetBookById(book.ID),
	})
}

// Delete Book By Id
func DeleteBook(c *gin.Context) {
	//Book ID
	id := c.Params.ByName("id")
	//Book Details
	var book BookDTO
	book.ID = id
	title := findBookTitle(id)
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Book ID :: %s --- Title :: %s", c.Request.URL, c.Request.Method, id, title)
	//IsBookExists
	if !isBookExists(book.ID, "") {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s  doesn't exists", book.ID),
		})
		return
	}
	//Delete Book
	deleteBookById(id)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book having book ID - %s and Title - %s is deleted", book.ID, title),
	})
}

// Check already exists book record
func isBookExists(id, title string) bool {
	for i := 0; i < len(BooksData); i++ {
		if (BooksData[i].ID == id) || (BooksData[i].Title == title) {
			log.Warn().Msgf("Book having book ID - %s and Title - %s already exists", id, findBookTitle(id))
			return true
		}
	}
	return false
}

// Check already exists book record
func deleteBookById(id string) {
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == id {
			BooksData = append(BooksData[:i], BooksData[i+1:]...)
		}
	}
}

// Return Book Title By ID
func findBookTitle(id string) string {
	var title string
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == id {
			title = BooksData[i].Title
		}
	}
	return title
}

// Get Book By Id
func GetBookById(id string) BookDTO {
	var book BookDTO
	for i := 0; i < len(BooksData); i++ {
		if BooksData[i].ID == id {
			book.ID = BooksData[i].ID
			book.Title = BooksData[i].Title
			book.Author = BooksData[i].Author
			book.PublishedAt = BooksData[i].PublishedAt
			book.Publisher = BooksData[i].Publisher
			book.ISBN = BooksData[i].ISBN
			book.Language = BooksData[i].Language
			book.Pages = BooksData[i].Pages
			book.Price = BooksData[i].Price
		}
	}
	return book
}
