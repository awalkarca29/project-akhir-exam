package main

import (
	"log"
	"project-akhir-exam/controller"
	"project-akhir-exam/repository"
	"project-akhir-exam/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//!! Database
	dsn := "root:@tcp(127.0.0.1:3306)/project_akhir_exam?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//!! Repository
	roleRepository := repository.NewRoleRepository(db)
	userRepository := repository.NewUserRepository(db)

	//!! Service
	roleService := service.NewRoleService(roleRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()

	//!! Middleware
	// authMiddleware := middleware.AuthMiddleware(authService, userService)
	// authMiddlewareAdmin := middleware.AuthMiddlewareAdmin(authService, userService)

	//!! Controller
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	router.Static("/photo", "./public/user")
	router.Static("/image", "./public/event")
	api := router.Group("/api/v1")

	//!! Role Route
	api.POST("/role", roleController.CreateRole) //!! DEVELOP ONLY

	//!! User Route
	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)

	//!! Event Route

	//!! Ticket Route

	router.Run()
}
