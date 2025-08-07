package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Uploader() (*S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3Uploader{
		Client:     client,
		BucketName: os.Getenv("S3_BUCKET"),
	}, nil
}

func (u *S3Uploader) UploadFile(file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s/%d-%s", folder, time.Now().Unix(), fileHeader.Filename)

	_, err := u.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &u.BucketName,
		Key:    &key,
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    "public-read", // Cho FE truy cập
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.BucketName, key)
	return url, nil
}
