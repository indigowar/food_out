package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type storage struct {
	c *minio.Client

	bucketName   string
	bucketExists bool
}

func (s *storage) Save(ctx context.Context, name string, contentType string, size int64, data []byte) error {
	if err := s.checkBucketExistance(ctx); err != nil {
		return err
	}

	_, err := s.c.PutObject(
		ctx,
		s.bucketName,
		name,
		bytes.NewReader(data),
		size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to save the object: %w", err)
	}

	return nil
}

func (s *storage) Load(ctx context.Context, name string) ([]byte, string, error) {
	if err := s.checkBucketExistance(ctx); err != nil {
		return nil, "", err
	}

	object, err := s.c.GetObject(ctx, s.bucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", fmt.Errorf("failed to load the object: %w", err)
	}

	stat, err := object.Stat()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get stat of object: %w", err)
	}

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read the object: %w", err)
	}

	return data, stat.ContentType, nil
}

func (s *storage) checkBucketExistance(ctx context.Context) error {
	if s.bucketExists {
		return nil
	}

	exists, err := s.c.BucketExists(ctx, s.bucketName)
	if err != nil {
		return fmt.Errorf("failed to check the bucket existance: %w", err)
	}
	if exists {
		return nil
	}

	err = s.c.MakeBucket(ctx, s.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return fmt.Errorf("failed to create a bucket: %w", err)
	}

	s.bucketExists = true
	return nil
}
