package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 파일업로드 test
func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	// file read
	path := "/mnt/c/Users/Annotation AI-I/Downloads/task_12월 2주차 상품데이타 폴리건-41-2020_12_08_03_52_54-cvat for images 1.1/images/_[회전]IMG_4481.JPG"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	// make form file
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	w, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	// file copy
	io.Copy(w, file)
	writer.Close()

	// 실제 url call
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	// form data라는걸 명시 해줘야 함
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// 업로드한 파일과 업로드된 파일 비교
	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	// uploadfile과 originfile Read
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)

	// 비교
	assert.Equal(originData, uploadData)
}
