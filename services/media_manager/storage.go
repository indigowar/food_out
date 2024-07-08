package main

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type storage struct {
	c *minio.Client

	bucketExists bool
}

func (s *storage) Save(ctx context.Context, name string, ext string, size int64, data []byte) error {
	// if !s.bucketExists {
	// 	exists, err := s.c.BucketExists(ctx, bucketName)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !exists {
	// 		err = s.c.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	//
	// 	s.bucketExists = true
	// }
	//
	// _, err := s.c.PutObject(ctx, bucketName, name, bytes.NewReader(data), size, minio.PutObjectOptions{
	// 	ContentType: ext,
	// })
	//
	// return err

	// TODO: Implement
	panic("not implemented")
}

func (s *storage) Load(ctx context.Context, name string) ([]byte, string, error) {
	// result, err := s.c.GetObject(ctx, bucketName, name, minio.GetObjectOptions{})
	//
	// if err != nil {
	// 	return nil, "", err
	// }

	// TODO: Implement
	panic("not implemented")
}
