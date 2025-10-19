package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Decalre constant
const (
	ERROR400                   = "err-400" // Bad Request
	ERROR500                   = "err-500" // Internal Server Error
	WRONGREQUESTBODY           = "Invalid Request Body"
	REQUIREDREQUESTBODYMISSING = "Required Request Body Missing"
)

// Books Data
var BooksData []BookDTO

// Cart Data
var CartData []CheckOut

// Cart Items
var carts []CartItem

// Book Structure
type BookDTO struct {
	ID          string  `json:"id" validate:"required"`           // Unique identifier for the book
	Title       string  `json:"title" validate:"required"`        // Title of the book
	Author      string  `json:"author" validate:"required"`       // Author's name
	Publisher   string  `json:"publisher" validate:"required"`    // Publisher's name
	PublishedAt string  `json:"published_at" validate:"required"` // Publication date (could be string or time.Time)
	ISBN        string  `json:"isbn" validate:"required"`         // ISBN number
	Pages       int     `json:"pages" validate:"required"`        // Number of pages
	Language    string  `json:"language" validate:"required"`     // Language of the book
	Price       float64 `json:"price" validate:"required"`        // Price of the book
}

// Update Book
type PatchBookDTO struct {
	ID       string  `json:"id" validate:"required"`       // Unique identifier for the book
	Pages    int     `json:"pages" validate:"required"`    // Number of pages
	Language string  `json:"language" validate:"required"` // Language of the book
	Price    float64 `json:"price" validate:"required"`    // Price of the book
}

// success dto
type SuccessDTO struct {
	SuccessCode    string `json:"status_code"`
	SuccessMessage string `json:"status_message,omitempty"`
	Total          int    `json:"total,omitempty"`
	CustomMessage  any    `json:"response,omitempty"`
}

// Error DTO - DTO to display error in response
type ErrorDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// CartItem represents a book added to the user's shopping cart.
type AddItem struct {
	BookID string `json:"book_id"`
}

// Cart
type CartItem struct {
	BookID   string  `json:"book_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CheckOut struct {
	Cart  []CartItem `json:"cart"`
	Total float64    `json:"total"`
}

// Main
func main() {
	//Initialization of gin-gonic
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info().Msgf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	//Register endpoint
	r.POST("/book/create", CreateBook)
	r.GET("/book/books", GetBooks)
	r.GET("/book/book/:id", GetBook)
	r.DELETE("/book/book/:id", DeleteBook)
	r.PATCH("/book/book/:id", PatchBook)
	r.PUT("/book/book/:id", PutBook)
	r.POST("/cart/cart", AddToCart)
	r.GET("/cart/viewcart", ViewCart)
	//Book Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// Save a new record
func CreateBook(c *gin.Context) {
	//Declare DTO for Book
	var book BookDTO
	//BindJSON
	jsonRes := Bindjson(c, &book)
	if jsonRes {
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
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
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
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
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &book)
	if jsonValid {
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

// BindJson Structure
func Bindjson(c *gin.Context, data any) bool {
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Data :: %+v", c.Request.URL, c.Request.Method, data)
	//Check JSON according to given structure
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    ERROR400,
			ErrorMessage: WRONGREQUESTBODY,
		})
		//Print Log
		log.Error().Msgf(WRONGREQUESTBODY + " :: " + err.Error())
		return true
	}
	return false
}

// Validate JSON Structure
func ValidateJson(c *gin.Context, data any) bool {
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Data :: %+v", c.Request.URL, c.Request.Method, data)
	//Validate JSON according to given structure
	validation := validator.New()
	if err := validation.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    ERROR400,
			ErrorMessage: REQUIREDREQUESTBODYMISSING,
		})
		//Print Log
		log.Error().Msgf(REQUIREDREQUESTBODYMISSING + " :: " + err.Error())
		return true
	}
	return false
}

// Check already exists book record
func isBookExists(id, title string) bool {
	for i := 0; i < len(BooksData); i++ {
		if (BooksData[i].ID == id) || (BooksData[i].Title == title) || (BooksData[i].Title == "") {
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

// Cart
func AddToCart(c *gin.Context) {
	//Declare DTO for Book
	var item AddItem
	//BindJSON
	jsonRes := Bindjson(c, &item)
	if jsonRes {
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &item)
	if jsonValid {
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, item)
	//Check Already Exists Book Record
	if !isBookExists(item.BookID, "") {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Book having book ID - %s and Title - %s doesn't exists", item.BookID, findBookTitle(item.BookID)),
		})
		return
	}
	CheckCartItem(item)
	log.Info().Msgf("Cart Item :: %+v", carts)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "Book Added in Your Cart",
		CustomMessage:  carts,
	})
}

func CheckCartItem(item AddItem) {
	var cart CartItem
	bookDetails := GetBookById(item.BookID)
	cart.BookID = bookDetails.ID
	cart.Quantity = 1
	cart.Price = bookDetails.Price
	if len(carts) == 0 {
		carts = append(carts, cart)
	} else {
		for i := 0; i < len(carts); i++ {
			if carts[i].BookID == item.BookID {
				carts[i].Quantity = carts[i].Quantity + 1
				return
			}
		}
		carts = append(carts, cart)
	}
}

// View Cart
func ViewCart(c *gin.Context) {
	//Store Total checkout value
	var total float64
	//Calculate Total Price
	for i := 0; i < len(carts); i++ {
		total = total + (float64(carts[i].Quantity) * carts[i].Price)
	}

	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s ", c.Request.URL, c.Request.Method)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode: fmt.Sprintf("%d", http.StatusOK),
		Total:       len(CartData),
		CustomMessage: CheckOut{
			Cart:  carts,
			Total: total,
		},
	})
}
