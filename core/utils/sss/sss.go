package sss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

func UploadPhoto(data []byte, recipeId uint64) (link string, err error) {
	ext := "png"

	t := time.Now()
	link = fmt.Sprintf("%d%d%d%d%d%d-%d", t.Year(),
		t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second(), recipeId) + "-photo." + ext

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("IOS_BUCKET_NAME")),
		Key:    aws.String(link),
		Body:   strings.NewReader(string(data)),
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	link = "https://huvalkiosbucket2.s3.eu-north-1.amazonaws.com/" + link

	return link, nil
}
