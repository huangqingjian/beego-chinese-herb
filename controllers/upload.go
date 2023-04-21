package controllers

import (
	"beego-chinese-herb/component"
	"beego-chinese-herb/models"
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
	"strings"
)

// 上传api
type UploadController struct {
	BaseController
}

// @Title 图片上传
// @Description 图片上传
// @Param file formData	[]byte true "upload img"
// @Success 200 {object} models.Response
// @Failure 200 error response object
// @router /img [post]
func (u *UploadController) UploadImg() {
	f, h, err := u.GetFile("file")
	if err != nil {
		log.Error("getfile err[%v] ", err)
	}
	defer f.Close()
	// 文件存储路径
	picturePath, _ := web.AppConfig.String("imagepath")
	fileName := h.Filename
	fileName = component.GetUuid() + fileName[strings.LastIndex(fileName, "."):]
	filePath := picturePath + "/" + fileName
	u.SaveToFile("file", filePath)
	// 文件访问路径
	visitPath := "/picture/" + fileName
	u.Data["json"] = models.ResponseSuccess(visitPath)
	u.ServeJSON()
}