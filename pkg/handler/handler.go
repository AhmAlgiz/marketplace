package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}

	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}

		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
		}
	}

	return router
}
