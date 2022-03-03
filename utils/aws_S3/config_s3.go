package awss3

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/gommon/log"
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

func Upload(sess *session.Session, file multipart.File, fileHeader *multipart.FileHeader) string {

	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	fileTemp := "" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	_, err := s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("airbnb-app"),
		ACL:         aws.String("public-read-write"),
		ContentType: aws.String(http.DetectContentType(buffer)),
		Key:         aws.String(fileTemp),
		Body:        bytes.NewReader(buffer),
	})

	if err != nil {
		log.Error("Upload error : ", err)
	}

	return fileTemp
}
