package service

import (
	"HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"path"
)

// FileUpload 文件上传
func FileUpload(ctx *gin.Context)  {
	//创建文件存储文件目录
	if _, err := os.Stat("./upload_file"); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("./upload_file", os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(http.StatusText(http.StatusInternalServerError)))
				global.Logger.Error(err.Error())
				return
			}
		}
	}

	//判断是上传类型为多个/单个
	uploadType := ctx.DefaultQuery("upload_type", "single")

	//单文件上传
	if uploadType == "single" {
		uploadFile, err := ctx.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
		}
		fileType := path.Ext(uploadFile.Filename)
		filePath := common.Md5encoder(uploadFile.Filename)
		file := models.File{
			FileName: uploadFile.Filename,
			FileSize: uploadFile.Size,
			FileType: fileType,
			FilePath: filePath,
		}
		// 使用md5文件名保存
		uploadFile.Filename = filePath
		err = ctx.SaveUploadedFile(uploadFile, "./upload_file/"+filePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
			return
		}
		// 存储到数据库
		_, err = dao.UploadSingle(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, common.Success.WithData(file))
	}

	//多文件上传
	if uploadType == "multiple" {
		ctx.JSON(http.StatusOK,common.Success.WithMsg("upload file"))

	}

}

// FileDownload 文件下载
func FileDownload(ctx *gin.Context)  {
	fileName := ctx.Param("fileName")
	file, err := dao.FindFileFileName(fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}
	if file == nil {
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg("file not found"))
		return
	}

	// Content-Type 设置为 application/octet-stream，告诉浏览器 一个二进制文件流。
	ctx.Header("Content-Type", "application/octet-stream")
	// Content-Disposition 设置为 attachment; filename="文件名"，下载文件的默认文件名。
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	//确保了文件在传输过程中不会因编码问题而损坏
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File("./upload_file/"+file.FilePath)
}