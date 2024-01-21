package v1

import (
	"strconv"

	"github.com/blog/internal/controller/http/models"
	"github.com/gin-gonic/gin"
)

// @ID 			 postcomment
// @Router		/comment [POST]
// @Summary		post comment
// @Tags        Comment
// @Description	Here comment will be created
// @Accept 		json
// @Produce		json
// @Param 		body body models.CommentReq true "CommentInfo"
// @Success 	201 {object} models.Response{data=models.CommentRes} "comment data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
func (h *handlerV1) CommentCreate(c *gin.Context) {
	body := models.CommentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Blog().CommentPost(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.Created, res)
}

// @ID 			 getcomment
// @Router		/comment [GET]
// @Summary		get comment
// @Tags        Comment
// @Description	Here comment will be give data
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "CommentInfo"
// @Success 	200 {object} models.Response{data=models.CommentRes} "comment data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
func (h *handlerV1) CommentGet(c *gin.Context) {
	id := c.Query("id")
	comId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}

	res, err := h.Storage.Blog().CommentGet(comId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.OK, res)
}

// @ID 			 udpatecomment
// @Router		/comment [PUT]
// @Summary		update comment
// @Tags        Comment
// @Description	Here comment will be update
// @Accept 		json
// @Produce		json
// @Param 		body body models.CommentRes true "CommentData"
// @Success 	200 {object} models.Response{data=models.Response} "comment data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
func (h *handlerV1) CommentUpdate(c *gin.Context) {
	body := models.CommentRes{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().CommentPut(&body)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.OK, gin.H{
		"ok":      true,
		"message": "updated successfully",
	})
}

// @ID 			 deltecomment
// @Router		/comment [DELETE]
// @Summary		delete comment
// @Tags        Comment
// @Description	Here comment will be deleted
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "CommentInfo"
// @Success 	200 {object} models.Response{data=models.Response} "comment data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
func (h *handlerV1) CommentDelete(c *gin.Context) {
	id := c.Query("id")
	comId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}
	err = h.Storage.Blog().CommentDel(comId)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, models.OK, gin.H{
		"ok":      true,
		"message": "deleted successfully",
	})
}
