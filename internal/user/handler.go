package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	service UserService
}

func NewUserHandler(service UserService) *userhandler {
	return &userhandler{
		service: service,
	}
}

func (h *userhandler) GetUser(c *gin.Context) {
	email := c.Query("email")
	id := c.Query("id")

	if email != "" {
		user, err := h.service.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"user": user})
		return
	}

	if id != "" {
		user, err := h.service.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "email or id required"})
}

func (h *userhandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}