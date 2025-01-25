package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"project-akhir-exam/helper"
	"project-akhir-exam/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserController(userService service.UserService, authService service.AuthService) *userController {
	return &userController{userService, authService}
}

func (h *userController) Register(c *gin.Context) {
	var input service.RegisterInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input.Email)
	if err != nil {
		errorMessage := gin.H{"errors": "Email checking failed"}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isEmailAvailable {
		errorMessage := gin.H{"errors": "Email already been used"}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("public/user/%d-%s", rand.Int(), file.Filename)
	// path := fmt.Sprintf("public/product/%d-%s", input.ProductID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.Register(input, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.RoleID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := helper.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input service.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID, loginUser.RoleID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := helper.FormatUser(loginUser, token)

	response := helper.APIResponse("Successfully Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// func (h *userController) UploadPhoto(c *gin.Context) {
// 	file, err := c.FormFile("photo")
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}

// 		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser").(entity.User)
// 	userID := currentUser.ID

// 	path := fmt.Sprintf("public/user/%d-%s", userID, file.Filename)

// 	err = c.SaveUploadedFile(file, path)
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}

// 		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	_, err = h.userService.UploadPhoto(userID, path)
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}

// 		response := helper.APIResponse("Failed to upload photo", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	data := gin.H{"is_uploaded": true}

// 	response := helper.APIResponse("Photo successfully uploaded", http.StatusOK, "success", data)
// 	c.JSON(http.StatusOK, response)
// }
