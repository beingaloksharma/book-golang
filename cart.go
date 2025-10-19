package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Cart Data
var CartData []CheckOut

// Cart Items
var carts []CartItem

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

// Checkout
type CheckOut struct {
	Cart  []CartItem `json:"cart"`
	Total float64    `json:"total"`
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
		Total:       len(carts),
		CustomMessage: CheckOut{
			Cart:  carts,
			Total: total,
		},
	})
}
