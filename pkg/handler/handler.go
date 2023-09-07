package handler

import (
	"github.com/AhmAlgiz/marketplace/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		items := api.Group("/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/id/:id", h.getItemById)
			items.GET("/username/:username", h.getItemByUsername)
			items.GET("/title/:title", h.getItemByTitle)
			items.PUT("/", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}

		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserById)
			users.PUT("/update/", h.updateUser)
		}
	}

	return router
}
