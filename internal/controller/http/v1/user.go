package v1

import (
	"strconv"

	"github.com/blog/internal/controller/http/models"
	"github.com/gin-gonic/gin"
)

// Register User
// @ID 			 adduser
// @Router		/user [POST]
// @Summary		add user
// @Tags        User
// @Description	Here user will be registred
// @Accept 		json
// @Produce		json
// @Param 		body body models.UserReq true "User"
// @Success 	201 {object} models.Response{data=models.UserRes} "User data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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

// @ID 			 getuser
// @Router		/user [GET]
// @Summary		get user
// @Tags        User
// @Description	Here user will be give data
// @Accept 		json
// @Produce		json
// @Param 		id  query int true "UserInfo"
// @Success 	200 {object} models.Response{data=models.UserRes} "User data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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
	h.handleResponse(c, models.OK, res)
}

// @ID 			 putuser
// @Router		/user [PUT]
// @Summary		PUT user
// @Tags        User
// @Description	Here user will be updated
// @Accept 		json
// @Produce		json
// @Param 		body body models.UserRes true "User"
// @Success 	200 {object} models.Response{data=models.Response} "User data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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
	h.handleResponse(c, models.OK, gin.H{
		"ok":      true,
		"message": "updated successfully",
	})
}

// @ID 			 deleteuser
// @Router		/user [DELETE]
// @Summary		delete user
// @Tags        User
// @Description	Here user will be deleted
// @Accept 		json
// @Produce		json
// @Param 		id query int true "Delete User"
// @Success 	200 {object} models.Response{data=models.Response} "data"
// @Response 	400 {object} models.Response{data=string} "Bad Request"
// @Failure 	500 {object} models.Response{data=string} "Server Error"
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
	h.handleResponse(c, models.OK, gin.H{
		"ok":      true,
		"message": "deleted successfully",
	})
}
