package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Database
var Users []ModelUser
var UserAddress = make(map[string][]RequestUserAddressToAdd)

// User Table
type ModelUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// DTO - Data Transfer Object
// Login Request
type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Request to Add Address
type RequestUserAddressToAdd struct {
	Add string `json:"address"`
}

// Request to Get User Profile
type ResponseUserProfile struct {
	Name     string                    `json:"name"`
	Username string                    `json:"username"`
	Address  []RequestUserAddressToAdd `json:"address"`
	Books    []ModelBook               `json:"books"`
	Cart     []CheckOut                `json:"cart"`
}

// Save a new record
func CreateUser(c *gin.Context) {
	//Request User Data
	var user ModelUser
	//BindJSON
	jsonRes := Bindjson(c, &user)
	if jsonRes {
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &user)
	if jsonValid {
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, user)
	//Check Already Exists Book Record
	if isUserExists(user.Username) {
		//Response
		c.JSON(http.StatusConflict, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Username - %s does not have an account", user.Username),
		})
		//Print Incoming Request
		log.Warn().Msgf("Username - %s does not have an account", user.Username)
		return
	}
	// Save user in table
	Users = append(Users, user)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "New User Created Successfully",
		CustomMessage: map[string]string{
			"username": user.Username,
			"password": strings.Repeat("*", len(user.Password)),
		},
	})
}

// Login User
func LoginUser(c *gin.Context) {
	//Request for User Login
	var user RequestLogin
	//BindJSON
	jsonRes := Bindjson(c, &user)
	if jsonRes {
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &user)
	if jsonValid {
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, user)
	//Check Already Exists User Record
	if !isUserExists(user.Username) {
		//Response
		c.JSON(http.StatusNotFound, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Username - %s don't have a account", user.Username),
		})
		//Print Incoming Request
		log.Warn().Msgf("Username - %s don't have a account", user.Username)
		return
	}
	//Check user credentials
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == user.Username && Users[i].Password == user.Password {
			//Generate Token
			token, _ := GenerateToken(user.Username)
			log.Info().Msgf("Login SuccessFully for Username - %s and Password - %s", user.Username, strings.Repeat("*", len(user.Password)))
			c.JSON(http.StatusOK, SuccessDTO{
				SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
				SuccessMessage: fmt.Sprintf("Login SuccessFully for Username - %s", user.Username),
				CustomMessage:  token,
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, ErrorDTO{
		ErrorCode:    fmt.Sprintf("%d", http.StatusUnauthorized),
		ErrorMessage: "Invalid username or password",
	})
}

// Add User Address
func UserAdd(c *gin.Context) {
	//To Store Active Username
	activeUsername := c.GetString("username")
	//Declare DTO for Book
	var add RequestUserAddressToAdd
	//BindJSON
	jsonRes := Bindjson(c, &add)
	if jsonRes {
		return
	}
	//Validate JSON
	jsonValid := ValidateJson(c, &add)
	if jsonValid {
		return
	}
	//Print Incoming Request
	log.Info().Msgf("Username - %s :: Requested :: Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", activeUsername, c.Request.URL, c.Request.Method, add)
	// Address
	if isAddExists(add, activeUsername) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Address - %s already exists", add.Add),
		})
		//Print Incoming Request
		log.Warn().Msgf("Address - %s already exists", add.Add)
		return
	}
	if val, ok := UserAddress[activeUsername]; ok {
		UserAddress[activeUsername] = append(val, add)
	} else {
		UserAddress[activeUsername] = []RequestUserAddressToAdd{
			{Add: add.Add},
		}
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "User Address added Successfully",
		CustomMessage:  UserAddress[activeUsername],
	})
}

func GetProfile(c *gin.Context) {
	//Username
	activeUsername := c.Params.ByName("username")
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Username  :: %v", c.Request.URL, c.Request.Method, activeUsername)
	//IsUserExists
	if !isUserExists(activeUsername) {
		//Response
		c.JSON(http.StatusNotFound, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Username - %s does not have an account", activeUsername),
		})
		//Print Incoming Request
		log.Warn().Msgf("Username - %s does not have an account", activeUsername)
		return
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Profile of Username %s", activeUsername),
		CustomMessage: ResponseUserProfile{
			Name:     GetName(activeUsername),
			Username: activeUsername,
			Address:  UserAddress[activeUsername],
			Books:    BooksData[activeUsername],
			Cart:     CartData[activeUsername],
		},
	})
}

// Check Is User Exists
func isUserExists(username string) bool {
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username {
			return true
		}
	}
	return false
}

// Check Is User Exists
func isAddExists(add RequestUserAddressToAdd, username string) bool {
	if val, ok := UserAddress[username]; ok {
		for i := 0; i < len(val); i++ {
			if val[i] == add {
				return true
			}
		}
	}
	return false
}

// Get Name
func GetName(username string) string {
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username {
			return Users[i].Name
		}
	}
	return ""
}
