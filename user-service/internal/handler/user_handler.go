package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/service"
)

// UserHandler handles HTTP requests for user-related features
type UserHandler struct {
	svc service.UserService
}

// NewUserHandler returns a new instance of UserHandler
func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// RegisterRoutes maps the user-related endpoints
func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	// Public routes
	router.POST("/register", h.Register)
	
	// You can add more routes here like /login, /profile
}

// HealthCheck provides a simple status check for the service
func (h *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "User Service is up and running",
	})
}

// Register handles the user registration request
func (h *UserHandler) Register(c *gin.Context) {
	var input db_users_gen.CreateUserParams

	// 1. Bind JSON input to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Warn("invalid registration input", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	// 2. Call service layer
	// Note: We'll need to update the service layer to handle the params properly
	user, err := h.svc.RegisterUser(c.Request.Context(), input)
	if err != nil {
		slog.Error("failed to register user", "error", err, "email", input.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	// 3. Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    user,
	})
}
