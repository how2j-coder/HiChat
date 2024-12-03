package service

import (
	"HiChat/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileUpload(ctx *gin.Context)  {
	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["file"]
		fmt.Println(len(files))
		for _, file := range files {
			fmt.Println("Uploaded file:", file.Filename)
		}
	} else {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK,common.Success.WithMsg("upload file"))
}