package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Database
var CartData = map[string][]CheckOut{}
var carts = make(map[string][]CartItem)

// Cart Item
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
	Username string     `json:"-"`
	Cart     []CartItem `json:"cart"`
	Total    float64    `json:"total"`
}

// Add to Cart handler
// @Schemes http
// @Description Add Book to Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request_body body AddItem true "Book to add to Cart"
// @Success 200 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /cart [post]
func AddToCart(c *gin.Context) {
	activeUsername := c.GetString("username")
	var item AddItem

	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Info().Msgf("Username :: %s :: Request URL :: %s --- Method :: %s --- Request Body :: %+v", activeUsername, c.Request.URL, c.Request.Method, item)

	if !isBookExists(item.BookID, "", activeUsername) {
		bookTitle := findBookTitle(item.BookID, activeUsername)
		c.JSON(http.StatusConflict, gin.H{
			"error_code":    fmt.Sprintf("%d", http.StatusConflict),
			"error_message": fmt.Sprintf("Book having book ID - %s and Title - %s doesn't exist", item.BookID, bookTitle),
		})
		log.Warn().Msgf("Requested Book ID - %s and Title - %s does not exist", item.BookID, bookTitle)
		return
	}

	CheckCartItem(c, item)

	log.Info().Msgf("Cart Items: %+v", carts[activeUsername])

	c.JSON(http.StatusOK, gin.H{
		"success_code":    fmt.Sprintf("%d", http.StatusOK),
		"success_message": "Book added to your cart",
		"cart":            carts[activeUsername],
	})
}

// Add or update cart item
func CheckCartItem(c *gin.Context, item AddItem) {
	activeUsername := c.GetString("username")
	bookDetails := GetBookById(item.BookID, activeUsername)

	userCart := carts[activeUsername]
	for i := range userCart {
		if userCart[i].BookID == item.BookID {
			userCart[i].Quantity++
			carts[activeUsername] = userCart
			return
		}
	}

	// Item not found, append new item
	newCartItem := CartItem{
		BookID:   bookDetails.ID,
		Quantity: 1,
		Price:    bookDetails.Price,
	}
	carts[activeUsername] = append(userCart, newCartItem)
}

// View Cart handler
// @Schemes http
// @Description View Cart Items
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /cart [get]
func ViewCart(c *gin.Context) {
	activeUsername := c.GetString("username")
	userCart := carts[activeUsername]

	//Total Price
	var totalPrice float64
	for _, item := range userCart {
		totalPrice += float64(item.Quantity) * item.Price
	}

	//Total Quantity
	var totalQUnatity int
	for _, item := range userCart {
		totalQUnatity += item.Quantity
	}

	log.Info().Msgf("Request URL :: %s --- Method :: %s", c.Request.URL, c.Request.Method)

	c.JSON(http.StatusOK, gin.H{
		"success_code": fmt.Sprintf("%d", http.StatusOK),
		"total_items":  totalQUnatity,
		"total_price":  totalPrice,
		"cart":         userCart,
	})
}

// Get cart checkout details by username
func GetCartDetailsByUserName(username string) CheckOut {
	userCart := carts[username]
	var total float64
	for _, item := range userCart {
		total += float64(item.Quantity) * item.Price
	}
	//Return
	return CheckOut{
		Username: username,
		Cart:     userCart,
		Total:    total,
	}
}
