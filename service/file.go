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
	"strconv"
	"sync"
	"time"
)

type FileRes struct{
	FileName string `json:"file_name"`;
	FilePath string `json:"file_path"`;
	FileId string `json:"file_id"`
}

// FileUploadSingle 文件上传
func FileUploadSingle(ctx *gin.Context)  {
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

	//单文件上传
	uploadFile, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}
	fileType := path.Ext(uploadFile.Filename)
	uploadTime := strconv.FormatInt(time.Now().Unix(), 10)
	filePath := common.Md5encoder(uploadFile.Filename + "-" + uploadTime)
	userId, _ := ctx.Get(global.AuthCtxFiled)
	file := models.File{
		FileName: uploadFile.Filename,
		FileSize: uploadFile.Size,
		FileType: fileType,
		FilePath: filePath,
		UserID: userId.(string),
	}
	// 使用md5文件名保存
	uploadFile.Filename = filePath
	err = ctx.SaveUploadedFile(uploadFile, "./upload_file/"+filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}
	// 存储到数据库
	singleFile, err := dao.UploadSingle(file)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.Success.WithData(FileRes{
		singleFile.FileName,
		singleFile.FilePath,
		singleFile.ID,
	}))
}

// FileUploadMultiple 多文件上传
func FileUploadMultiple(ctx *gin.Context)  {
	ctx.Request.Header.Set("Authorization", "")
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err, form)
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}
	files:= form.File["file"]
	if len(files) == 0 {
		ctx.JSON(http.StatusOK, common.ParamsNilError.WithMsg("文件不能为空"))
		return
	}
	// 文件数据处理
	var saveFiles []models.File
	var wg sync.WaitGroup
	cCtx := ctx.Copy()
	userId, _ := ctx.Get(global.AuthCtxFiled)
	for i, f := range files {
		wg.Add(1)
		go func(i int) {
			uploadTime := strconv.FormatInt(time.Now().Unix(), 10)
			fileType := path.Ext(f.Filename)
			filePath := common.Md5encoder(f.Filename + "-" + uploadTime + "-" + strconv.Itoa(i))
			_ = cCtx.SaveUploadedFile(files[i], "./upload_file/"+filePath)
			saveFiles = append(saveFiles, models.File{
				FileName: f.Filename,
				FileSize: f.Size,
				FileType: fileType,
				FilePath: filePath,
				UserID: userId.(string),
			})
			wg.Done()
		}(i)
	}
	wg.Wait()

	//存储到数据库
	multipleFile, err := dao.UploadMultiple(saveFiles)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error.WithMsg(err.Error()))
		return
	}

	var sendData []FileRes
	for _, file := range *multipleFile {
		sendData = append(sendData, FileRes{
			FileId: file.ID,
			FileName: file.FileName,
			FilePath: file.FilePath,
		})
	}

	ctx.JSON(http.StatusOK,common.Success.WithData(multipleFile))
}

// FileDownload 文件下载
func FileDownload(ctx *gin.Context)  {
	httpContentType := map[string]string{
		".avi": "video/avi",
		".mp3": "audio/mp3",
		".mp4": "video/mp4",
		".wmv": "video/x-ms-wmv",
		".asf":  "video/x-ms-asf",
		".rm":   "application/vnd.rn-realmedia",
		".rmvb": "application/vnd.rn-realmedia-vbr",
		".mov":  "video/quicktime",
		".m4v":  "video/mp4",
		".flv":  "video/x-flv",
		".jpg":  "image/jpeg",
		".png":  "image/png",
		".txt":  "text/plain; charset=utf-8",
	}
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
	contentType, exist := httpContentType[file.FileType]
	if !exist {
		contentType = "application/octet-stream"
	}
	ctx.Header("Content-Type", contentType)
	// Content-Disposition 设置为 attachment; filename="文件名"，下载文件的默认文件名。
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	//确保了文件在传输过程中不会因编码问题而损坏
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File("./upload_file/"+file.FilePath)
}