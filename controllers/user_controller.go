package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/app/request"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/app/response"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/helpers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/models"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/services"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
	authService helpers.JwtService
}

func (h *UserControllerImpl) Register(c *gin.Context) {
	var input request.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse("Input invalid", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, _ := h.userService.IsEmailAvailable(input.Email)

	if !isEmailAvailable {
		response := helpers.APIResponse("Email has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isUsernameAvailable, _ := h.userService.IsUsernameAvailable(input.Username)

	if !isUsernameAvailable {
		response := helpers.APIResponse("Username has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helpers.APIResponse("Generate token is failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := response.RegisterFormatUser(newUser, token)

	response := helpers.APIResponse("Account successfully registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Login(c *gin.Context) {
	var input request.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helpers.APIResponse("Generate token is failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := response.RegisterFormatUser(loggedinUser, token)

	response := helpers.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Update(c *gin.Context) {
	var inputData request.UpdateInput
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response := helpers.APIResponse("Invalid get user ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(models.User)

	inputData.User = currentUser

	if userID != currentUser.ID {
		response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorMessage := gin.H{"errors": "errors"}
		response := helpers.APIResponse("Input invalid", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	checkUsername, _ := h.userService.IsUsernameAvailable(inputData.Username)
	if !checkUsername && currentUser.Username != inputData.Username {
		response := helpers.APIResponse("username has been  used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	checkEmail, _ := h.userService.IsEmailAvailable(inputData.Email)
	if !checkEmail && currentUser.Email != inputData.Email {
		response := helpers.APIResponse("email has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUser(userID, inputData)
	if err != nil {
		response := helpers.APIResponse("error on update user", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("Successfully updated user", http.StatusOK, "Sukses", response.GetFormatUser(updatedUser))
	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	currentUser := c.MustGet("currentUser").(models.User)

	userDetail, err := h.userService.GetUserByID(userId)
	if err != nil {
		response := helpers.APIResponse("Failed to get detail user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userDetail.ID != currentUser.ID {
		response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = h.userService.DeleteUser(currentUser.ID)
	if err != nil {
		response := helpers.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Successfully deleted user", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func NewUserController(userService services.UserService, authService helpers.JwtService) UserController {
	return &UserControllerImpl{userService: userService, authService: authService}
}
