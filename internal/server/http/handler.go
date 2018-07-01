package http

import (
	"net/http"

	"github.com/fhbkoller/workshop-go/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService domain.UserService
}

func NewHandler(userService domain.UserService) http.Handler {
	handler := &handler{
		userService: userService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(handler.recovery())

	v1 := router.Group("/v1")
	v1.POST("/users", handler.postUser)
	v1.GET("/users/:id", handler.getUser)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
