package handler

import (
	"net/http"
	"strconv"

	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/gin-gonic/gin"
)

type getUsersResponse struct {
	Data []structures.GetUser `json:"data"`
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input structures.UpdateUser
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateUser(input, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusRespone{
		Status: true,
	})
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id parameter")
		return
	}

	sl, err := h.services.GetUserById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUsersResponse{
		Data: sl,
	})
}
