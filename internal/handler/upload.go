package handler

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/model"
	"com/chat/service/internal/types"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/logger"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	ErrBodyTooLarge = errors.New("http: request body too large")
)

type UploadHandler interface {
	UploadSingleFile(c *gin.Context)
	UploadMultiFile(c *gin.Context)
	GetFile(c *gin.Context)
}

type uploadHandler struct {
	iDao dao.UploadDao
}

var _ UploadHandler = (*uploadHandler)(nil)

func NewUploadHandler() UploadHandler {
	return &uploadHandler{
		iDao: dao.NewUploadDao(database.GetDB()),
	}
}

// UploadSingleFile params: type file
func (h *uploadHandler) UploadSingleFile(c *gin.Context) {
	// 设置最大内存限制为8MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 160<<20)
	// 获取上传的文件
	file, err := c.FormFile("file")

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			logger.Warn("UploadSingleFile: file is missing")
			response.Error(c, ecode.ErrCreateUpload)
			return
		}
		if strings.Contains(err.Error(), "request body too large") {
			logger.Warn("UploadSingleFile: request body too large")
			response.Error(c, ecode.ErrCreateUpload.RewriteMsg("file too large, just 8m"))
			return
		}
		response.Error(c, ecode.ErrCreateUpload)
		return
	}

	fileType, err := detectFileType(file)       // 文件类型参数
	mimeType := file.Header.Get("Content-Type") // 从文件头获取MIME类型
	size := file.Size                           // 文件大小(字节)

	fileKey := generateUUID()
	filename := filepath.Base(file.Filename)

	dstPath := "files/" + fileKey + "." + fileType
	dstFile, err := os.Create(dstPath)
	if err != nil {
		logger.Warn("UploadSingleFile: save uploaded file err")
		response.Error(c, ecode.InternalServerError)
		return
	}
	defer dstFile.Close()

	// 获取MD5
	md5Value, _ := h.GetMD5Key(file, dstFile)

	uploadFile := &model.Upload{
		Type:             fileType,
		MimeType:         mimeType,
		Size:             size,
		MD5:              md5Value,
		Key:              fileKey,
		Storage:          "local",
		OriginalFileName: filename,
		RelativeFile:     dstPath,
		URL:              "/file/" + time.Now().Format("2006/01/02") + "/" + fileKey,
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.CreateUpload(ctx, uploadFile)
	if err != nil {
		logger.Warn("UploadSingleFile: save uploaded file err")
		response.Error(c, ecode.InternalServerError)
		return
	}
	response.Success(c, uploadFile)
}

func (h *uploadHandler) UploadMultiFile(c *gin.Context) {
	// 获取 multipart form
	form, err := c.MultipartForm()
	if err != nil {
		logger.Warn("UploadMultiFile: multipart form err")
		response.Error(c, ecode.InternalServerError)
		return
	}

	// 获取所有文件
	files := form.File["files"]
	if len(files) == 0 {
		logger.Warn("UploadMultiFile: no files to upload")
		response.Error(c, ecode.InternalServerError)
		return
	}

	// 响应数据
	results := make([]types.CreatUploadMultiReq, 0, len(files))

	// 并发处理每个文件
	var wg sync.WaitGroup
	// 限制最大并发数
	concurrency := 5
	if len(files) <= 5 {
		concurrency = len(files)
	}
	resultChan := make(chan types.CreatUploadMultiReq, concurrency)

	// 并发保存文件
	for _, fileHeaders := range files {
		wg.Add(1)
		go func(fh *multipart.FileHeader) {
			defer wg.Done()

			fileType, err := detectFileType(fileHeaders)       // 文件类型参数
			mimeType := fileHeaders.Header.Get("Content-Type") // 从文件头获取MIME类型
			size := fileHeaders.Size                           // 文件大小(字节)
			filename := filepath.Base(fileHeaders.Filename)

			result := types.CreatUploadMultiReq{
				Type:             fileType,
				MimeType:         mimeType,
				Size:             size,
				Storage:          "local",
				OriginalFileName: filename,
			}
			// 打开文件流
			srcFile, err := fh.Open()
			if err != nil {
				result.Error = err.Error()
				resultChan <- result
				return
			}
			defer srcFile.Close()

			// 创建文件路径
			fileKey := generateUUID()

			dstPath := "files/" + fileKey +  fileType
			result.RelativeFile = dstPath

			// 创建目标文件
			dstFile, err := os.Create(dstPath)
			if err != nil {
				result.Error = err.Error()
				resultChan <- result
				return
			}
			defer dstFile.Close()

			// 流式处理（带MD5计算）
			hash := md5.New()
			buf := make([]byte, 20<<20) // 20MB缓冲区
			totalRead := int64(0)

			for {
				n, err := srcFile.Read(buf)
				if err != nil && err != io.EOF {
					result.Error = err.Error()
					resultChan <- result
					return
				}
				if n == 0 {
					break
				}

				// 写入分块数据
				if _, err := dstFile.Write(buf[:n]); err != nil {
					result.Error = err.Error()
					resultChan <- result
					return
				}

				// 更新MD5
				hash.Write(buf[:n])
				totalRead += int64(n)
			}

			// 验证完整性
			if totalRead != fh.Size {
				result.Error = "incomplete transfer"
				_ = os.Remove(dstPath)
				resultChan <- result
				return
			}

			result.RelativeFile = filename
			result.Key = fileKey
			result.MD5 = hex.EncodeToString(hash.Sum(nil))
			result.URL = "/file/" + time.Now().Format("2006/01/02") + "/" + fileKey

			resultChan <- result
		}(fileHeaders)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	for res := range resultChan {
		fmt.Println(res.Error, "99999")
		results = append(results, res)
	}

	response.Success(c, results)
}

func (h *uploadHandler) GetFile(c *gin.Context) {
	filePath := c.Param("year")
	response.Success(c, filePath)
}

func (h *uploadHandler) GetMD5Key(uploadFile *multipart.FileHeader, saveFile *os.File) (string, error) {
	size := uploadFile.Size

	// 打开上传文件流
	srcFile, err := uploadFile.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	// 同时写入文件和计算MD5
	hash := md5.New()

	if size > 100<<20 {
		// 分块处理（4MB/块）
		buf := make([]byte, 4<<20) // 4MB缓冲区
		for {
			n, err := srcFile.Read(buf)
			if err != nil && err != io.EOF {
				return "", err
			}
			if n == 0 {
				break
			}

			// 同时写入和计算MD5
			if _, err := saveFile.Write(buf[:n]); err != nil {
				return "", err
			}
			hash.Write(buf[:n])
		}

	} else {
		// 只需一次文件读写，内存效率高 中小文件上传（<100MB）
		multiWriter := io.MultiWriter(saveFile, hash)

		// 复制文件内容
		if _, err = io.Copy(multiWriter, srcFile); err != nil {
			return "", err
		}
	}
	key := hex.EncodeToString(hash.Sum(nil))
	return key, nil
}

func detectFileType(fileHeader *multipart.FileHeader) (fileType string, miniType error) {
	// 打开上传的文件
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close() // 确保关闭文件

	// 检测文件类型
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}

	return mime.Extension(), nil
}

func generateUUID() string {
	// 生成带连字符的UUID
	uuidWithHyphens := uuid.NewV4().String()
	// 移除所有连字符
	return strings.Replace(uuidWithHyphens, "-", "", -1)
}