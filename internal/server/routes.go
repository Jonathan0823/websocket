package server

import (
	"net/http"
	"websocket/internal/auth"
	"websocket/internal/middleware"
	"websocket/internal/user"
	"websocket/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	hub := ws.NewHub()
	wsHandler := ws.NewWsHandler(hub)
	go hub.Run()

	authRepo := auth.NewAuthRepository(s.db.GetDB())
	userRepo := user.NewUserRepository(s.db.GetDB())

	authService := auth.NewAuthService(authRepo)
	userService := user.NewUserService(userRepo)

	authHandler := auth.NewAuthHandler(authService)
	userHandler := user.NewUserHandler(userService)

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	userGroup := r.Group("/user").Use(middleware.JWTMiddleware(authService))
	{
		userGroup.GET("/", userHandler.GetUser)
		userGroup.GET("/all", userHandler.GetUsers)
	}

	wsGroup := r.Group("/ws")
	{
		wsGroup.POST("/room", wsHandler.CreateRoom)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
