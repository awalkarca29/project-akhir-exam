package main

import (
	"project-akhir-exam/config"
	"project-akhir-exam/controller"
	"project-akhir-exam/middleware"
	"project-akhir-exam/repository"
	"project-akhir-exam/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//!! Database
	db := config.InitDB()

	//!! Repository
	roleRepository := repository.NewRoleRepository(db)
	userRepository := repository.NewUserRepository(db)
	eventRepository := repository.NewEventRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	reportRepo := repository.NewReportRepository(db)

	//!! Service
	roleService := service.NewRoleService(roleRepository)
	userService := service.NewUserService(userRepository)
	eventService := service.NewEventService(eventRepository)
	ticketService := service.NewTicketService(ticketRepository)
	reportService := service.NewReportService(reportRepo)
	authService := service.NewAuthService()

	//!! Middleware
	authMiddleware := middleware.AuthMiddleware(authService, userService)
	authMiddlewareAdmin := middleware.AuthMiddlewareAdmin(authService, userService)

	//!! Controller
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService, authService)
	eventController := controller.NewEventController(eventService)
	ticketController := controller.NewTicketController(ticketService)
	reportController := controller.NewReportController(reportService)

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
	api.GET("/events/:id/tickets", authMiddlewareAdmin, ticketController.GetEventTickets)
	api.GET("/tickets", authMiddleware, ticketController.GetUserTickets)
	api.GET("/tickets/:id", authMiddleware, ticketController.GetTicket)
	api.POST("/tickets", authMiddleware, ticketController.CreateTicket)
	api.PATCH("/tickets/:id/pay", authMiddleware, ticketController.MarkPaid)
	api.PATCH("/tickets/:id/cancel", authMiddleware, ticketController.MarkCancel)
	// api.PATCH("/tickets/:id/status", authMiddlewareAdmin, ticketController.MarkStatus)

	//!! Report Route
	api.GET("/reports/summary", authMiddlewareAdmin, reportController.SummaryHandler)
	api.GET("/reports/event", authMiddlewareAdmin, reportController.EventHandler)

	router.Run()
}
