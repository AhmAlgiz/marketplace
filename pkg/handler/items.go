package handler

import (
	"net/http"
	"strconv"

	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/gin-gonic/gin"
)

type getAllItemsResponse struct {
	Data []structures.Item `json:"data"`
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input structures.Item
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userId

	id, err := h.services.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id parameter")
	}

	sl, err := h.services.GetItemById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: sl,
	})
}

func (h *Handler) getItemByTitle(c *gin.Context) {

}

func (h *Handler) getItemByUsername(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
