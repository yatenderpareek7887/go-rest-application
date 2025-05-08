package controllers

import (
	userdto "go-rest-application/src/dto/user"
	genricerror "go-rest-application/src/genric_error"
	"go-rest-application/src/services"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		for _, r := range value {
			if !unicode.IsSpace(r) {
				return true
			}
		}
		return false
	})
}

// GetAllUsers
// @Summary Get all users
// @Description Get list of all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} user.User
// @Failure 500 {object} genricerror.ErrorResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser
// @Summary Create a new user
// @Description Create a new user with provided details
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  user body userdto.CreateUserRequest true "User to create"
// @Success 201 {object} user.User
// @Failure 400 {object} genricerror.ErrorResponse
// @Failure 500 {object} genricerror.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var userDto userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	// Validate struct fields
	if err := validate.Struct(userDto); err != nil {
		c.JSON(http.StatusBadRequest, genricerror.ErrorResponse{Message: "Validation failed: " + err.Error()})
		return
	}

	existingUsers, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: "Failed to load users: " + err.Error()})
		return
	}
	// Check for duplicate email
	for _, existingUser := range existingUsers {
		if existingUser.Email == userDto.Email {
			c.JSON(http.StatusConflict, genricerror.ErrorResponse{Message: "User with this email already exists"})
			return
		}
	}

	createdUser, err := services.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// UpdateUser
// @Summary Update an existing user
// @Description Replace an existing user's details
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Param  user body userdto.CreateUserRequest true "User details to update"
// @Success 200 {object} user.User
// @Failure 400 {object} genricerror.ErrorResponse
// @Failure 404 {object} genricerror.ErrorResponse
// @Failure 500 {object} genricerror.ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUserdto userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&updateUserdto); err != nil {
		c.JSON(http.StatusBadRequest, genricerror.ErrorResponse{Message: err.Error()})
		return
	}
	// Validate struct fields
	if err := validate.Struct(updateUserdto); err != nil {
		c.JSON(http.StatusBadRequest, genricerror.ErrorResponse{Message: "Validation failed: " + err.Error()})
		return
	}
	existingUsers, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: "Failed to load users: " + err.Error()})
		return
	}
	// Check for duplicate email
	for _, existingUser := range existingUsers {
		if existingUser.Email == updateUserdto.Email {
			c.JSON(http.StatusConflict, genricerror.ErrorResponse{Message: "User with this email already exists"})
			return
		}
	}

	updatedUser, err := services.UpdateUser(id, updateUserdto)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, genricerror.ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// PartialUpdateUser
// @Summary Partially update an existing user
// @Description Update only specific fields of an existing user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Param  user body userdto.CreateUserRequest true "User fields to update"
// @Success 200 {object} user.User
// @Failure 400 {object} genricerror.ErrorResponse
// @Failure 404 {object} genricerror.ErrorResponse
// @Failure 500 {object} genricerror.ErrorResponse
// @Router /users/{id} [patch]
func PartialUpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUserdto userdto.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserdto); err != nil {
		c.JSON(http.StatusBadRequest, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	updatedUser, err := services.UpdatePartialUser(id, updateUserdto)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, genricerror.ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser
// @Summary Delete an existing user
// @Description Delete a user by their ID
// @Tags Users
// @Param  id path int true "User ID"
// @Success 204 {object} user.User
// @Failure 404 {object} genricerror.ErrorResponse
// @Failure 500 {object} genricerror.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, genricerror.ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, genricerror.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
