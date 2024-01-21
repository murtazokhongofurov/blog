package v1

import (
	"strconv"

	"github.com/blog/internal/controller/http/models"
	"github.com/gin-gonic/gin"
)


// @ID 			 addpost
// @Router		/post [POST]
// @Summary		get post
// @Tags        Post
// @Description	Here post will be give data
// @Accept 		json
// @Produce		json
// @Param 		body body models.PostReq true "Post Data"
// @Success 	201 {object} models.Response{data=models.PostRes} "post data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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


// @ID 			 getpost
// @Router		/post [GET]
// @Summary		get post
// @Tags        Post
// @Description	Here post will be give data
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "PostInfo"
// @Success 	200 {object} models.Response{data=models.PostRes} "post data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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


// @ID 			 putpost
// @Router		/post [PUT]
// @Summary		PUT post
// @Tags        Post
// @Description	Here post will be updated
// @Accept 		json
// @Produce		json
// @Param 		body body models.PostRes true "Post Data"
// @Success 	200 {object} models.Response{data=models.Response} "post data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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


// @ID 			 deltepost
// @Router		/post [DELETE]
// @Summary		delete post
// @Tags        Post
// @Description	Here post will be deleted
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "PostInfo"
// @Success 	200 {object} models.Response{data=models.Response} "post data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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
