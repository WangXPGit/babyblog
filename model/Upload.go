package model

import (
	"babyblog/utils"
	"babyblog/utils/errmsg"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var AccessKey = utils.AccessKey
var ScretKey = utils.ScretKey
var Bucket = utils.Bucket
var AliyunServer = utils.AliyunServer

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "<yourObjectName>"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "<yourLocalFileName>"
	// 创建OSSClient实例。
	client, err := oss.New(AliyunServer, AccessKey, ScretKey)
	var url string
	if err != nil {
		return url, errmsg.ERROR
	}
	// 获取存储空间。
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		return url, errmsg.ERROR
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return url, errmsg.ERROR
	}
}
