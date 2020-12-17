package sss

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-north-1"),
	Credentials: credentials.NewStaticCredentials(os.Getenv("IOS_BUCKET_ID"),
		os.Getenv("IOS_BUCKET_SECRET"), ""),
}))

var svc = s3.New(sess)

func UploadPhoto(form *multipart.Form, recipeId uint64) (link string, err error) {
	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return "", errors.New("no file in multipart form")
	}

	file, err := fileHeaders[0].Open()
	if err != nil {
		return "", err
	}
	defer func() {
		err = file.Close()
	}()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	splitName := strings.Split(fileHeaders[0].Filename, ".")
	ext := splitName[len(splitName)-1]

	t := time.Now()
	link = fmt.Sprintf("%d%d%d%d%d%d-%d", t.Year(),
		t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second(), recipeId) + "-avatar." + ext

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("IOS_BUCKET_NAME")),
		Key:    aws.String(link),
		Body:   strings.NewReader(buf.String()),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	link = "https://huvalkiosbucket2.s3.eu-north-1.amazonaws.com/" + link

	return link, nil
}
