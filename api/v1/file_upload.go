package v1

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type File struct {
	File *multipart.FileHeader  `form:"file" binding:"required"`
}

// @Router /file-upload [post]
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Success 201 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UploadFile(c *gin.Context) {
	var file File
	fmt.Println("---------------------------------------1")
	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("---------------------------------------2")

	id := uuid.New()
	filename := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()
	fmt.Println("---------------------------------------3")

	if _, err := os.Stat(dst + "/media"); os.IsNotExist(err) {
		os.Mkdir(dst + "/media", os.ModePerm)
	}

	filePath := "/media/" + filename
	err = c.SaveUploadedFile(file.File, dst + filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("---------------------------------------4")

	c.JSON(http.StatusCreated, gin.H{
		"filename": filePath,
	})
}