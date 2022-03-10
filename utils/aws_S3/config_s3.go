package awss3

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
	"github.com/lithammer/shortuuid"
)

func InitS3(key, secret, region string) *session.Session {
	connect, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				"AKIAVOMUO3KKNSP4RXWR",
				"o3T3ozzKzrdIfiDTPMVFMgP7NWfpFm75hxtX2Cww",
				"",
			),
		},
	)
	if err != nil {
		log.Error("aws S3 Config error", err)
	}

	return connect
}

func Upload(sess *session.Session, fileHeader multipart.FileHeader) string {

	var uid = shortuuid.New()

	var manager = s3manager.NewUploader(sess)
	var src, err = fileHeader.Open()
	if err != nil {
		log.Info(err)
	}
	defer src.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	var res, err1 = manager.Upload(
		&s3manager.UploadInput{
			Bucket:       aws.String("airbnb-app"),
			Key:          aws.String(uid),
			ACL:          aws.String("public-read-write"),
			Body:         bytes.NewReader(buffer),
			ContentType:  aws.String(http.DetectContentType(buffer)),
			StorageClass: aws.String("STANDARD"),
		},
	)

	if err1 != nil {
		log.Info(res)
		log.Error(err1)
	}

	var url = "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/" + uid

	return url
}

// func Update(ses *session.Session, name string, fileHeader multipart.FileHeader) string {

// 	var src, err0 = fileHeader.Open()
// 	if err0 != nil {
// 		log.Info(err0)
// 	}
// 	defer src.Close()

// 	size := fileHeader.Size
// 	buffer := make([]byte, size)
// 	src.Read(buffer)

// 	var svc = s3.New(ses)

// 	var input = &s3.PutObjectInput{
// 		Bucket:       aws.String("airbnb-app"),
// 		Key:          aws.String(name),
// 		ACL:          aws.String("public-read-write"),
// 		Body:         bytes.NewReader(buffer),
// 		ContentType:  aws.String(http.DetectContentType(buffer)),
// 		StorageClass: aws.String("STANDARD"),
// 	}

// 	var res, err = svc.PutObject(input)

// 	if err != nil {
// 		log.Info(res)
// 		if errA, ok := err.(awserr.Error); ok {
// 			switch errA.Code() {
// 			default:
// 				log.Info(errA.Error())
// 			}
// 		} else {
// 			log.Info(err.Error())
// 		}
// 	}

// 	return ""
// }
