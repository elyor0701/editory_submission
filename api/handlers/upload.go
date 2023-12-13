package handlers

import (
	"context"
	"editory_submission/api/http"
	"editory_submission/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"os"
	"strings"
)

type UploadResponse struct {
	Filename string `json:"filename"`
}

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type Path struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
}

// Upload godoc
// @ID upload
// @Security ApiKeyAuth
// @Router /upload [POST]
// @Summary Upload
// @Description Upload
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} http.Response{data=Path} "Path"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Upload(c *gin.Context) {
	var (
		file          File
		defaultBucket = "submission-editory"
		filePath      string
	)
	err := c.ShouldBind(&file)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	fName, _ := uuid.NewRandom()
	file.File.Filename = strings.ReplaceAll(file.File.Filename, " ", "")
	file.File.Filename = fmt.Sprintf("%s_%s", fName.String(), file.File.Filename)
	filePath = file.File.Filename
	dst, _ := os.Getwd()

	minioClient, err := minio.New(h.cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccessKey, ""),
		Secure: h.cfg.MinioProtocol,
	})

	h.log.Info("info", logger.String("MinioEndpoint: ", h.cfg.MinioEndpoint), logger.String("access_key: ",
		h.cfg.MinioAccessKeyID), logger.String("access_secret: ", h.cfg.MinioSecretAccessKey))

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = c.SaveUploadedFile(file.File, dst+"/"+file.File.Filename)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	splitContentType := strings.Split(file.File.Header["Content-Type"][0], "/")

	switch splitContentType[0] {
	case "image":
		filePath = "image/" + filePath
	case "video":
		filePath = "video/" + filePath
	default:
		filePath = "docs/" + filePath
	}

	_, err = minioClient.FPutObject(
		context.Background(),
		defaultBucket,
		filePath,
		dst+"/"+file.File.Filename,
		minio.PutObjectOptions{ContentType: file.File.Header["Content-Type"][0]},
	)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		err = os.Remove(dst + "/" + file.File.Filename)
		if err != nil {
			h.log.Error("cant remove file", logger.String("path", dst+"/"+file.File.Filename))
		}
		return
	}

	err = os.Remove(dst + "/" + file.File.Filename)
	if err != nil {
		h.log.Error("cant remove file")
	}

	h.handleResponse(c, http.Created, Path{
		Filename: filePath,
		Hash:     fName.String(),
	})
}
