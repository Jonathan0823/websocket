package auth

import (
	"net/http"
	"websocket/internal/models"

	"github.com/gin-gonic/gin"
)

type authhandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *authhandler {
	return &authhandler{
		service: service,
	}
}

func (h *authhandler) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *authhandler) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login sucess"})
}
