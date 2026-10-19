package storage

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/hexanecreations/docsai/backend/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New(cfg *config.Config) (*minio.Client, error) {

	client, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			cfg.MinioAccessKey,
			cfg.MinioSecretKey,
			"",
		),
		Secure: cfg.MinioUseSSL,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

type Service interface {
	Upload(
		ctx context.Context,
		file *multipart.FileHeader,
		objectKey string,
	) (string, error)

	Delete(
		ctx context.Context,
		objectKey string,
	) error

	GetPresignedURL(
		ctx context.Context,
		objectKey string,
		expiry time.Duration,
	) (string, error)
}

type service struct {
	client *minio.Client
	bucket string
}

func NewService(
	client *minio.Client,
	bucket string,
) Service {
	return &service{
		client: client,
		bucket: bucket,
	}
}

func (s *service) Upload(
	ctx context.Context,
	file *multipart.FileHeader,
	objectKey string,
) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = s.client.PutObject(
		ctx,
		s.bucket,
		objectKey,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

func (s *service) Delete(
	ctx context.Context,
	objectKey string,
) error {

	return s.client.RemoveObject(
		ctx,
		s.bucket,
		objectKey,
		minio.RemoveObjectOptions{},
	)
}

func (s *service) GetPresignedURL(
	ctx context.Context,
	objectKey string,
	expiry time.Duration,
) (string, error) {

	presignedURL, err := s.client.PresignedGetObject(
		ctx,
		s.bucket,
		objectKey,
		expiry,
		url.Values{},
	)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
