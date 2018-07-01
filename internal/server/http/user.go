package http

import (
	"net/http"

	"github.com/fhbkoller/workshop-go/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) postUser(c *gin.Context) {
	var user *domain.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	if err := h.userService.Create(user); err != nil {
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

func (h *handler) getUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.userService.Retrieve(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
