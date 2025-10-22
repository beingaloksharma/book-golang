package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Database
var BooksData = map[string][]ModelBook{}

// Book Model
type ModelBook struct {
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

// PatchBookDTO for updates
type PatchBookDTO struct {
	ID       string  `json:"id" validate:"required"`
	Pages    int     `json:"pages" validate:"required"`
	Language string  `json:"language" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

// Save a new record
// @Schemes http
// @Description Create a new Book record
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request_body body ModelBook true "Book Data"
// @Success 200 {object} SuccessDTO
// @Failure 409 {object} ErrorDTO
// @Router /book [post]
func CreateBook(c *gin.Context) {
	//To Store Active Username
	activeUsername := c.GetString("username")
	//Declare DTO for Book
	var book ModelBook
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
	log.Info().Msgf("Username - %s :: Requested :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", activeUsername, c.Request.URL, c.Request.Method, book)
	//Check Already Exists Book Record
	if isBookExists(book.ID, book.Title, activeUsername) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s and Title - %s already exists", book.ID, book.Title),
		})
		// Print Log
		log.Warn().Msgf("Book having book ID - %s and Title - %s already exists", book.ID, book.Title)
		return
	}
	Books := BooksData[activeUsername]
	Books = append(Books, book)
	BooksData[activeUsername] = Books
	log.Info().Msgf("Books Data for Username - %s :: %+v", activeUsername, Books)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "Book Record Created",
		CustomMessage:  Books,
	})
}

// Get All Books
// @Schemes http
// @Description Get All Books for User
// @Tags Book
// @Produce json
// @Security BearerAuth
// @Success 200 {object} SuccessDTO
// @Router /book/books [get]
func GetBooks(c *gin.Context) {
	//To Store Active Username
	activeUsername := c.GetString("username")
	//Print Incoming Request
	log.Info().Msgf("Username - %s :: Requested :: Request URL :: %s --- Request Method :: %s ", activeUsername, c.Request.URL, c.Request.Method)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:   fmt.Sprintf("%d", http.StatusOK),
		Total:         len(BooksData[activeUsername]),
		CustomMessage: BooksData[activeUsername],
	})
}

// Get Book By Id
// @Schemes http
// @Description Get Book By ID
// @Tags Book
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} SuccessDTO
// @Failure 404 {object} ErrorDTO
// @Router /book/{id} [get]
func GetBook(c *gin.Context) {
	//To Store Active Username
	activeUsername := c.GetString("username")
	//Book ID
	id := c.Params.ByName("id")
	//Book Model
	var book ModelBook
	//Print Incoming Request
	log.Info().Msgf("Username - %s :: Requested :: Request URL :: %s --- Request Method :: %s --- Book ID :: %s", activeUsername, c.Request.URL, c.Request.Method, id)
	//IsBookExists
	if !isBookExists(id, "", activeUsername) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s  don't exists", book.ID),
		})
		// Print Log
		log.Warn().Msgf("Book having book ID - %s don't exists", id)
		return
	}
	bookRecord := BooksData[activeUsername]
	//If Book exists, save the response in book
	for i := 0; i < len(bookRecord); i++ {
		if bookRecord[i].ID == id {
			book.ID = bookRecord[i].ID
			book.Title = bookRecord[i].Title
			book.Author = bookRecord[i].Author
			book.PublishedAt = bookRecord[i].PublishedAt
			book.Publisher = bookRecord[i].Publisher
			book.ISBN = bookRecord[i].ISBN
			book.Language = bookRecord[i].Language
			book.Pages = bookRecord[i].Pages
			book.Price = bookRecord[i].Price
		}
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:   fmt.Sprintf("%d", http.StatusOK),
		CustomMessage: book,
	})
}

// Put Book By Id
// @Schemes http
// @Description Partially update a Book by ID
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Param request_body body PatchBookDTO true "Book Patch Data"
// @Success 200 {object} SuccessDTO
// @Failure 400 {object} ErrorDTO
// @Failure 404 {object} ErrorDTO
// @Router /book/{id} [patch]
func PatchBook(c *gin.Context) {
	activeUsername := c.GetString("username")
	id := c.Params.ByName("id")
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
	log.Info().Msgf("Username - %s :: Request URL :: %s Method :: %s Book :: %+v", activeUsername, c.Request.URL, c.Request.Method, book)

	if book.ID != id {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusBadRequest),
			ErrorMessage: fmt.Sprintf("URL Parameter ID %s and Book ID %s are different", id, book.ID),
		})
		log.Warn().Msgf("URL param ID %s and book ID %s are different", id, book.ID)
		return
	}

	if !isBookExists(book.ID, "", activeUsername) {
		c.JSON(http.StatusNotFound, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book with ID %s does not exist", book.ID),
		})
		log.Warn().Msgf("Book with ID %s does not exist", book.ID)
		return
	}

	bookRecord := BooksData[activeUsername]
	for i := range bookRecord {
		if bookRecord[i].ID == book.ID {
			bookRecord[i].Pages = book.Pages
			bookRecord[i].Language = book.Language
			bookRecord[i].Price = book.Price
			break
		}
	}
	BooksData[activeUsername] = bookRecord

	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book ID %s record has been updated successfully", book.ID),
		CustomMessage:  GetBookById(book.ID, activeUsername),
	})
}

// @Schemes http
// @Description Fully update a Book by ID
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Param request_body body ModelBook true "Book Data"
// @Success 200 {object} SuccessDTO
// @Failure 400 {object} ErrorDTO
// @Failure 404 {object} ErrorDTO
// @Router /book/{id} [put]
func PutBook(c *gin.Context) {
	activeUsername := c.GetString("username")
	id := c.Params.ByName("id")
	var book ModelBook
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
	log.Info().Msgf("Username - %s :: Request URL :: %s Method :: %s Book :: %+v", activeUsername, c.Request.URL, c.Request.Method, book)

	if book.ID != id {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusBadRequest),
			ErrorMessage: fmt.Sprintf("URL Parameter ID %s and Book ID %s are different", id, book.ID),
		})
		log.Warn().Msgf("URL param ID %s and book ID %s are different", id, book.ID)
		return
	}

	if !isBookExists(book.ID, "", activeUsername) {
		c.JSON(http.StatusNotFound, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book with ID %s does not exist", book.ID),
		})
		log.Warn().Msgf("Book with ID %s does not exist", book.ID)
		return
	}

	bookRecord := BooksData[activeUsername]
	for i := range bookRecord {
		if bookRecord[i].ID == book.ID {
			bookRecord[i] = book
			break
		}
	}
	BooksData[activeUsername] = bookRecord

	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book ID %s record has been updated successfully", book.ID),
		CustomMessage:  GetBookById(book.ID, activeUsername),
	})
}

// Delete Book By Id
// @Schemes http
// @Description Delete a Book by ID
// @Tags Book
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} SuccessDTO
// @Failure 404 {object} ErrorDTO
// @Router /book/{id} [delete]
func DeleteBook(c *gin.Context) {
	activeUsername := c.GetString("username")
	id := c.Params.ByName("id")
	title := findBookTitle(id, activeUsername)
	log.Info().Msgf("Username - %s :: Request URL :: %s Method :: %s Book ID :: %s Title :: %s", activeUsername, c.Request.URL, c.Request.Method, id, title)

	if !isBookExists(id, "", activeUsername) {
		c.JSON(http.StatusNotFound, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Book with ID %s does not exist", id),
		})
		log.Warn().Msgf("Book with ID %s does not exist", id)
		return
	}

	deleteBookById(id, activeUsername)

	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Book with ID %s and Title %s has been deleted", id, title),
	})
}

// Check already exists book record
func isBookExists(id, title, username string) bool {
	bookRecord := BooksData[username]
	for i := range bookRecord {
		if id != "" && bookRecord[i].ID == id {
			return true
		}
		if title != "" && bookRecord[i].Title == title {
			return true
		}
	}
	return false
}

// Check already exists book record
func deleteBookById(id string, username string) {
	bookRecord := BooksData[username]
	for i := 0; i < len(bookRecord); i++ {
		if bookRecord[i].ID == id {
			bookRecord = append(bookRecord[:i], bookRecord[i+1:]...)
			BooksData[username] = bookRecord
			return
		}
	}
}

// Return Book Title By ID
func findBookTitle(id, username string) string {
	var title string
	bookRecord := BooksData[username]
	for i := 0; i < len(bookRecord); i++ {
		if bookRecord[i].ID == id {
			title = bookRecord[i].Title
		}
	}
	return title
}

// Get Book By Id
func GetBookById(id string, username string) ModelBook {
	var book ModelBook
	bookRecord := BooksData[username]
	for i := 0; i < len(bookRecord); i++ {
		if bookRecord[i].ID == id {
			book.ID = bookRecord[i].ID
			book.Title = bookRecord[i].Title
			book.Author = bookRecord[i].Author
			book.PublishedAt = bookRecord[i].PublishedAt
			book.Publisher = bookRecord[i].Publisher
			book.ISBN = bookRecord[i].ISBN
			book.Language = bookRecord[i].Language
			book.Pages = bookRecord[i].Pages
			book.Price = bookRecord[i].Price
		}
	}
	return book
}
