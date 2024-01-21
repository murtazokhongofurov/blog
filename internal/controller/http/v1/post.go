package v1

import (
	"strconv"

	"github.com/blog/internal/controller/http/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) PostCreate(c *gin.Context) {
	body := models.PostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Blog().PostCreate(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, res)
}

func (h *handlerV1) PostGet(c *gin.Context) {
	id := c.Query("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}

	res, err := h.Storage.Blog().PostGet(postId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, res)
}

func (h *handlerV1) PostUpdate(c *gin.Context) {
	body := models.PostRes{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().PostPut(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, gin.H{
		"ok":      true,
		"message": "updated successfully",
	})
}

func (h *handlerV1) PostDelete(c *gin.Context) {
	id := c.Query("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().PostDel(postId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, gin.H{
		"ok":      true,
		"message": "deleted successfully",
	})
}
