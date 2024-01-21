package v1

import (
	"github.com/blog/config"
	"github.com/blog/internal/controller/http/models"
	"github.com/blog/internal/storage"
	"github.com/blog/pkg/logger"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	Cfg     *config.Config
	Storage storage.StorageI
	Log     logger.Logger
}

type HandlerV1Option struct {
	Cfg     *config.Config
	Storage storage.StorageI
	Log     logger.Logger
}

func New(optoins *HandlerV1Option) *handlerV1 {
	return &handlerV1{
		Cfg:     optoins.Cfg,
		Storage: optoins.Storage,
		Log:     optoins.Log,
	}
}

func (h *handlerV1) handleResponse(c *gin.Context, status models.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, models.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
