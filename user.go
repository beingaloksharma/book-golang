package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Users Data
var Users []User

// To Store Current Username
var activeUser string

// User Address
var UserAddress []Address

// User
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Address struct {
	Add string `json:"address"`
}

type UserAddr struct {
	Username string    `json:"username"`
	UserAdds []Address `json:"address"`
}

type UserProfile struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	UserAdds []Address `json:"address"`
}

// Save a new record
func CreateUser(c *gin.Context) {
	//Declare DTO for Book
	var user User
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
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Username - %s already exists", user.Username),
		})
		return
	}
	Users = append(Users, user)
	log.Info().Msgf("User Data :: %+v", Users)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "User Created Successfully",
		CustomMessage: map[string]string{
			"username": user.Username,
			"password": strings.Repeat("*", len(user.Password)),
		},
	})
}

func LoginUser(c *gin.Context) {
	var user LoginDetails
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
	if !isUserExists(user.Username) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Username - %s doesn't exists", user.Username),
		})
		return
	}
	//Check user credentials
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == user.Username && Users[i].Password == user.Password {
			activeUser = user.Username
			log.Info().Msgf("Login SuccessFully for Username - %s", user.Username)
			c.JSON(http.StatusOK, SuccessDTO{
				SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
				SuccessMessage: fmt.Sprintf("Login SuccessFully for Username - %s", user.Username),
			})
			return
		}
	}
	//Response
	c.JSON(http.StatusOK, ErrorDTO{
		ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
		ErrorMessage: "You don't have account with us",
	})
}

func UserAdd(c *gin.Context) {
	//Declare DTO for Book
	var add Address
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
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s  --- Request Body :: %+v", c.Request.URL, c.Request.Method, add)
	// Address
	if isAddExists(add, activeUser) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
			ErrorMessage: fmt.Sprintf("Address - %s already exists", add.Add),
		})
		return
	}
	UserAddress = append(UserAddress, add)
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: "User Address added Successfully",
		CustomMessage: UserAddr{
			Username: activeUser,
			UserAdds: UserAddress,
		},
	})
}

func GetProfile(c *gin.Context) {
	//Username
	username := c.Params.ByName("username")
	//Print Incoming Request
	log.Info().Msgf("Request URL :: %s --- Request Method :: %s --- Username  :: %v", c.Request.URL, c.Request.Method, username)
	//IsUserExists
	if !isUserExists(username) {
		//Response
		c.JSON(http.StatusOK, ErrorDTO{
			ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
			ErrorMessage: fmt.Sprintf("Username - %s  don't exists", username),
		})
		return
	}
	//Response
	c.JSON(http.StatusOK, SuccessDTO{
		SuccessCode:    fmt.Sprintf("%d", http.StatusOK),
		SuccessMessage: fmt.Sprintf("Profile of Username %s", username),
		CustomMessage: UserProfile{
			Name:     GetName(username),
			Username: username,
			UserAdds: UserAddress,
		},
	})
}

// Check Is User Exists
func isUserExists(username string) bool {
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username {
			log.Warn().Msgf("Username  %s already exists", username)
			return true
		}
	}
	return false
}

// Check Is User Exists
func isAddExists(add Address, username string) bool {
	for i := 0; i < len(Users); i++ {
		for j := 0; j < len(UserAddress); j++ {
			if UserAddress[j].Add == add.Add && Users[i].Username == username {
				log.Warn().Msgf("Address  %s already exists", add)
				return true
			}
		}
	}
	return false
}

// Get Name
func GetName(username string) string {
	var name string
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == username {
			name = Users[i].Name
		}
	}
	return name
}
