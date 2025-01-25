package main

import (
	"log"
	"project-akhir-exam/controller"
	"project-akhir-exam/middleware"
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
	eventRepository := repository.NewEventRepository(db)

	//!! Service
	roleService := service.NewRoleService(roleRepository)
	userService := service.NewUserService(userRepository)
	eventService := service.NewEventService(eventRepository)
	authService := service.NewAuthService()

	//!! Middleware
	// authMiddleware := middleware.AuthMiddleware(authService, userService)
	authMiddlewareAdmin := middleware.AuthMiddlewareAdmin(authService, userService)

	//!! Controller
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService, authService)
	eventController := controller.NewEventController(eventService)

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
	api.GET("/events", eventController.GetAllEvents)
	api.GET("/events/:id", eventController.GetEvent)
	api.POST("/events", authMiddlewareAdmin, eventController.CreateEvent)
	api.PUT("/events/:id", authMiddlewareAdmin, eventController.UpdateEvent)
	api.DELETE("/events/:id", authMiddlewareAdmin, eventController.DeleteEvent)

	//!! Ticket Route
	// api.GET("/events/:id/tickets", authMiddlewareAdmin, ticketController.GetEventTickets)
	// api.GET("/tickets", authMiddleware, ticketController.GetUserTickets)
	// api.GET("/tickets/:id", authMiddleware, ticketController.GetTicket)
	// api.POST("/tickets", authMiddleware, ticketController.CreateTicket)
	// api.PUT("/tickets/:id/pay", authMiddleware, ticketController.MarkPaid)
	// api.PUT("/tickets/:id/cancel", authMiddleware, ticketController.MarkCancel)
	// api.PUT("/tickets/:id/status", authMiddlewareAdmin, ticketController.MarkStatus)

	router.Run()
}
