package v1

import (
	"github.com/blog/config"
	_ "github.com/blog/internal/controller/http/docs" // docs
	v1 "github.com/blog/internal/controller/http/v1"
	"github.com/blog/internal/middleware"
	"github.com/blog/internal/storage"
	"github.com/blog/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Cfg     config.Config
	Storage storage.StorageI
	Log     logger.Logger
}

// @title           Blog API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(opt *Options) *gin.Engine {
	router := gin.New()
	basicAuth := middleware.BasicAuth()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	handlerV1 := v1.New(&v1.HandlerV1Option{
		Cfg:     &opt.Cfg,
		Storage: opt.Storage,
		Log:     opt.Log,
	})

	router.Use(basicAuth.Middleware)
	
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is running!!!",
		})
	})

	api := router.Group("/v1")

	user := api.Group("/user")
	user.POST("", handlerV1.UserRegister)
	user.GET("", handlerV1.UserGet)
	user.PUT("", handlerV1.UserUpdate)
	user.DELETE("", handlerV1.UserDelete)

	post := api.Group("/post")
	post.POST("", handlerV1.PostCreate)
	post.GET("", handlerV1.PostGet)
	post.PUT("", handlerV1.PostUpdate)
	post.DELETE("", handlerV1.PostDelete)

	comment := api.Group("/comment")
	comment.POST("", handlerV1.CommentCreate)
	comment.GET("", handlerV1.CommentGet)
	comment.PUT("", handlerV1.CommentUpdate)
	comment.DELETE("", handlerV1.CommentDelete)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
