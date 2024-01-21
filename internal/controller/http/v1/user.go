package v1

import (
	"strconv"

	"github.com/blog/internal/controller/http/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) UserRegister(c *gin.Context) {
	body := models.UserReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Blog().UserPost(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, res)
}

func (h *handlerV1) UserGet(c *gin.Context) {
	id := c.Query("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}

	res, err := h.Storage.Blog().UserGet(userId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, res)
}

func (h *handlerV1) UserUpdate(c *gin.Context) {
	body := models.UserRes{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().UserPut(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, gin.H{
		"ok":      true,
		"message": "updated successfully",
	})
}

func (h *handlerV1) UserDelete(c *gin.Context) {
	id := c.Query("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().UserDel(userId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, gin.H{
		"ok":      true,
		"message": "deleted successfully",
	})
}
