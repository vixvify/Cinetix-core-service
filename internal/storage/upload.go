package storage

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewR2Client() *s3.Client {
	endpoint := "https://" + os.Getenv("R2_ACCOUNT_ID") + ".r2.cloudflarestorage.com"

	return s3.New(s3.Options{
		Region: "auto",
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				os.Getenv("R2_ACCESS_KEY"),
				os.Getenv("R2_SECRET_KEY"),
				"",
			),
		),
		BaseEndpoint: aws.String(endpoint),
	})
}