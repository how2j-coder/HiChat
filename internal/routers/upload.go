package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init()  {
	apiV1RouterFns = append(apiV1RouterFns, func(r *gin.RouterGroup) {
		uploadFileRouter(r, handler.NewUploadHandler())
	})
}

func uploadFileRouter(group *gin.RouterGroup, h handler.UploadHandler)  {
	g := group.Group("/upload")
	g.POST("/single", h.UploadSingleFile)
	g.POST("/multi", h.UploadMultiFile)
	g.GET("/file/:year/:mouth/:day/:key", h.GetFile)
}