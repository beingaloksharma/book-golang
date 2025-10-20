package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Decalre constant
const (
	ERROR400                   = "err-400" // Bad Request
	ERROR401                   = "err-401" // Unauthorized
	ERROR500                   = "err-500" // Internal Server Error
	WRONGREQUESTBODY           = "Invalid Request Body"
	REQUIREDREQUESTBODYMISSING = "Required Request Body Missing"
	UNAUTHORIZED               = "Access is Unauthorized" // Unauthorized
)

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

// Main
func main() {
	//Initialization of gin-gonic
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info().Msgf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": Users,
		})
	})
	r.POST("/signup", CreateUser)
	r.POST("/signin", LoginUser)
	r.Use(JwtAuthMiddleware())
	user := r.Group("/user")
	{
		user.POST("/user/address", UserAdd)
		user.GET("/user/profile/:username", GetProfile)
	}
	//Register endpoint
	book := r.Group("/book")
	{
		book.POST("", CreateBook)
		book.GET("/books", GetBooks)
		book.GET("/:id", GetBook)
		book.DELETE("/:id", DeleteBook)
		book.PATCH("/:id", PatchBook)
		book.PUT("/:id", PutBook)
	}
	cart := r.Group("/cart")
	{
		cart.POST("", AddToCart)
		cart.GET("", ViewCart)
	}
	//Book Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Start Server To Listen
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}

// BindJson Structure
func Bindjson(c *gin.Context, data any) bool {
	//Print Incoming Request
	//log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Data :: %+v", c.Request.URL, c.Request.Method, data)
	//Check JSON according to given structure
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    ERROR400,
			ErrorMessage: WRONGREQUESTBODY,
		})
		//Print Log
		log.Error().Msgf("%s :: %s ", WRONGREQUESTBODY, err.Error())
		return true
	}
	return false
}

// Validate JSON Structure
func ValidateJson(c *gin.Context, data any) bool {
	//Print Incoming Request
	//log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Data :: %+v", c.Request.URL, c.Request.Method, data)
	//Validate JSON according to given structure
	validation := validator.New()
	if err := validation.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorDTO{
			ErrorCode:    ERROR400,
			ErrorMessage: REQUIREDREQUESTBODYMISSING,
		})
		//Print Log
		log.Error().Msgf("%s :: %s ", REQUIREDREQUESTBODYMISSING, err.Error())
		return true
	}
	return false
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			log.Error().Msgf("User  %s ", UNAUTHORIZED)
			c.JSON(http.StatusUnauthorized, ErrorDTO{
				ErrorCode:    ERROR401,
				ErrorMessage: UNAUTHORIZED,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
