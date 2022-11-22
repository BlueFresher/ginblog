package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

const UPLOAD_PATH string = "/home/hyq/Desktop/app/ginblog/asset/"

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	// image, err := os.Create(UPLOAD_PATH + fileHeader.Filename)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer image.Close()
	// _, err = io.Copy(image, file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// url := UPLOAD_PATH + fileHeader.Filename
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
