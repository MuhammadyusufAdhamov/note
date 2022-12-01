package api

import (
	v1 "github.com/MuhammadyusufAdhamov/note/api/v1"
	"github.com/MuhammadyusufAdhamov/note/config"
	"github.com/MuhammadyusufAdhamov/note/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/MuhammadyusufAdhamov/note/api/docs"
)

type RouterOptions struct {
	Cfg *config.Config
	Storage storage.StorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      localhost:8000
// @BasePath  /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg: opt.Cfg,
		Storage: opt.Storage,
	})

	router.Static("/media", "./media")

	apiV1 := router.Group("/v1")

	apiV1.GET("/user/:id", handlerV1.GetNote)
	apiV1.GET("/user", handlerV1.GetAllUsers)
	apiV1.POST("/user", handlerV1.CreateUser)
	apiV1.PUT("/user/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/user/:id", handlerV1.DeleteUser)

	apiV1.GET("/note/:id", handlerV1.GetUser)
	apiV1.GET("/note", handlerV1.GetAllnotes)
	apiV1.POST("/note", handlerV1.CreateNote)
	apiV1.PUT("/note/:id", handlerV1.UpdateNote)
	apiV1.DELETE("/note/:id", handlerV1.DeleteNote)

	apiV1.POST("/file-upload", handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}