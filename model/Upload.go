package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var Secretkey = utils.Secretkey
var Bucket = utils.Bucket
var ImgUrl = utils.QiNiuServer

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {

	// return url, errmsg.SUCCESS
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, Secretkey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
