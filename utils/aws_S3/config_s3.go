package awss3

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
	"github.com/lithammer/shortuuid"
)

func InitS3(key, secret, region string) *session.Session {
	connect, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				key,
				secret,
				"",
			),
		},
	)
	if err != nil {
		log.Error("aws S3 Config error", err)
	}

	return connect
}

func Upload(sess *session.Session, file multipart.File, fileHeader multipart.FileHeader) (string, error) {

	var uid = shortuuid.New()

	var manager = s3manager.NewUploader(sess)
	var src, errO = fileHeader.Open()
	if errO != nil {
		log.Info(errO)
	}
	defer src.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	var res, err = manager.Upload(
		&s3manager.UploadInput{
			Bucket:       aws.String("airbnb-app"),
			Key:          aws.String(uid),
			ACL:          aws.String("public-read-write"),
			Body:         bytes.NewReader(buffer),
			ContentType:  aws.String(http.DetectContentType(buffer)),
			StorageClass: aws.String("STANDARD"),
		},
	)

	if err != nil {
		log.Info(res)
		log.Error(err)
	}

	var url = "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/" + uid

	return url, nil
}

func UpdateUpload(sess *session.Session, filename string, file multipart.File, fileHeader multipart.FileHeader) string {

	var src, err0 = fileHeader.Open()
	if err0 != nil {
		log.Info(err0)
	}
	defer src.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	var svc = s3.New(sess)

	var input = &s3.PutObjectInput{
		Bucket:       aws.String("airbnb-app"),
		Key:          aws.String(filename),
		ACL:          aws.String("public-read-write"),
		Body:         bytes.NewReader(buffer),
		ContentType:  aws.String(http.DetectContentType(buffer)),
		StorageClass: aws.String("STANDARD"),
	}

	var res, err = svc.PutObject(input)
	if err != nil {
		log.Info(res)
	}

	return "success to update image"
}
