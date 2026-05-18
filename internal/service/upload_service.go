package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"server/internal/storage"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFile(fileName string, file io.Reader, contentType string) (string, error) {
	client := storage.NewR2Client()

	key := fmt.Sprintf("%d-%s", time.Now().Unix(), fileName)

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("R2_BUCKET")),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := os.Getenv("R2_PUBLIC_URL") + "/" + key
	return url, nil
}

func DeleteFile(fileURL string) error {
	client := storage.NewR2Client()

	key := strings.TrimPrefix(fileURL, os.Getenv("R2_PUBLIC_URL")+"/")

	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET")),
		Key:    aws.String(key),
	})

	return err
}