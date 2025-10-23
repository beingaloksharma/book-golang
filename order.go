package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Database
var OrderData = make(map[string][]Orders)

// Order Table
type Orders struct {
	OrderID   string     `json:"order_id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	Cart      []CartItem `json:"cart" validate:"required"`
	Address   string     `json:"address" validate:"required"`
	OrderDate string     `json:"order_date" validate:"required"`
	Total     float64    `json:"total" validate:"required"`
	Status    Status     `json:"status"`
	Payment   Payment    `json:"payment,omitempty"`
}

// Status structure for tracking
type Status struct {
	Type            string `json:"type" validate:"required,oneof=pending shipped delivered"`
	Reason          string `json:"reason,omitempty"`
	TrackingNumber  string `json:"tracking_number,omitempty"`
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
}

// Payment structure for tracking
type Payment struct {
	Paid      bool   `json:"paid"`
	Method    string `json:"method,omitempty" validate:"required,oneof=UPI Card COD"`
	PaidOn    string `json:"paid_on,omitempty"`
	Reference string `json:"reference,omitempty"`
}

// Create new order
// @Schemes http
// @Description Place an Order
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} SuccessDTO
// @Failure 400 {object} ErrorDTO
// @Router /order [post]
func OrderDetails(c *gin.Context) {
	activeUsername := c.GetString("username")
	log.Info().Msgf("Username - %s :: Requested URL - %s :: Method - %s", activeUsername, c.Request.URL, c.Request.Method)

	// Validate address
	if _, ok := UserAddress[activeUsername]; !ok || len(UserAddress[activeUsername]) == 0 {
		respondError(c, http.StatusBadRequest, "User must add an address before placing order")
		log.Warn().Msgf("No address found for user - %s", activeUsername)
		return
	}

	// Validate cart
	checkout := GetCartDetailsByUserName(activeUsername)
	if len(checkout.Cart) == 0 {
		respondError(c, http.StatusBadRequest, "Cart is empty, please add items before placing an order")
		log.Warn().Msgf("Cart empty for user - %s", activeUsername)
		return
	}

	// Create order
	order := Orders{
		OrderID:   generateOrderNumber(),
		Name:      GetName(activeUsername),
		Cart:      checkout.Cart,
		Address:   UserAddress[activeUsername][0].Add,
		OrderDate: time.Now().Format("2006-01-02 15:04:05"),
		Total:     checkout.Total,
		Status:    Status{Type: "pending", Reason: "Awaiting payment"},
		Payment:   Payment{Paid: false, Method: "", PaidOn: "", Reference: ""},
	}

	// Save in-memory
	OrderData[activeUsername] = append(OrderData[activeUsername], order)
	delete(carts, activeUsername) // Clear cart after order placement

	// Generate order summary text file
	err := generateOrderFile(activeUsername, order)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate order summary file")
	}

	// Respond success
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "Order placed successfully",
		CustomMessage:  order,
	})

	log.Info().Msgf("Order placed for user - %s :: %+v", activeUsername, order)
}

// Get all orders for user
// @Schemes http
// @Description Get all Orders for User
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} SuccessDTO
// @Router /order [get]
func GetOrders(c *gin.Context) {
	activeUsername := c.GetString("username")
	log.Info().Msgf("Username - %s :: Requested Orders", activeUsername)

	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "Fetched all orders successfully",
		Total:          len(OrderData[activeUsername]),
		CustomMessage:  OrderData[activeUsername],
	})
}

// Get specific order by ID
// @Schemes http
// @Description Get Order by ID
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} SuccessDTO
// @Failure 404 {object} ErrorDTO
// @Router /order/{id} [get]
func GetOrderByID(c *gin.Context) {
	activeUsername := c.GetString("username")
	orderID := c.Params.ByName("id")
	log.Info().Msgf("Username - %s :: Requested Order ID - %s", activeUsername, orderID)

	for _, o := range OrderData[activeUsername] {
		if o.OrderID == orderID {
			c.JSON(http.StatusOK, SuccessDTO{
				SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
				SuccessMessage: "Fetched order successfully",
				CustomMessage:  o,
			})
			return
		}
	}

	respondError(c, http.StatusNotFound, fmt.Sprintf("Order with ID %s not found", orderID))
	log.Warn().Msgf("Order not found - ID %s, User %s", orderID, activeUsername)
}

// Helper for JSON Error Responses
func respondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorDTO{
		ErrorCode:    fmt.Sprintf("%d", statusCode),
		ErrorMessage: message,
	})
}

// Generate Random Order Number
func generateOrderNumber() string {
	rand.Seed(time.Now().UnixNano())
	min := int64(1000000000)
	max := int64(9999999999)
	return fmt.Sprintf("%s%d", ORDERPREFIX, rand.Int63n(max-min+1)+min)
}

// Generate order summary file (text)
func generateOrderFile(username string, order Orders) error {
	// Create directory if not present
	dir := "orders_output"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	// Define file name
	fileName := fmt.Sprintf("%s_%s.txt", username, order.OrderID)
	filePath := filepath.Join(dir, fileName)

	// Prepare file content
	content := fmt.Sprintf(
		"Order Summary\n"+
			"==============\n"+
			"Username: %s\n"+
			"Name: %s\n"+
			"Order ID: %s\n"+
			"Order Date: %s\n"+
			"Delivery Address: %s\n\n"+
			"Items Ordered:\n",
		username, order.Name, order.OrderID, order.OrderDate, order.Address,
	)

	// Append cart details
	for _, item := range order.Cart {
		content += fmt.Sprintf("- BookID: %s | Quantity: %d | Price: %.2f\n",
			item.BookID, item.Quantity, item.Price)
	}
	content += fmt.Sprintf("\nTotal Amount: ₹%.2f\n", order.Total)

	// Append status info
	content += fmt.Sprintf("\nOrder Status: %s\n", order.Status.Type)
	if order.Status.Reason != "" {
		content += fmt.Sprintf("Status Reason: %s\n", order.Status.Reason)
	}
	if order.Status.TrackingNumber != "" {
		content += fmt.Sprintf("Tracking Number: %s\n", order.Status.TrackingNumber)
	}
	if order.Status.ShippingCarrier != "" {
		content += fmt.Sprintf("Shipping Carrier: %s\n", order.Status.ShippingCarrier)
	}

	// Append payment details
	content += fmt.Sprintf("\nPayment Details:\n")
	content += fmt.Sprintf("Paid: %t\n", order.Payment.Paid)
	content += fmt.Sprintf("Method: %s\n", order.Payment.Method)
	content += fmt.Sprintf("Paid On: %s\n", order.Payment.PaidOn)
	content += fmt.Sprintf("Reference: %s\n", order.Payment.Reference)

	content += fmt.Sprintf("\nGenerated On: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// Write to file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	log.Info().Msgf("Order summary file generated: %s", filePath)
	return nil
}

// Update order status
// @Schemes http
// @Description Update Order Status
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param status body Status true "Status Info"
// @Success 200 {object} SuccessDTO
// @Failure 404 {object} ErrorDTO
// @Router /order/{id}/status [put]
func UpdateOrderStatus(c *gin.Context) {
	activeUsername := c.GetString("username")
	orderID := c.Param("id")

	var newStatus Status
	if err := c.ShouldBindJSON(&newStatus); err != nil {
		respondError(c, http.StatusBadRequest, "Invalid status data")
		return
	}

	orders := OrderData[activeUsername]
	for i := range orders {
		if orders[i].OrderID == orderID {
			orders[i].Status = newStatus
			OrderData[activeUsername][i] = orders[i]

			// Update order summary file
			err := generateOrderFile(activeUsername, orders[i])
			if err != nil {
				log.Error().Err(err).Msg("Failed to update order summary file after status update")
			}

			c.JSON(http.StatusOK, SuccessDTO{
				SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
				SuccessMessage: "Order status updated successfully",
				CustomMessage:  orders[i],
			})
			log.Info().Msgf("Order status updated for user - %s :: %+v", activeUsername, orders[i])
			return
		}
	}

	respondError(c, http.StatusNotFound, fmt.Sprintf("Order with ID %s not found", orderID))
	log.Warn().Msgf("Order not found for status update - ID %s, User %s", orderID, activeUsername)
}

// Update order payment
// @Schemes http
// @Description Update Order Payment
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param payment body Payment true "Payment Info"
// @Success 200 {object} SuccessDTO
// @Failure 404 {object} ErrorDTO
// @Router /order/{id}/payment [put]
func UpdateOrderPayment(c *gin.Context) {
	activeUsername := c.GetString("username")
	orderID := c.Param("id")

	var newPayment Payment
	if err := c.ShouldBindJSON(&newPayment); err != nil {
		respondError(c, http.StatusBadRequest, "Invalid payment data")
		return
	}

	orders := OrderData[activeUsername]
	for i := range orders {
		if orders[i].OrderID == orderID {
			orders[i].Payment = newPayment
			OrderData[activeUsername][i] = orders[i]

			// Update order summary file
			err := generateOrderFile(activeUsername, orders[i])
			if err != nil {
				log.Error().Err(err).Msg("Failed to update order summary file after payment update")
			}

			c.JSON(http.StatusOK, SuccessDTO{
				SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
				SuccessMessage: "Order payment updated successfully",
				CustomMessage:  orders[i],
			})
			log.Info().Msgf("Order payment updated for user - %s :: %+v", activeUsername, orders[i])
			return
		}
	}

	respondError(c, http.StatusNotFound, fmt.Sprintf("Order with ID %s not found", orderID))
	log.Warn().Msgf("Order not found for payment update - ID %s, User %s", orderID, activeUsername)
}
