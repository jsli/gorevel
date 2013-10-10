package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/robfig/revel"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	uploadPath string
	imageExts  string = ".jpg.jpeg.png"
)

func saveFile(r *revel.Request, formField string) string {
	file, header, err := r.FormFile(formField)
	if err != nil {
		return ""
	}
	defer file.Close()

	uuid := strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	ext := filepath.Ext(header.Filename)
	fileName := uuid + ext

	os.MkdirAll(uploadPath, 0777)

	f, err := os.OpenFile(uploadPath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(f, file)
	}

	return fileName
}

func deleteFile(fileName string) error {
	err := os.Remove(uploadPath + fileName)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func getFileExt(r *revel.Request, formField string) (bool, string) {
	file, header, err := r.FormFile(formField)
	if err != nil {
		return false, ""
	}
	defer file.Close()

	return true, strings.ToLower(filepath.Ext(header.Filename))
}

// (c Controller, fileExts 允许的文件类型, formField 表单字段, message 表单字段提示信息)
func checkFileExt(c *revel.Controller, fileExts, formField, message string) {
	if ok, ext := getFileExt(c.Request, formField); ok && !strings.Contains(fileExts, ext) {
		err := &revel.ValidationError{
			Message: message,
			Key:     formField,
		}
		c.Validation.Errors = append(c.Validation.Errors, err)
	}
}
